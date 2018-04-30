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
const dev = true

func main() {

	if len(os.Args) > 3 || len(os.Args) < 3 {
		log.Fatalf("improper invocation\n./Mining-Stake [stake|listen] [eth|rtc]")
	}

	if os.Args[1] != "stake" && os.Args[1] != "listen" {
		log.Fatalf("%s is not valid, must be stake or listen\n", os.Args[1])
	}

	if os.Args[1] == "stake" && os.Args[2] != "rtc" && os.Args[2] != "eth" {
		log.Fatalf("%s is not valid, must be rtc or eth\n", os.Args[2])
	}

	// fethc the send grid API key
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if len(apiKey) < 15 && dev == false {
		log.Fatal("invalid sendgrid key detected")
	}

	ethKey := os.Getenv("ETH_KEY")
	if len(ethKey) < 25 {
		log.Fatal("invalid eth key")
	}

	ethPass := os.Getenv("ETH_PASS")
	if len(ethPass) < 1 {
		log.Fatal("invalid eth pass")
	}

	manager := &manager.Manager{
		Key:            ethKey,
		Password:       ethPass,
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
		if weekday.String() == "Saturday" {
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
