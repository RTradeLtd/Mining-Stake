package manager

import (
	"log"

	"github.com/RTradeLtd/Mining-Stake/database"
	"github.com/ethereum/go-ethereum/common"
)

const dbPath = "stakers.db"
const bucketName = "stakers"
const tokenLockupAddress = "0x"
const rpcURL = "http://127.0.0.1:8545"
const ipcPath = "/home/solidity/.ethereum/rinkeby/geth.ipc"
const key = "..."
const password = "password123"

func main() {

	manager := &Manager{}
	manager.Bolt = &database.BoltDB{
		StakeIDBucketName:          bucketName,
		TokenLockupContractAddress: common.HexToAddress(tokenLockupAddress),
	}

	//setup the bolt database
	if err := manager.Bolt.Setup(dbPath, bucketName, manager.Bolt.TokenLockupContractAddress); err != nil {
		log.Fatal(err)
	}

	// setup our connection to the rpc backend
	manager.EstablishRPCConnection(rpcURL)

	if err := manager.AuthenticateWithNetwork(); err != nil {
		log.Fatal(err)
	}
}
