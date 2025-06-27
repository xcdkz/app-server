package coingeckoterminalapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type networkLinks struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Last  string `json:"last"`
}

type networkListData struct {
	Data  []networkList `json:"data"`
	Links networkLinks  `json:"links"`
}

type networkAttributes struct {
	Name                     string `json:"name"`
	CoingeckoAssetPlatformID string `json:"coingeckoAssetPlatformId"`
}

type networkList struct {
	ID          string            `json:"id"`
	NetworkType string            `json:"type"`
	Attributes  networkAttributes `json:"attributes"`
}

const CoingeckoTerminalBaseURL = "https://api.geckoterminal.com/api/v2"

func Networks() (*http.Response, error) {
	return http.Get(CoingeckoTerminalBaseURL + "/networks")
}

func NetworksList() (*networkListData, error) {
	networks, err := Networks()
	if err != nil {
		return nil, err
	}
	defer networks.Body.Close()
	if networks.StatusCode != http.StatusOK {
		return nil, errors.New("status Code not OK")
	}
	bodyBytes, err := io.ReadAll(networks.Body)
	if err != nil {
		return nil, err
	}
	var nl networkListData
	if err = json.Unmarshal(bodyBytes, &nl); err != nil {
		return nil, err
	}
	return &nl, nil
}
