package main

import (
	"fmt"
	"log"

	"github.com/RTradeLtd/Mining-Stake/database"
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

	manager := &manager.Manager{
		Password: password,
		Key:      key,
		IpcPath:  ipcPath,
		RPCURL:   rpcURL,
		Bolt: &database.BoltDB{
			StakeIDBucketName:          bucketName,
			TokenLockupContractAddress: common.HexToAddress(tokenLockupAddress),
		}}

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

	var stakes = make(map[common.Address]uint64)
	stakes = manager.Bolt.FetchStakeIDs()
	fmt.Println(stakes)
	manager.ConstructRtcPayoutData()
}
