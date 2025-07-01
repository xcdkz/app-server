package simple

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func TokenPrice(
	id string,
	contract_addresses string,
	vs_currencies string,
	include_market_cap bool,
	include_24hr_vol bool,
	include_24hr_change bool,
	include_last_updated_at bool,
	precision string,
) (*http.Response, error) {
	urlString := "contract_addresses=" + contract_addresses + "&vs_currencies=" + vs_currencies
	if include_market_cap {
		urlString += "&include_market_cap=true"
	}
	if include_24hr_vol {
		urlString += "&include_24hr_vol=true"
	}
	if include_24hr_change {
		urlString += "&include_24hr_change=true"
	}
	if include_last_updated_at {
		urlString += "&include_last_updated_at=true"
	}
	if len(precision) != 0 {
		urlString += "&precision=" + precision
	}
	urlString = CoingeckoAPIURL + "/token_price/" + id + "?" + url.PathEscape(urlString)
	client := http.Client{}
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"x-cg-demo-api-key": {CoingeckoAPIKey},
	}
	return client.Do(req)
}

func TokenPriceMap(
	id string,
	contract_addresses string,
	vs_currencies string,
	include_market_cap bool,
	include_24hr_vol bool,
	include_24hr_change bool,
	include_last_updated_at bool,
	precision string,
) (*map[string]map[string]float64, error) {
	tpm, err := TokenPrice(id, contract_addresses, vs_currencies, include_market_cap, include_24hr_vol, include_24hr_change, include_last_updated_at, precision)
	if err != nil {
		return nil, err
	}
	defer tpm.Body.Close()
	if tpm.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK")
	}
	bodyBytes, err := io.ReadAll(tpm.Body)
	if err != nil {
		return nil, err
	}
	var tokenPriceMap map[string]map[string]float64
	if err = json.Unmarshal(bodyBytes, &tokenPriceMap); err != nil {
		return nil, err
	}
	return &tokenPriceMap, nil
}
