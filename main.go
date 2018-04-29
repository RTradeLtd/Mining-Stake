package main

import (
	"log"
	"os"
	"time"

	"github.com/sendgrid/sendgrid-go"

	"github.com/RTradeLtd/Mining-Stake/database"
	"github.com/RTradeLtd/Mining-Stake/listener"
	"github.com/RTradeLtd/Mining-Stake/manager"
	"github.com/ethereum/go-ethereum/common"
)

const dbPath = "stakers.db"
const bucketName = "stakers"
const tokenLockupAddress = "0x5ae6c285eeb2e5a9234956cbcf9dea2c97c3a773"
const rpcURL = "http://127.0.0.1:8545"
const ipcPath = "/home/solidity/.ethereum/rinkeby/geth.ipc"
const key = `{"address":"069ba77207ad40b7d386f8e2979a9337a36f991c","crypto":{"cipher":"aes-128-ctr","ciphertext":"b1218c0a8d354cddcb288d021a1e76a5a8617e32b78cff0d9769b6b663851516","cipherparams":{"iv":"f1f1e9461f2e17c3ca6866173b953860"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"b2aaac5ac70f95d81b41ca1042e6ddbb1b73f456ba03c6feb8d1eda7137b571b"},"mac":"4e1eadd303f63936806808059dc1fc3be0a9706de50e5b0f824e1a4bd1310e87"},"id":"54fe5587-3f4f-45b6-b895-d905275faaf5","version":3}`
const password = "password123"

func main() {

	if len(os.Args) > 3 || len(os.Args) < 3 {
		log.Fatalf("improper invocation\n./Mining-Stake [stake|listen] [eth|rtc]")
	}

	if os.Args[1] != "stake" && os.Args[1] != "listen" {
		log.Fatalf("%s is not valid, must be stake or listen\n", os.Args[1])
	}

	if os.Args[2] != "rtc" && os.Args[2] != "eth" {
		log.Fatalf("%s is not valid, must be rtc or eth\n", os.Args[2])
	}

	// fethc the send grid API key
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if len(apiKey) < 15 {
		log.Fatal("invalid sendgrid key detected")
	}

	manager := &manager.Manager{
		Password:       password,
		Key:            key,
		SendGridAPIKey: apiKey,
		SendGridClient: sendgrid.NewSendClient(apiKey),
		IpcPath:        ipcPath,
		RPCURL:         rpcURL,
		Bolt: &database.BoltDB{
			StakeIDBucketName:          bucketName,
			TokenLockupContractAddress: common.HexToAddress(tokenLockupAddress),
		},
	}

	//authenticate with the eht network
	if err := manager.AuthenticateWithNetwork(); err != nil {
		log.Fatal(err)
	}

	//setup the bolt database
	if err := manager.Bolt.Setup(dbPath, bucketName, manager.Bolt.TokenLockupContractAddress); err != nil {
		log.Fatal(err)
	}

	// setup our connection to the rpc backend
	manager.EstablishRPCConnection()

	// connect to the backend
	if err := manager.AuthenticateWithNetwork(); err != nil {
		log.Fatal(err)
	}

	// parse through block headers
	if err := manager.ParseBlockStatistics(); err != nil {
		log.Fatal(err)
	}

	if os.Args[1] == "stake" {
		Stake(manager)
	} else if os.Args[1] == "listen" {
		Listen(manager)
	} else {
		log.Fatal("invalid run option detected")
	}
}

// Stake is used to initiate stake payouts
func Stake(manager *manager.Manager) {
	if os.Args[1] == "rtc" {
		manager.ConstructRtcPayoutData()
	} else if os.Args[1] == "eth" {
		currDate := time.Now()
		weekday := currDate.Weekday()
		if weekday.String() == "Sunday" {
			manager.ConstructEthPayoutData()
		} else {
			log.Fatal("today is not Sunday")
		}
	}
}

// Listen is used to listen for ongoing stakes
func Listen(manager *manager.Manager) {
	listener.EventParser(manager)
}
