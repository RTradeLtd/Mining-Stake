package main

import (
	"log"
	"fmt"
	"strings"
	"math/big"
	"os"

	// bbolt will be used to store active stakers
	bbolt "github.com/coreos/bbolt"
	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	
	"./token_lockup"
)

const key = `{"address":"d72f0d88384c05c3d95c870ba98ac2d606939c65","crypto":{"cipher":"aes-128-ctr","ciphertext":"589a88ccbdaa312595343c907e944c8b9d9e133d443b43d4efa71c6c7cea26d0","cipherparams":{"iv":"4429d785f61dd7d37d7813a8a422d941"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f92dbdb8c2c4686a839978d9dab36601a2e950d001b6d7131dd9a22c68f32da1"},"mac":"9037da8e700215e1d79043a4fcac847768d27e28dfcd3ce16f094eb1d837f1e1"},"id":"6472fa0e-80e4-475a-8f35-ede98c37641e","version":3}`

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

func retrieveBucketInformationForAddress(address common.Address, db *bbolt.DB) (*big.Int) {
	var response []byte
    db.View(func(tx *bbolt.Tx) error {
        bucket := tx.Bucket([]byte("stakers"))
        response = bucket.Get([]byte(address.Bytes()))
        //id.SetBytes([]byte(response))
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
			//fmt.Printf("Retrieving bolt information to test %v", retrieveBoltInformationForAddress(evLog.Depositer, db))
			calculateActiveHashRate(contract, evLog.Depositer, db)
		}
	}
}

func sendEmail(_depositer common.Address, _amountStaked *big.Int, _duration *big.Int, _khSec *big.Int, _id *big.Int) {
	content := fmt.Sprintf("<br>Staker: 0x%x<br><br>RTC Staked: %v<br><br>Weeks Staked: %v<br><br>KhSec: %v<br><br>Stake Id: %v<br>",_depositer, _amountStaked, _duration, _khSec, _id)
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

	db := bBoltSetup("stakers.db")

	if len(os.Getenv("SENDGRID_API_KEY")) <= 15 {
		log.Fatal("no valid send grid api key detected in environment variable")
	}
	// db := bBoltSetup("staker.db")
	// defer db.Close()
	fmt.Println("initiating ipc connection")
	client, err := ethclient.Dial("/home/solidity/.ethereum/geth.ipc")
	if err != nil {
		log.Fatal("error connecting to blockchain ", err)
	} else {
		fmt.Println("ipc connection successfully established")
	}

	fmt.Println("unlocking eth account")
	auth, err := bind.NewTransactor(strings.NewReader(key), "password123")
	if err != nil {
		log.Fatalf("error unlocking account")
	} else {
		fmt.Println("unlock successful", auth)
	}

	tokenLockup, err := TokenLockup.NewTokenLockup(common.HexToAddress("0x5Ae6C285eeB2e5a9234956cbCf9dea2C97C3A773"), client)	

	minStake, err := tokenLockup.MINSTAKE(nil)
	if err != nil {
		log.Fatal("error connecting to contract", err)
	} else {
		fmt.Println("contract connection successful, min stake ", minStake)
	}

	eventParser(tokenLockup, db)
}