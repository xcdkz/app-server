package coins

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type tickersTickersMarketStruct struct {
	Name                  *string `json:"name"`
	Identifier            *string `json:"identifier"`
	Has_trading_incentive bool    `json:"has_trading_incentive"`
	Logo                  *string `json:"logo"`
}

type tickersTickersStruct struct {
	Base                      *string                     `json:"base"`
	Target                    *string                     `json:"target"`
	Market                    *tickersTickersMarketStruct `json:"market"`
	Last                      *float64                    `json:"last"`
	Volume                    *float64                    `json:"volume"`
	Cost_to_move_up_usd       *float64                    `json:"cost_to_move_up_usd"`
	Cost_to_move_down_usd     *float64                    `json:"cost_to_move_down_usd"`
	Converted_last            map[string]float64          `json:"converted_last"`
	Converted_volume          map[string]float64          `json:"converted_volume"`
	Trust_score               *string                     `json:"trust_score"`
	Bid_ask_spread_percentage *float64                    `json:"bid_ask_spread_percentage"`
	Timestamp                 *string                     `json:"timestamp"`
	Last_traded_at            *string                     `json:"last_traded_at"`
	Last_fetch_at             *string                     `json:"last_fetch_at"`
	Is_anomaly                *bool                       `json:"is_anomaly"`
	Is_stale                  *bool                       `json:"is_stale"`
	Trade_url                 *string                     `json:"trade_url"`
	Token_info_url            *string                     `json:"token_info_url"`
	Coin_id                   *string                     `json:"coin_id"`
	Target_coin_id            *string                     `json:"target_coin_id"`
}

type tickersStruct struct {
	Name    *string                `json:"name"`
	Tickers []tickersTickersStruct `json:"tickers"`
}

func Tickers(
	id string,
	exchange_ids string,
	include_exchange_logo bool,
	pageInt string,
	order string,
	depth bool,
	dex_pair_format string,
) (*http.Response, error) {
	urlAddress := ""
	if include_exchange_logo {
		urlAddress += "include_exchange_logo=true"
	} else {
		urlAddress += "include_exchange_logo=false"
	}
	if len(exchange_ids) > 0 {
		urlAddress += "&exchange_ids=" + exchange_ids
	}
	if len(pageInt) > 0 {
		urlAddress += "&page=" + pageInt
	}
	if len(order) > 0 {
		urlAddress += "&order=" + order
	}
	if depth {
		urlAddress += "&depth=true"
	}
	if len(dex_pair_format) > 0 {
		urlAddress += "&dex_pair_format=" + dex_pair_format
	}
	urlAddress = CoingeckoAPIURL + "/" + id + "/tickers?" + url.PathEscape(urlAddress)
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

func TickersStruct(
	id string,
	exchange_ids string,
	include_exchange_logo bool,
	pageInt string,
	order string,
	depth bool,
	dex_pair_format string,
) (*tickersStruct, error) {
	tickers, err := Tickers(id, exchange_ids, include_exchange_logo, pageInt, order, depth, dex_pair_format)
	if err != nil {
		return nil, err
	}
	defer tickers.Body.Close()
	if tickers.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK, " + tickers.Status)
	}
	bodyTypes, err := io.ReadAll(tickers.Body)
	if err != nil {
		return nil, err
	}
	var t tickersStruct
	if err = json.Unmarshal(bodyTypes, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
