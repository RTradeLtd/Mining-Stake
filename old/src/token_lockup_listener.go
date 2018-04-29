package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/RTradeLtd/Mining-Stake/TokenLockup"
	bbolt "github.com/coreos/bbolt"
	gopass "github.com/howeyc/gopass"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func bBoltSetup(dbPath string) *bbolt.DB {
	// setup the bbolt database
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal("error opening bolt database ", err)
	}
	// create bucket if it does not exist
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("stakers"))
		return err
	})
	if err != nil {
		log.Fatal("error establishing staker bucket in database")
	}
	return db
}

func updateBboltDb(address common.Address, id *big.Int, db *bbolt.DB) {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("stakers"))
		if err != nil {
			log.Fatal("error connecting to stakers bucket in bolt database ", err)
		}
		err = bucket.Put([]byte(address.Bytes()), []byte(id.Bytes()))
		return err
	})
	if err != nil {
		log.Fatal("error updating bolt stakers bucket ", err)
	}
}

func retrieveBucketInformationForAddress(address common.Address, db *bbolt.DB) *big.Int {
	var response []byte
	db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("stakers"))
		response = bucket.Get([]byte(address.Bytes()))
		return nil
	})
	i := new(big.Int)
	i.SetBytes(response)
	return i
}

// this is used to calculate a users currently active hash rate so we can easily factor multiple stake payments into a single payment
func calculateActiveHashRate(contract *TokenLockup.TokenLockup, address common.Address, db *bbolt.DB) *big.Int {
	var one = big.NewInt(1)
	start := big.NewInt(0)
	end := retrieveBucketInformationForAddress(address, db)
	khSecSum := big.NewInt(0)
	//    rtcStaked, khSec, depositDate, releaseDate, id, enabled :=
	for i := new(big.Int).Set(start); i.Cmp(end) == -1; i.Add(i, one) {
		_, khSec, _, _, _, enabled, err := contract.GetStakerStruct(nil, address, i)
		if err != nil {
			log.Fatal("error calculcating hash rate ", err)
		}
		if enabled == true {
			khSecSum.Add(khSecSum, khSec)
		}
	}
	return khSecSum
}

func buildPayoutData(contract *TokenLockup.TokenLockup, addresses []common.Address, db *bbolt.DB) {
	var m = make(map[common.Address]*big.Int)
	for i := 0; i < len(addresses); i++ {
		m[addresses[i]] = calculateActiveHashRate(contract, addresses[i], db)
	}
}

func eventParser(contract *TokenLockup.TokenLockup, db *bbolt.DB) {
	var ch = make(chan *TokenLockup.TokenLockupStakeDeposited)
	sub, err := contract.WatchStakeDeposited(nil, ch)
	if err != nil {
		log.Fatal("error creationg event subscription for stake deposited ", err)
	} else {
		fmt.Println("succesfully established event subsription for stake deposited")
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal("error parsing event ", err)
		case evLog := <-ch:
			fmt.Println("successfully retrieved log")
			fmt.Printf(
				"%v, %v, %v, %v, %v\n", evLog.Depositer, evLog.Amount, evLog.WeeksStaked, evLog.KhSec, evLog.Id)
			sendEmail(evLog.Depositer, evLog.Amount, evLog.WeeksStaked, evLog.KhSec, evLog.Id)
			updateBboltDb(evLog.Depositer, evLog.Id, db)
			calculateActiveHashRate(contract, evLog.Depositer, db)
		}
	}
}

func sendEmail(_depositer common.Address, _amountStaked *big.Int, _duration *big.Int, _khSec *big.Int, _id *big.Int) {
	content := fmt.Sprintf("<br>Staker: 0x%x<br><br>RTC Staked: %v<br><br>Weeks Staked: %v<br><br>KhSec: %v<br><br>Stake Id: %v<br>", _depositer, _amountStaked, _duration, _khSec, _id)
	from := mail.NewEmail("stake-sendgrid-api", "sgapi@rtradetechnologies.com")
	subject := "New Stake Deposit Detected In Staking Contract"
	to := mail.NewEmail("Mining Stake", "stake@rtradetechnologies.com")

	mContent := mail.NewContent("text/html", content)
	m := mail.NewV3MailInit(from, subject, to, mContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(m)
	if err != nil {
		fmt.Println("error sending message ", err)
	} else {
		fmt.Println("message send status code ", response.StatusCode)
	}
}

func main() {

	// setup scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter raw json key file")
	scanner.Scan()
	key := scanner.Text()

	fmt.Println("enter password to decrypt key")
	pass, err := gopass.GetPasswd()
	if err != nil {
		log.Fatal("error reading password ", err)
	}

	fmt.Println("enter address of staking contract")
	scanner.Scan()
	contractAddress := scanner.Text()

	// setup the bolt database
	db := bBoltSetup("stakers.db")

	// make sure we can use our send grid system
	if len(os.Getenv("SENDGRID_API_KEY")) <= 15 {
		log.Fatal("no valid send grid api key detected in environment variable, please enter your api key then re-run")
	}

	// connect to network
	fmt.Println("please enter path to ipc file")
	scanner.Scan()
	fmt.Println("initiating ipc connection")
	client, err := ethclient.Dial(scanner.Text())
	if err != nil {
		log.Fatal("error connecting to blockchain ", err)
	} else {
		fmt.Println("ipc connection successfully established")
	}

	fmt.Println("unlocking eth account")
	auth, err := bind.NewTransactor(strings.NewReader(key), string(pass))
	if err != nil {
		log.Fatalf("error unlocking account")
	} else {
		fmt.Println("unlock successful", auth)
	}

	tokenLockup, err := TokenLockup.NewTokenLockup(common.HexToAddress(contractAddress), client)

	initializer := big.NewInt(0)
	minStake, err := tokenLockup.MINSTAKE(nil)
	// if minStake is equal to 0, we can consider the contract to have not been properly deployed
	if err != nil || minStake.Cmp(initializer) == 0 {
		log.Fatal("error connecting to contract", err)
	}

	eventParser(tokenLockup, db)
}
