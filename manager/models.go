package manager

import (
	"github.com/RTradeLtd/Mining-Stake/TokenLockup"
	"github.com/RTradeLtd/Mining-Stake/database"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onrik/ethrpc"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/spf13/cobra"
)

// Response used to hold response data from cmc
type Response struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	Rank               string `json:"rank"`
	PriceUsd           string `json:"price_usd"`
	PriceBtc           string `json:"price_btc"`
	TwentyFourHrVolume string `json:"24h_volume_usd"`
	MarketCapUsd       string `json:"market_cap_usd"`
	AvailableSupply    string `json:"available_supply"`
	TotalSupply        string `json:"total_supply"`
	MaxSupply          string `json:"null"`
	PercentChange1h    string `json:"percent_change_1h"`
	PercentChange24h   string `json:"percent_change_24h"`
	PercentChange7d    string `json:"percent_change_7d"`
	LastUpdate         string `json:"last_updated"`
}

// Manager is a general purpose struct to interface with the
// token lockup contract
type Manager struct {
	ContractHandler *TokenLockup.TokenLockup
	Bolt            *database.BoltDB
	Block           *BlockStatistics
	RPC             *ethrpc.EthRPC
	EthClient       *ethclient.Client
	SendGridClient  *sendgrid.Client
	TransactOpts    *bind.TransactOpts
	Cmd             *cobra.Command
	SendGridAPIKey  string
	Password        string
	Key             string
	IpcPath         string
	RPCURL          string
}

// BlockStatistics hold block related statistics
type BlockStatistics struct {
	DiffTh       float64
	BlockTimeSec float64
	BlockReward  float64
	EthPrice     float64
}
