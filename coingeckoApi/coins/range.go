package coins

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type rangeStruct struct {
	Prices        [][]float64 `json:"prices"`
	Market_caps   [][]float64 `json:"market_caps"`
	Total_volumes [][]float64 `json:"total_volumes"`
}

func Range(
	id string,
	vs_currency string,
	from_int string,
	to_int string,
	precision string,
) (*http.Response, error) {
	urlAddress := "vs_currency=" + vs_currency + "&from=" + from_int + "&to=" + to_int
	if len(precision) > 0 {
		urlAddress += "&precision=" + precision
	}
	urlAddress = CoingeckoAPIURL + "/" + id + "/market_chart/range?" + url.PathEscape(urlAddress)
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

func RangeStruct(
	id string,
	vs_currency string,
	from_int string,
	to_int string,
	precision string,
) (*rangeStruct, error) {
	rangeR, err := Range(id, vs_currency, from_int, to_int, precision)
	if err != nil {
		return nil, err
	}
	defer rangeR.Body.Close()
	if rangeR.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK, " + rangeR.Status)
	}
	bodyTypes, err := io.ReadAll(rangeR.Body)
	if err != nil {
		return nil, err
	}
	var r rangeStruct
	if err = json.Unmarshal(bodyTypes, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
