package main

import (
	"app-server/coingeckoApi/simple"
	"encoding/json"
)

const CoingeckoTerminalBaseURL = "https://api.geckoterminal.com/api/v2"

func strPtr(s string) *string { return &s }

func main() {
	/* COINGECKO TERMINAL NETWORKS TEST*/
	// nl, err := coingeckoterminalapi.NetworksList()
	// if err != nil {
	// 	panic(err)
	// }
	// t, err := json.MarshalIndent(*nl, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(t))

	/* PING TEST */
	// ping, err := coingeckoapi.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// defer ping.Body.Close()
	// println(ping.StatusCode)
	//
	sp, err := simple.PriceMaps(
		"pln,eur,usd",
		"bitcoin,ethereum",
		"Bitcoin,Ethereum",
		"btc,eth",
		"",
		false,
		true,
		false,
		false,
		"5")
	if err != nil {
		panic(err)
	}
	t, err := json.MarshalIndent(*sp, "", " ")
	if err != nil {
		panic(err)
	}
	println(string(t))

	tpm, err := simple.TokenPriceMap(
		"ethereum",
		"0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
		"pln",
		true,
		true,
		true,
		true,
		"5")
	if err != nil {
		panic(err)
	}
	t, err = json.MarshalIndent(*tpm, "", " ")
	if err != nil {
		panic(err)
	}
	println(string(t))

	svca, err := simple.SupportedVsCurrenciesArray()
	if err != nil {
		panic(err)
	}
	t, err = json.MarshalIndent(&svca, "", " ")
	if err != nil {
		panic(err)
	}
	println(string(t))
}
