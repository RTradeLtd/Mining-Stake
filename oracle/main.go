package oracle

import (
	"log"
	"os"

	"github.com/RTradeLtd/Mining-Stake/database"
	"github.com/RTradeLtd/Mining-Stake/manager"
	"github.com/ethereum/go-ethereum/common"
	sendgrid "github.com/sendgrid/sendgrid-go"
)

/*
This package is used for oraclize style updates of the smart contract with off-chain data.

We thought about using Oraclize, however due to the extreme long term costs occured by 10 minute updates of off-chain data,
the only way we could reasonably implement that would be to offload the fees onto the end-user, which we cant justify.
*/

const rtcPrice = 0.125 //
const fixedRTCPrice = true
const dbPath = "stakers.db"
const bucketName = "stakers"
const tokenLockupAddress = "0x5ae6c285eeb2e5a9234956cbcf9dea2c97c3a773"
const rpcURL = "http://127.0.0.1:8545"
const ipcPath = "/home/solidity/.ethereum/rinkeby/geth.ipc"
const key = `{"address":"069ba77207ad40b7d386f8e2979a9337a36f991c","crypto":{"cipher":"aes-128-ctr","ciphertext":"b1218c0a8d354cddcb288d021a1e76a5a8617e32b78cff0d9769b6b663851516","cipherparams":{"iv":"f1f1e9461f2e17c3ca6866173b953860"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"b2aaac5ac70f95d81b41ca1042e6ddbb1b73f456ba03c6feb8d1eda7137b571b"},"mac":"4e1eadd303f63936806808059dc1fc3be0a9706de50e5b0f824e1a4bd1310e87"},"id":"54fe5587-3f4f-45b6-b895-d905275faaf5","version":3}`
const password = "password123"

// Oracle is used to interface with the RTrade Oracle network
type Oracle struct {
	Manager     *manager.Manager
	EthUsdPrice float64
	RtcUsdPrice float64
}

func setup() *Oracle {
	oracle := Oracle{
		Manager: &manager.Manager{}}
	return &oracle
}

// UpdateEthPrice is used to update the eth usd price
func (o *Oracle) UpdateEthPrice() {
	o.EthUsdPrice = o.Manager.RetrieveEthUsdPrice()
}

// UpdateRtcPrice is used to update the rtc usd price
func (o *Oracle) UpdateRtcPrice() {
	if fixedRTCPrice {
		o.RtcUsdPrice = rtcPrice
	}
}

func main() {
	oracle := setup()
	oracle.UpdateEthPrice()
	oracle.UpdateRtcPrice()

	// fethc the send grid API key
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if len(apiKey) < 15 {
		log.Fatal("invalid sendgrid key detected")
	}

	oracle.Manager = &manager.Manager{
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

	if err := oracle.Manager.AuthenticateWithNetwork(); err != nil {
		log.Fatal(err)
	}
}
