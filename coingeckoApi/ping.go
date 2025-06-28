package coingeckoapi

import "net/http"

func Ping() (*http.Response, error) {
	return http.Get(CoingeckoAPIURL + "/ping" + "?x_cg_demo_api_key=" + CoingeckoAPIKey)
}
