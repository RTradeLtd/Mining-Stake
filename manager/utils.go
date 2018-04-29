package manager

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/RTradeLtd/Mining-Stake/TokenLockup"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onrik/ethrpc"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

	tokenLockup, err := TokenLockup.NewTokenLockup(m.Bolt.TokenLockupContractAddress, client)
	if err != nil {
		return err
	}
	m.ContractHandler = tokenLockup
	m.EthClient = client
	m.TransactOpts = auth
	return nil
}

// SendNotificationEmail is used to send us an email when we detect a stake in the system
func (m *Manager) SendNotificationEmail(depositer common.Address, amountStaked *big.Int, duration *big.Int, khSec *big.Int, id *big.Int) (int, error) {
	content := fmt.Sprintf("<br>Staker: 0x%x<br><br>RTC Staked: %v<br><br>Weeks Staked: %v<br><br>KhSec: %v<br><br>Stake Id: %v<br>", depositer, amountStaked, duration, khSec, id)
	from := mail.NewEmail("stake-sendgrid-api", "sgapi@rtradetechnologies.com")
	subject := "New Stake Deposit Detected In Staking Contract"
	to := mail.NewEmail("Mining Stake", "stake@rtradetechnologies.com")

	mContent := mail.NewContent("text/html", content)
	mail := mail.NewV3MailInit(from, subject, to, mContent)

	response, err := m.SendGridClient.Send(mail)
	if err != nil {
		return 0, err
	} else {
		return response.StatusCode, nil
	}
}
