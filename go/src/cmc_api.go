package main

import (
	"fmt"
	//"strings"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)
/*
	
id	"ethereum"
name	"Ethereum"
symbol	"ETH"
rank	"2"
price_usd	"411.063"
price_btc	"0.0601683"
24h_volume_usd	"1157210000.0"
market_cap_usd	"40588180291.0"
available_supply	"98739561.0"
total_supply	"98739561.0"
max_supply	null
percent_change_1h	"0.39"
percent_change_24h	"4.27"
percent_change_7d	"-0.16"
last_updated	"1523395454"*/

type Response struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Rank string `json:"rank"`
	PriceUsd string `json:"price_usd"`
	PriceBtc string `json:"price_btc"`
	TwentyFourHrVolume string `json:"24h_volume_usd"`
	MarketCapUsd string `json:"market_cap_usd"`
	AvailableSupply string `json:"available_supply"`
	TotalSupply string `json:"total_supply"`
	MaxSupply string `json:"null"`
	PercentChange1h string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d string `json:"percent_change_7d"`
	LastUpdate string `json:"last_updated"`
}

var responseObject Response

func main() {
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
	fmt.Println(decode)
}
/*
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }

    s, err := getStations([]byte(body))
func getStations(body []byte) (*StationAPIResponse, error) {
    var s = new(StationAPIResponse)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}
*/
