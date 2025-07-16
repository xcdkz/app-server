package coins

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func Ohlc(
	id string,
	vs_currency string,
	days string,
	precision string,
) (*http.Response, error) {
	urlAddress := "vs_currency=" + vs_currency + "&days=" + days
	if len(precision) > 0 {
		urlAddress += "&precision=" + precision
	}
	urlAddress = CoingeckoAPIURL + "/" + id + "/ohlc?" + url.PathEscape(urlAddress)
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

func OhlcStruct(
	id string,
	vs_currency string,
	days string,
	precision string,
) (*[][]float64, error) {
	ohlc, err := Ohlc(id, vs_currency, days, precision)
	if err != nil {
		return nil, err
	}
	defer ohlc.Body.Close()
	if ohlc.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK, " + ohlc.Status)
	}
	bodyTypes, err := io.ReadAll(ohlc.Body)
	if err != nil {
		return nil, err
	}
	var o [][]float64
	if err = json.Unmarshal(bodyTypes, &o); err != nil {
		return nil, err
	}

	return &o, nil
}
