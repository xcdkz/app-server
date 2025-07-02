package coins

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type roiStruct struct {
	times      float64 `json:"times"`
	currency   string  `json:"currency"`
	percentage float64 `json:"percentage"`
}

type marketsStruct struct {
	Id                               string     `json:"id"`
	Symbol                           string     `json:"symbol"`
	Name                             string     `json:"name"`
	Image                            string     `json:"image"`
	Current_price                    float64    `json:"current_price"`
	Market_cap                       float64    `json:"market_cap"`
	Market_cap_rank                  float64    `json:"market_cap_rank"`
	Fully_diluted_valuation          float64    `json:"fully_diluted_valuation"`
	Total_volume                     float64    `json:"total_volume"`
	High_24h                         float64    `json:"high_24h"`
	Low_24h                          float64    `json:"low_24h"`
	Price_change_24h                 float64    `json:"price_change_24h"`
	Price_change_percentage_24h      float64    `json:"price_change_percentage_24h"`
	Market_cap_change_24h            float64    `json:"market_cap_change_24h"`
	Market_cap_change_percentage_24h float64    `json:"market_cap_change_percentage_24h"`
	Circulating_supply               float64    `json:"circulating_supply"`
	Total_supply                     float64    `json:"total_supply"`
	Max_supply                       float64    `json:"max_supply"`
	Ath                              float64    `json:"ath"`
	Ath_change_percentage            float64    `json:"ath_change_percentage"`
	Ath_date                         string     `json:"ath_date"`
	Atl                              float64    `json:"atl"`
	Atl_change_percentage            float64    `json:"atl_change_percentage"`
	Atl_date                         string     `json:"atl_date"`
	Roi                              *roiStruct `json:"roi"`
	Last_updated                     string     `json:"last_updated"`
}

func Markets(
	vs_currency string,
	ids string,
	names string,
	symbols string,
	include_tokens string,
	category string,
	order string,
	per_page_int string,
	page_int string,
	sparkline bool,
	price_change_percentage string,
	locale string,
	precision string,
) (*http.Response, error) {
	urlAddress := "vs_currency=" + vs_currency
	if len(ids) > 0 {
		urlAddress += "&ids=" + ids
	}
	if len(names) > 0 {
		urlAddress += "&names=" + names
	}
	if len(symbols) > 0 {
		urlAddress += "&symbols=" + symbols
	}
	if len(include_tokens) > 0 {
		urlAddress += "&include_tokens=" + include_tokens
	}
	if len(category) > 0 {
		urlAddress += "&category=" + category
	}
	if len(order) > 0 {
		urlAddress += "&order=" + order
	}
	if len(per_page_int) > 0 {
		urlAddress += "&per_page=" + per_page_int
	}
	if len(page_int) > 0 {
		urlAddress += "&page=" + page_int
	}
	if sparkline {
		urlAddress += "&sparkline=true"
	}
	if len(price_change_percentage) > 0 {
		urlAddress += "&price_change_percentage=" + price_change_percentage
	}
	if len(locale) > 0 {
		urlAddress += "&locale=" + locale
	}
	if len(precision) > 0 {
		urlAddress += "&precision=" + precision
	}

	urlAddress = CoingeckoAPIURL + "/markets?" + url.PathEscape(urlAddress)
	println(urlAddress)
	client := http.Client{}
	req, err := http.NewRequest("GET", urlAddress, nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"x-cg-demo-api-key": {CoingeckoAPIKey},
	}

	return client.Do(req)
}

func MarketsStruct(
	vs_currency string,
	ids string,
	names string,
	symbols string,
	include_tokens string,
	category string,
	order string,
	per_page_int string,
	page_int string,
	sparkline bool,
	price_change_percentage string,
	locale string,
	precision string,
) (*[]marketsStruct, error) {
	markets, err := Markets(vs_currency, ids, names, symbols, include_tokens, category, order, per_page_int, page_int, sparkline, price_change_percentage, locale, precision)
	if err != nil {
		return nil, err
	}
	defer markets.Body.Close()
	if markets.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK, " + markets.Status)
	}
	bodyBytes, err := io.ReadAll(markets.Body)
	if err != nil {
		return nil, err
	}
	var ms []marketsStruct
	if err = json.Unmarshal(bodyBytes, &ms); err != nil {
		return nil, err
	}
	return &ms, nil
}
