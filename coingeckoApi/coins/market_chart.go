package coins

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type marketChartStruct struct {
	Prices        [][]float64 `json:"prices"`
	Market_caps   [][]float64 `json:"market_caps"`
	Total_volumes [][]float64 `json:"total_volumes"`
}

func MarketChart(
	id string,
	vs_currency string,
	days string,
	interval string,
	precision string,
) (*http.Response, error) {
	urlAddress := "vs_currency=" + vs_currency + "&days=" + days
	if len(interval) > 0 {
		urlAddress += "&interval=" + interval
	}
	if len(precision) > 0 {
		urlAddress += "&precision=" + precision
	}
	urlAddress = CoingeckoAPIURL + "/" + id + "/market_chart?" + url.PathEscape(urlAddress)
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

func MarketChartStruct(
	id string,
	vs_currency string,
	days string,
	interval string,
	precision string,
) (*marketChartStruct, error) {
	marketChart, err := MarketChart(id, vs_currency, days, interval, precision)
	if err != nil {
		return nil, err
	}
	defer marketChart.Body.Close()
	if marketChart.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK, " + marketChart.Status)
	}
	bodyTypes, err := io.ReadAll(marketChart.Body)
	if err != nil {
		return nil, err
	}
	var mc marketChartStruct
	if err = json.Unmarshal(bodyTypes, &mc); err != nil {
		return nil, err
	}

	return &mc, nil
}
