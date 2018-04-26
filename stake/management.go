package stake

import (
	"log"
	"math/big"

	"github.com/RTradeLtd/Mining-Stake/database"
	"github.com/ethereum/go-ethereum/common"
)

func CalculateActiveHashRate(address common.Address, b *database.BoltDB) *big.Int {
	var one = big.NewInt(1)
	var zero = big.NewInt(0)
	start := big.NewInt(0)
	end := b.RetrieveStakeIDInformationForAddress(address)
	khSecSum := big.NewInt(0)
	// generate new big int, and set it to start
	// compare i to end, if less than end (-1) continue, increment counter by 1
	if end.Cmp(one) == 0 {
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
	} else {
		_, khSec, _, _, _, enabled, err := contract.GetStakerStruct(nil, address, zero)
		if err != nil {
			log.Fatal("error calculating active hash rate")
		}
		if enabled == true {
			return khSec
		} else {
			return zero
		}
	}
}