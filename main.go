package main

import (
	coingeckoapi "app-server/coingeckoApi"
	coingeckoterminalapi "app-server/coingeckoTerminalApi"
	"encoding/json"
)

const CoingeckoTerminalBaseURL = "https://api.geckoterminal.com/api/v2"

func main() {
	nl, err := coingeckoterminalapi.NetworksList()
	if err != nil {
		panic(err)
	}
	t, err := json.MarshalIndent(*nl, "", " ")
	if err != nil {
		panic(err)
	}
	println(string(t))

	ping, err := coingeckoapi.Ping()
	if err != nil {
		panic(err)
	}
	defer ping.Body.Close()
	println(ping.StatusCode)
}
