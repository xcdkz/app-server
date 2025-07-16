package main

import (
	"app-server/coingeckoApi/coins"
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
	// sp, err := simple.PriceMaps(
	// 	"pln,eur,usd",
	// 	"bitcoin,ethereum",
	// 	"Bitcoin,Ethereum",
	// 	"btc,eth",
	// 	"",
	// 	false,
	// 	true,
	// 	false,
	// 	false,
	// 	"5")
	// if err != nil {
	// 	panic(err)
	// }
	// t, err := json.MarshalIndent(*sp, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(t))
	//
	// tpm, err := simple.TokenPriceMap(
	// 	"ethereum",
	// 	"0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
	// 	"pln",
	// 	true,
	// 	true,
	// 	true,
	// 	true,
	// 	"5")
	// if err != nil {
	// 	panic(err)
	// }
	// t, err = json.MarshalIndent(*tpm, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(t))
	//
	// svca, err := simple.SupportedVsCurrenciesArray()
	// if err != nil {
	// 	panic(err)
	// }
	// t, err = json.MarshalIndent(&svca, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(t))
	//
	// cl, err := coins.ListStructs(true)
	// if err != nil {
	// 	panic(err)
	// }
	// tc, err := json.MarshalIndent(&cl, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(tc))
	//
	// markets, err := coins.MarketsStruct("pln", "bitcoin,ethereum", "", "", "", "", "", "", "", false, "", "", "")
	// if err != nil {
	// 	panic(err)
	// }
	// ms, err := json.MarshalIndent(&markets, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(ms))
	//
	// identifier, err := coins.IdStruct("wrapped-eeth", true, true, true, true, true, true, "")
	// if err != nil {
	// 	panic(err)
	// }
	// i, err := json.MarshalIndent(&identifier, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(i))

	// tickers, err := coins.TickersStruct("ethereum", "", true, "", "", true, "")
	// if err != nil {
	// 	panic(err)
	// }
	// t, err := json.MarshalIndent(&tickers, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(t))

	// history, err := coins.HistoryStruct("ethereum", "30-12-2024", true)
	// if err != nil {
	// 	panic(err)
	// }
	// h, err := json.MarshalIndent(&history, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(h))

	// marketChart, err := coins.MarketChartStruct("ethereum", "pln", "1", "daily", "5")
	// if err != nil {
	// 	panic(err)
	// }
	// mc, err := json.MarshalIndent(&marketChart, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(mc))

	rangeS, err := coins.RangeStruct("bitcoin", "pln", "1711929600", "1712275200", "5")
	if err != nil {
		panic(err)
	}
	r, err := json.MarshalIndent(&rangeS, "", " ")
	if err != nil {
		panic(err)
	}
	println(string(r))
}
