package simple

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func SupportedVsCurrencies() (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", CoingeckoAPIURL+"/supported_vs_currencies", nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"x-cg-demo-api-key": {CoingeckoAPIKey},
	}
	return client.Do(req)
}

func SupportedVsCurrenciesArray() (*[]string, error) {
	svca, err := SupportedVsCurrencies()
	if err != nil {
		return nil, err
	}
	defer svca.Body.Close()
	if svca.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK")
	}
	bodyBytes, err := io.ReadAll(svca.Body)
	if err != nil {
		return nil, err
	}
	var supportedVsCurrenciesArray []string
	if err = json.Unmarshal(bodyBytes, &supportedVsCurrenciesArray); err != nil {
		return nil, err
	}
	return &supportedVsCurrenciesArray, nil
}
