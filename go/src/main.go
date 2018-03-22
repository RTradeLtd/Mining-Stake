package main

import (
	"log"
	"fmt"
	"strings"
	"math/big"
	//ipfs_api "github.com/ipfs/go-ipfs-api"


	//ethereum "github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
  //"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/sendgrid/sendgrid-go"
	"./token_lockup"
)

const key = `{"address":"d72f0d88384c05c3d95c870ba98ac2d606939c65","crypto":{"cipher":"aes-128-ctr","ciphertext":"589a88ccbdaa312595343c907e944c8b9d9e133d443b43d4efa71c6c7cea26d0","cipherparams":{"iv":"4429d785f61dd7d37d7813a8a422d941"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f92dbdb8c2c4686a839978d9dab36601a2e950d001b6d7131dd9a22c68f32da1"},"mac":"9037da8e700215e1d79043a4fcac847768d27e28dfcd3ce16f094eb1d837f1e1"},"id":"6472fa0e-80e4-475a-8f35-ede98c37641e","version":3}`

func eventParser(contract *TokenLockup.TokenLockup) {
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
			depositer, amountStaked, duration, khsec, id := evLog.Depositer, evLog.Amount, evLog.WeeksStaked, evLog.KhSec, evLog.Id
			fmt.Printf(
				"%v, %v, %v, %v, %v", depositer, amount, duration, khsec, id)
		}
	}
}

func sendEmail(_depositer common.Address, _amountStaked big.Int, _duration big.Int, _khSec big.Int, _id big.Int) {
	content = fmt.Sprintf()
	from := mail.NewEmail("stake-sendgrid-api", "sgapi@rtradetechnologies.com")
    subject := "New Stake Deposit Detected In Staking Contract"
    to := mail.NewEmail("Mining Stake", "stake@rtradetechnologies.com")
    plainTextContent := "and easy to do anywhere, even with Go"
    htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
    message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
    client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := client.Send(message)
    if err != nil {
    log.Println(err)
    } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
    }
}


func main() {

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
		fmt.Println("unlock successful")
	}

	
	address, tx, tokenLockup, err := TokenLockup.DeployTokenLockup(auth, client)
	if err != nil {
			log.Fatal("error deploying token ", err)
	} else {
		fmt.Println("token deploy successful")
		fmt.Printf("Token address 0x%x\n", address)
		fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
		fmt.Println("Press enter once contract is sucessfully deployed")
		fmt.Scanln()
	}

	minStake, err := tokenLockup.MINSTAKE(nil)
	if err != nil {
		log.Fatal("error connecting to contract", err)
	} else {
		fmt.Println("contract connection successful, min stake ", minStake)
	}
	go eventParser(tokenLockup)
	fmt.Println("looping for enternity")
	for {}
}