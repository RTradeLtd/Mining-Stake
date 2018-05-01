package manager

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// UpdateEthUsdPrice is used to update the eth usd price on the destination contract
// note that it must be a whole number, ie 400, 500, not 500.10
func (m *Manager) UpdateEthUsdPrice(destinationContract common.Address) error {
	price := m.RetrieveEthUsdPriceNoDecimals()
	_, err := m.OracleContractHandler.UpdateEthPrice(m.TransactOpts, destinationContract, price)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRtcUsdPrice is used to update the RTC/USD price on the destination contract
func (m *Manager) UpdateRtcUsdPrice(destinationContract common.Address, price *big.Int) error {
	_, err := m.OracleContractHandler.UpdateRtcPrice(m.TransactOpts, destinationContract, price)
	if err != nil {
		return err
	}
	return nil
}
