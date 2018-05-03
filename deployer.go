package main

import (
	"fmt"
	"os"
)

func deployer() {

	key := os.Getenv("ETH_KEY")
	password := os.Getenv("ETH_PASS")

	fmt.Println(key, password)
}
