package coingeckoapi

import "net/http"

func Ping() (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", CoingeckoAPIURL+"/ping", nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"x-cg-demo-api-key": {CoingeckoAPIKey},
	}
	return client.Do(req)
}
