package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/RTradeLtd/Mining-Stake/database"
	"github.com/RTradeLtd/Mining-Stake/manager"
	"github.com/ethereum/go-ethereum/common"
)

const dbPath = "stakers.db"
const bucketName = "stakers"
const tokenLockupAddress = "0x"
const rpcURL = "http://127.0.0.1:8545"
const ipcPath = "/home/solidity/.ethereum/rinkeby/geth.ipc"
const key = `{"address":"d72f0d88384c05c3d95c870ba98ac2d606939c65","crypto":{"cipher":"aes-128-ctr","ciphertext":"589a88ccbdaa312595343c907e944c8b9d9e133d443b43d4efa71c6c7cea26d0","cipherparams":{"iv":"4429d785f61dd7d37d7813a8a422d941"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f92dbdb8c2c4686a839978d9dab36601a2e950d001b6d7131dd9a22c68f32da1"},"mac":"9037da8e700215e1d79043a4fcac847768d27e28dfcd3ce16f094eb1d837f1e1"},"id":"6472fa0e-80e4-475a-8f35-ede98c37641e","version":3}`
const password = "password123"

func main() {

	manager := &manager.Manager{}
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

	// connect to the backend
	if err := manager.AuthenticateWithNetwork(); err != nil {
		log.Fatal(err)
	}

	// parse through block headers
	if err := manager.ParseBlockStatistics(); err != nil {
		log.Fatal(err)
	}

	var stakes = make(map[common.Address]uint64)
	var hashRates = make(map[common.Address]*big.Int)
	var usdPayouts = make(map[common.Address]float64)
	stakes = manager.Bolt.FetchStakeIDs()
	for k, v := range stakes {
		hashRates[k] = manager.CalculateActiveHashRate(k)
		fmt.Println(k, v, usdPayouts)
	}
}
