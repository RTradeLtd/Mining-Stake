package manager

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

/*
	Contains methods related to stake management
*/

// ParseBlockStatistics is used to retrieve block params to
// allow us to calculate payout data
func (m *Manager) ParseBlockStatistics() error {
	//  hard coded value for testing phase
	// this will be pulled from the blockchain at runtime
	diffTh := float64(3200)
	ethUsd := m.RetrieveEthUsdPrice()
	currentBlockNum, err := m.RPC.EthBlockNumber()
	if err != nil {
		return err
	}
	previousBlockNum := currentBlockNum - 1
	currentBlock, err := m.RPC.EthGetBlockByNumber(currentBlockNum, false)
	if err != nil {
		return err
	}
	previousBlock, err := m.RPC.EthGetBlockByNumber(previousBlockNum, false)
	if err != nil {
		return err
	}

	currentBlockTimestamp := currentBlock.Timestamp
	previousBlockTimestamp := previousBlock.Timestamp
	blockTimeSec := currentBlockTimestamp - previousBlockTimestamp
	blockReward := float64(3) // this doesnt change often so we hard code
	m.Block = &BlockStatistics{
		DiffTh:       diffTh,
		BlockTimeSec: float64(blockTimeSec),
		BlockReward:  blockReward,
		EthPrice:     ethUsd,
	}
	return nil
}

// CalculateActiveHashRate used to calculate active hash rate for a staker
// active hashrate is defined as the combined hash rate of all actively enabled stakes
func (m *Manager) CalculateActiveHashRate(address common.Address) *big.Int {
	var one = big.NewInt(1)
	var zero = big.NewInt(0)
	start := big.NewInt(0)
	end := m.Bolt.RetrieveStakeIDInformationForAddress(address)
	khSecSum := big.NewInt(0)
	// generate new big int, and set it to start
	// compare i to end, if less than end (-1) continue, increment counter by 1
	if end.Cmp(one) == 0 {
		for i := new(big.Int).Set(start); i.Cmp(end) == -1; i.Add(i, one) {
			_, khSec, _, _, _, enabled, err := m.TokenLockupContractHandler.GetStakerStruct(nil, address, i)
			if err != nil {
				log.Fatal("error calculcating hash rate ", err)
			}
			if enabled == true {
				khSecSum.Add(khSecSum, khSec)
			}
		}
		return khSecSum
	}
	_, khSec, _, _, _, enabled, err := m.TokenLockupContractHandler.GetStakerStruct(nil, address, zero)
	if err != nil {
		log.Fatal("error calculating active hash rate ", err)
	}
	if enabled == true {
		return khSec
	} else {
		return zero
	}

}

// CalculateUsdPayout is used to calculate someone's USD payout given their hash rate
func (m *Manager) CalculateUsdPayout(mhSec float64, diffTH float64, blockTimeSec float64, blockReward float64, ethPrice float64) float64 {
	usdEarningsPerDay := float64((mhSec * 1e6 / ((diffTH / blockTimeSec) * 1000 * 1e9)) * ((60 / blockTimeSec) * blockReward) * (60 * 24) * (ethPrice))
	return usdEarningsPerDay
}

// CalculateEthPayout is used to calculate someone's ETH payout earnings for a week
func (m *Manager) CalculateEthPayout(mhSec float64, diffTH float64, blockTimeSec float64, blockReward float64) (float64, error) {
	ethEarningsPerDay := float64((mhSec * 1e6 / ((diffTH / blockTimeSec) * 1000 * 1e9)) * ((60 / blockTimeSec) * blockReward) * (60 * 24))
	return ethEarningsPerDay, nil
}

// ConstructRtcPayoutData is used to build payout rtc stake payout data
// current implementation routes to one address at a time
// to fix this we will need to rework some of the logic
func (m *Manager) ConstructRtcPayoutData() map[common.Address]*big.Int {
	var stakerMap = make(map[common.Address]uint64)
	var activeStakers = make(map[common.Address]*big.Int)
	stakerMap = m.Bolt.FetchStakeIDs()
	var addresses []common.Address
	var rtcs []*big.Int
	for k := range stakerMap {
		addresses = append(addresses, k)
		hash := m.CalculateActiveHashRate(k)
		if hash.Cmp(big.NewInt(0)) <= 0 {
			continue
		}
		// since we're dealing with big numbers, we can simply just divide by 10^18, we need to do that by utilizing big int variables
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
		dv := new(big.Int).Div(hash, exp)
		mHash := float64(dv.Int64()) / float64(1000)
		usdEarnings := m.CalculateUsdPayout(mHash, m.Block.DiffTh, m.Block.BlockTimeSec, m.Block.BlockReward, m.Block.EthPrice)
		percentUsd := new(big.Float).Mul(big.NewFloat(usdEarnings), big.NewFloat(0.1))
		percentUsdFloat, _ := percentUsd.Float64()
		fmt.Println("USD Float ", percentUsdFloat)
		// we are using a fixed RTC price for now
		rtcFloat := percentUsdFloat / 0.16
		rtc := FloatToBigInt(rtcFloat)
		rtcs = append(rtcs, rtc)
		activeStakers[k] = rtc
	}
	tx, err := m.TokenLockupContractHandler.RouteRtcRewards(m.TransactOpts, addresses, rtcs)
	if err != nil {
		log.Fatal("error routing token payments ", err)
	}
	fmt.Println("token payments routed successfully")
	fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
	return activeStakers
}

// ConstructEthPayoutData is used to build, and send
// eth payouts
func (m *Manager) ConstructEthPayoutData() map[common.Address]*big.Int {
	var stakerMap = make(map[common.Address]uint64)
	var activeStakers = make(map[common.Address]*big.Int)
	stakerMap = m.Bolt.FetchStakeIDs()
	var addresses []common.Address
	var eths []*big.Int
	var totalEarnings = big.NewInt(0)
	for addr := range stakerMap {
		addresses = append(addresses, addr)
		hash := m.CalculateActiveHashRate(addr)
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
		dv := new(big.Int).Div(hash, exp)
		mHash := float64(dv.Int64()) / float64(1000)
		ethEarnings, _ := m.CalculateEthPayout(mHash, m.Block.DiffTh, m.Block.BlockTimeSec, m.Block.BlockReward)
		ethEarningsBig := FloatToBigInt(ethEarnings)
		weekEarnings := new(big.Int).Mul(ethEarningsBig, big.NewInt(7))
		totalEarnings = new(big.Int).Add(totalEarnings, weekEarnings)
		eths = append(eths, weekEarnings)
		activeStakers[addr] = weekEarnings
	}
	m.TransactOpts.Value = totalEarnings
	tx, err := m.TokenLockupContractHandler.RouteEthReward(m.TransactOpts, addresses, eths)
	if err != nil {
		log.Fatal("error sending token ", err)
	}
	fmt.Printf("TX Hash 0x%x\n", tx.Hash())
	return activeStakers
}
