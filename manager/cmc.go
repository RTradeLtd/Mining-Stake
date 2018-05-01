package manager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

// RetrieveEthUsdPrice is used to retrieve eths usd pricing
func (m *Manager) RetrieveEthUsdPrice() float64 {
	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/ethereum/")
	if err != nil {
		log.Fatal("error reading website ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("error reading response ", err)
	}
	var decode []Response
	err = json.Unmarshal(body, &decode)
	if err != nil {
		log.Fatal("error unmarshling json ", err)
	}

	// TODO: add error handling
	f, _ := strconv.ParseFloat(decode[0].PriceUsd, 64)

	return f
}

// RetrieveEthUsdPriceNoDecimals is used to retrieve the eth usd price without decimals
// TODO: add error handling
func (m *Manager) RetrieveEthUsdPriceNoDecimals() *big.Int {
	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/ethereum/")
	if err != nil {
		log.Fatal("error reading website ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("error reading response ", err)
	}
	var decode []Response
	err = json.Unmarshal(body, &decode)
	if err != nil {
		log.Fatal("error unmarshling json ", err)
	}

	f, _ := strconv.ParseFloat(decode[0].PriceUsd, 64)

	bigF := big.NewFloat(f)
	bigFloatString := bigF.String()
	var s string
	for _, v := range bigFloatString {
		if string(v) == "." {
			break
		}
		s += string(v)
	}
	i, _ := strconv.ParseInt(s, 10, 64)
	return big.NewInt(i)
}
