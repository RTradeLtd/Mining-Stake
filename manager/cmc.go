package manager

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

	f, _ := strconv.ParseFloat(decode[0].PriceUsd, 64)

	return f
}
