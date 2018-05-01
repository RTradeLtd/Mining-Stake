package manager

import (
	"log"
	"math/big"
	"strings"

	"github.com/RTradeLtd/Mining-Stake/Oracle"
	"github.com/RTradeLtd/Mining-Stake/TokenLockup"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onrik/ethrpc"
)

var emptyString string

// FloatToBigInt used to convert a float to big int
func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

// EstablishRPCConnection is used to connect to our rpc node
func (m *Manager) EstablishRPCConnection() {
	m.RPC = ethrpc.New(m.RPCURL)
}

// AuthenticateWithNetwork is used to authenticate with the ethereum network
func (m *Manager) AuthenticateWithNetwork() error {
	client, err := ethclient.Dial(m.IpcPath)
	if err != nil {
		return err
	}
	auth, err := bind.NewTransactor(strings.NewReader(m.Key), m.Password)
	if err != nil {
		return err
	}

	tokenLockup, err := TokenLockup.NewTokenLockup(m.TokenLockupContractAddress, client)
	if err != nil {
		return err
	}
	oracleContract, err := Oracle.NewOracle(m.OracleContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	m.OracleContractHandler = oracleContract
	m.TokenLockupContractHandler = tokenLockup
	m.EthClient = client
	m.TransactOpts = auth
	return nil
}
