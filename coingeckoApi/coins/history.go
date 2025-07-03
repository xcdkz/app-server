package coins

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type historyDeveloperDataStruct struct {
	Forks                            int            `json:"forks"`
	Stars                            int            `json:"stars"`
	Subscribers                      int            `json:"subscribers"`
	Total_issues                     int            `json:"total_issues"`
	Closed_issues                    int            `json:"closed_issues"`
	Pull_request_merged              int            `json:"pull_request_merged"`
	Pull_request_contributors        int            `json:"pull_request_contributors"`
	Code_additions_deletions_4_weeks map[string]int `json:"code_additions_deletions_4_weeks"`
	Commit_count_4_weeks             int            `json:"commit_count_4_weeks"`
}

type historyStruct struct {
	Id                    string                        `json:"id"`
	Symbol                string                        `json:"symbol"`
	Name                  string                        `json:"name"`
	Localization          map[string]string             `json:"localization"`
	Image                 map[string]string             `json:"image"`
	Market_data           map[string]map[string]float64 `json:"market_data"`
	Community_data        map[string]float64            `json:"community_data"`
	Developer_data        historyDeveloperDataStruct    `json:"developer_data"`
	Public_interest_stats map[string]int                `json:"public_interest_stats"`
}

func History(
	id string,
	date string,
	localization bool,
) (*http.Response, error) {
	urlAddress := "date=" + date
	if localization {
		urlAddress += "&localization=true"
	}
	urlAddress = CoingeckoAPIURL + "/" + id + "/history?" + url.PathEscape(urlAddress)
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

func HistoryStruct(
	id string,
	date string,
	localization bool,
) (*historyStruct, error) {
	history, err := History(id, date, localization)
	if err != nil {
		return nil, err
	}
	defer history.Body.Close()
	if history.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK, " + history.Status)
	}
	bodyTypes, err := io.ReadAll(history.Body)
	if err != nil {
		return nil, err
	}
	var h historyStruct
	if err = json.Unmarshal(bodyTypes, &h); err != nil {
		return nil, err
	}
	return &h, nil
}
