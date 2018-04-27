package main

import (
	"log"

	"github.com/RTradeLtd/Mining-Stake/manager/commands"
)

/*
	Main file for the stake manager
*/

func main() { /*
		boltDB := database.BoltDB{}
		databasePath := os.Args[1]
		bucketName := os.Args[2]
		tokenLockupContractAddress := common.HexToAddress(os.Args[3])
		if err := boltDB.Setup(databasePath, bucketName, tokenLockupContractAddress); err != nil {
			log.Fatal(err)
		}*/
	Execute()
}

func Execute() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
