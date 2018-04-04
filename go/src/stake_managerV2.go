package main

import (
	"log"
	"fmt"
	"strings"
	"math/big"
	"encoding/binary"
	"bufio"
	"os"

	// bbolt will be used to store active stakers
	bbolt "github.com/coreos/bbolt"
	//prompt "github.com/c-bata/go-prompt"
	ishell "gopkg.in/abiosoft/ishell.v2"
	"github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/ethstats"
	//"github.com/sendgrid/sendgrid-go"
	//"github.com/sendgrid/sendgrid-go/helpers/mail"
	
	"./token_lockup"
)

const key = `{"address":"d72f0d88384c05c3d95c870ba98ac2d606939c65","crypto":{"cipher":"aes-128-ctr","ciphertext":"589a88ccbdaa312595343c907e944c8b9d9e133d443b43d4efa71c6c7cea26d0","cipherparams":{"iv":"4429d785f61dd7d37d7813a8a422d941"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f92dbdb8c2c4686a839978d9dab36601a2e950d001b6d7131dd9a22c68f32da1"},"mac":"9037da8e700215e1d79043a4fcac847768d27e28dfcd3ce16f094eb1d837f1e1"},"id":"6472fa0e-80e4-475a-8f35-ede98c37641e","version":3}`
//EarningsPerMonth = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24*30)*(EthPrice)
//EarningsPerDay = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24)*(EthPrice)

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

// used to update the database
func updateBboltDb(address common.Address, id *big.Int, db *bbolt.DB) {
	err := db.Update(func(tx *bbolt.Tx) error {
		// creates the bucket if it doesnt exist
		bucket, err := tx.CreateBucketIfNotExists([]byte("stakers"))
		if err != nil {
			log.Fatal("error connecting to stakers bucket in bolt database ", err)
		}
		// adds data to bucket
		err = bucket.Put([]byte(address.Bytes()), []byte(id.Bytes()))
		return err
	})
	// if err == nil, then we can assume the database wasn't updated properly
	if err != nil {
		log.Fatal("error updating bolt stakers bucket ", err)
	}
}

// used to retrieve the stake ID for an account
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

// used to iterate over the  bucket, returning a map with the contents
func iterateOverBucket(db *bbolt.DB) map[common.Address]uint64 {
	var m = make(map[common.Address]uint64)
	db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("stakers"))
	    // Iterate over items in sorted key order.
	    if err := bucket.ForEach(func(k, v []byte) error {
	    	address := k
	    	hexAddr := fmt.Sprintf("0x%x", address)
	    	stakeId, _ := binary.Uvarint(v)
	        fmt.Printf("Staker\t0x%x\nStake ID\t%v\n", address, stakeId)
	        m[common.HexToAddress(hexAddr)] = stakeId
	        return nil
	    }); err != nil {
	        return err
	    }
	    return nil
	})
	return m
}

// this is used to calculate a users currently active hash rate so we can easily factor multiple stake payments into a single payment
func calculateActiveHashRate(contract *TokenLockup.TokenLockup, address common.Address, db *bbolt.DB) *big.Int {
	var one = big.NewInt(1)
	start := big.NewInt(0)
	end := retrieveBucketInformationForAddress(address, db)
	khSecSum := big.NewInt(0)
	// generate new big int, and set it to start
	// compare i to end, if less than end (-1) continue, increment counter by 1
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

// WIP
func calculatePayout(khSec *big.Int) {
	//EarningsPerMonth = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24*30)*(EthPrice)
	//EarningsPerDay = (UserHashMh * 1e6 / ((difficultyTH / BlockTimeSec)*1000*1e9))*((60/ BlockTimeSec)*BlockReward)*(60*24)*(EthPrice)
}


// WIP (most likely will be discarded)
func buildPayoutData(contract *TokenLockup.TokenLockup, addresses []common.Address, db *bbolt.DB) {
	var m = make(map[common.Address]*big.Int)
	for i := 0; i < len(addresses); i++ {
		m[addresses[i]] = calculateActiveHashRate(contract, addresses[i], db)
	}
}

// used to create an RPC connection with the block chain
func establishRpcConnection(rpcUrl string) *rpc.Client {
	rpcClient, err := rpc.Dial(rpcUrl)
	if err != nil {
		log.Fatal("error establishing RPC connection ", err)
	}
	return rpcClient
}

// authenticates with the blockchain, and the staking contract
func authenticateWithContract()  (*ethclient.Client, *bind.TransactOpts, *TokenLockup.TokenLockup) {
	fmt.Println("initiating ipc connection")
	client, err := ethclient.Dial("/home/solidity/.ethereum/rinkeby/geth.ipc")
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

	return client, auth, tokenLockup
}

func main() {
	// create map to store stake data
	var m = make(map[common.Address]uint64)

	// setup rpc client connection
	fmt.Println("setting up rpc client")
	rpcClient := establishRpcConnection("http://127.0.0.1:8545")
	// setup connection to bolt database
	fmt.Println("setting up bolt database")
	db := bBoltSetup("stakers.db")

	// make sure we can interact with the contract
	fmt.Println("establishing connection with contract")
	_, _, tokenLockup := authenticateWithContract()


	// used to create a new shell
	fmt.Println("establishing shell")
    shell := ishell.New()

    // display welcome info.
    shell.Println("RTrade Technology Stake Manager")

	shell.AddCmd(&ishell.Cmd{
	    Name: "single-payout",
	    Help: "construct payout data for a single staker",
	    Func: func(c *ishell.Context) {
	        // disable the '>>>' for cleaner same line input.
	        c.ShowPrompt(false)
	        defer c.ShowPrompt(true) // yes, revert after login.
	        c.Print("Address: ")
	        address := c.ReadLine()
	        hash := calculateActiveHashRate(tokenLockup, common.HexToAddress(address), db)
	        c.Print(hash)
	    },
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "construct-payout-data",
		Help: "build payout data for active stakers",
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			m = iterateOverBucket(db)
			file, err := os.Create("test")
			writer := bufio.NewWriter(file)
			if err != nil {
				log.Fatal("error creating file")
			}
			for k, v := range m {
				hash := calculateActiveHashRate(tokenLockup, k, db)
				_, err = fmt.Fprintf(writer, "Address\t0x%x\nStake ID\t%v\nHash Rate\t%v\n\n\n", k, v, hash)
				if err != nil {
					log.Fatal("error writing to file")
				}
			}
			writer.Flush()
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "block-number",
		Help: "retrieve latest block number",
		Func: func(c *ishell.Context) {
			var lastBlock types.Block
			var num *big.Int
			err := rpcClient.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
			if err != nil {
				c.Print("unable to retrieve latest block ", err)
			}
			
			c.Print("latest block: ", num)
			},
		})

	shell.AddCmd(&ishell.Cmd{
		Name: "iterate-over-bucket",
		Help: "iterates over the stakers bucket",
		Func: func(c *ishell.Context) {
			m = iterateOverBucket(db)
			fmt.Println(m)
			c.Print("db iter finished\n")
		},
	})

    // run shell
    shell.Run()

}