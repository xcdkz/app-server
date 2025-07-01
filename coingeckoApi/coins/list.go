package coins

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type coinsList struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms"`
}

func List(include_platform bool) (*http.Response, error) {
	urlAddress := CoingeckoAPIURL + "/list"
	if include_platform {
		urlAddress += "?include_platform=true"
	}
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

func ListStructs(include_platform bool) (*[]coinsList, error) {
	list, err := List(include_platform)
	if err != nil {
		return nil, err
	}
	defer list.Body.Close()
	if list.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK, " + list.Status)
	}
	bodyBytes, err := io.ReadAll(list.Body)
	if err != nil {
		return nil, err
	}
	var ls []coinsList
	if err = json.Unmarshal(bodyBytes, &ls); err != nil {
		return nil, err
	}
	return &ls, nil
}
