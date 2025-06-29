package main

import "app-server/coingeckoApi/simple"

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
	simple.Price(
		"usd,eur",
		"bitcoin,ethereum",
		"",
		"btc,eth",
		"",
		false,
		true,
		false,
		false,
		"5")
}
