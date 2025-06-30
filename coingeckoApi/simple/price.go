package simple

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func Price(
	vs_currencies string,
	ids string,
	names string,
	symbols string,
	include_tokens string,
	include_market_cap bool,
	include_24hr_vol bool,
	include_24hr_change bool,
	include_last_updated_at bool,
	precision string,
) (*http.Response, error) {
	urlString := "vs_currencies=" + vs_currencies
	if len(ids) != 0 {
		urlString += "&ids=" + ids
	}
	if len(names) != 0 {
		urlString += "&names=" + names
	}
	if len(symbols) != 0 {
		urlString += "&symbols=" + symbols
	}
	if len(include_tokens) != 0 {
		urlString += "&include_tokens=" + include_tokens
	}
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
	urlString = CoingeckoAPIURL + "/price?" + url.PathEscape(urlString)
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

func PriceMaps(
	vs_currencies string,
	ids string,
	names string,
	symbols string,
	include_tokens string,
	include_market_cap bool,
	include_24hr_vol bool,
	include_24hr_change bool,
	include_last_updated_at bool,
	precision string,
) (*map[string]map[string]float64, error) {
	pm, err := Price(vs_currencies, ids, names, symbols, include_tokens, include_market_cap, include_24hr_vol, include_24hr_change, include_last_updated_at, precision)
	if err != nil {
		return nil, err
	}
	defer pm.Body.Close()
	if pm.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK")
	}
	bodyBytes, err := io.ReadAll(pm.Body)
	if err != nil {
		return nil, err
	}
	var pricemaps map[string]map[string]float64
	if err = json.Unmarshal(bodyBytes, &pricemaps); err != nil {
		return nil, err
	}
	return &pricemaps, nil
}
