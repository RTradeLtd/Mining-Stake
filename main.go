package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/RTradeLtd/Mining-Stake/bindings"
	"github.com/RTradeLtd/Mining-Stake/database"
	"github.com/RTradeLtd/Mining-Stake/listener"
	"github.com/RTradeLtd/Mining-Stake/manager"
	sendgrid "github.com/sendgrid/sendgrid-go"
)

const dbPath = "stakers.db"
const stakeBucket = "stakers"
const emailBucket = "emails"
const oracleBucket = "oracle"
const tokenLockupAddress = "0x5ae6c285eeb2e5a9234956cbcf9dea2c97c3a773"
const oracleContractAddress = "0x5ae6c285eeb2e5a9234956cbcf9dea2c97c3a773"
const rpcURL = "http://127.0.0.1:8545"
const ipcPath = "/home/solidity/.ethereum/rinkeby/geth.ipc"
const dev = true

var emptyString string

func deployRtc(manager *manager.Manager) {
	address, tx, _, err := bindings.DeployRTCoin(manager.TransactOpts, manager.EthClient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Contract Address %s\n", address.String())
	fmt.Printf("Transaction Hash 0x%x\n", tx.Hash())
}
func main() {

	if len(os.Args) > 3 || len(os.Args) < 3 {
		log.Fatalf("improper invocation\n./Mining-Stake [stake|listen|oracle|deploy] [eth|rtc|noop]")
	}

	if os.Args[1] != "stake" && os.Args[1] != "listen" && os.Args[1] != "oracle" && os.Args[1] != "deploy" {
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
		Key:                        ethKey,
		Password:                   ethPass,
		SendGridAPIKey:             apiKey,
		SendGridClient:             sendgrid.NewSendClient(apiKey),
		TokenLockupContractAddress: common.HexToAddress(tokenLockupAddress),
		OracleContractAddress:      common.HexToAddress(oracleContractAddress),
		IpcPath:                    ipcPath,
		RPCURL:                     rpcURL,
		Bolt:                       &database.BoltDB{},
	}

	//authenticate with the eth network
	// this loads both the oracle and token lockup contracts
	if err := manager.AuthenticateWithNetwork(); err != nil {
		log.Fatal(err)
	}

	//setup the bolt database
	if err := manager.Bolt.Setup(dbPath, stakeBucket, emailBucket, oracleBucket); err != nil {
		log.Fatal(err)
	}

	if os.Args[1] == "deploy" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("enter contract to deploy")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "rtc":
			deployRtc(manager)
			os.Exit(0)
		default:
			fmt.Println("unrecognized option")
			os.Exit(1)
		}
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
		m := manager.ConstructRtcPayoutData()
		manager.SendEmailsForStakePayout(m, "rtc")
	} else if os.Args[1] == "eth" {
		currDate := time.Now()
		weekday := currDate.Weekday()
		if weekday.String() == "Saturday" {
			m := manager.ConstructEthPayoutData()
			manager.SendEmailsForStakePayout(m, "eth")
		} else {
			log.Fatal("today is not Sunday")
		}
	}
}

// Listen is used to listen for ongoing stakes
func Listen(manager *manager.Manager) {
	listener.EventParser(manager)
}
