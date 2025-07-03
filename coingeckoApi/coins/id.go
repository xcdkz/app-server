package coins

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type linksStruct struct {
	Homepage                      []string            `json:"homepage"`
	Whitepaper                    *string             `json:"whitepaper"`
	Blockchain_site               []string            `json:"blockchain_site"`
	Official_forum_url            []string            `json:"official_forum_url"`
	Chat_url                      []string            `json:"chat_url"`
	Annoucement_url               []string            `json:"annoucement_url"`
	Snapshot_url                  *string             `json:"snapshot_url"`
	Twitter_screen_name           *string             `json:"twitter_screen_name"`
	Facebook_username             *string             `json:"facebook_username"`
	Bitcointalk_thread_identifier *string             `json:"bitcointalk_thread_identifier"`
	Telegram_channel_identifier   *string             `json:"telegram_channel_identifier"`
	Subreddit_url                 *string             `json:"subreddit_url"`
	Repos_url                     map[string][]string `json:"repos_url"`
}

type marketDataStruct struct {
	Current_price                                map[string]float64 `json:"current_price"`
	Total_value_locked                           *float64           `json:"total_value_locked"`
	Mcap_to_tvl_ratio                            *float64           `json:"mcap_to_tvl_ratio"`
	Fdv_to_tvl_ratio                             *float64           `json:"fdv_to_tvl_ratio"`
	Roi                                          *float64           `json:"roi"`
	Ath                                          map[string]float64 `json:"ath"`
	Ath_change_percentage                        map[string]float64 `json:"ath_change_percentage"`
	Ath_date                                     map[string]string  `json:"ath_date"`
	Atl                                          map[string]float64 `json:"atl"`
	Atl_change_percentage                        map[string]float64 `json:"atl_change_percentage"`
	Atl_date                                     map[string]string  `json:"atl_date"`
	Market_cap                                   map[string]float64 `json:"market_cap"`
	Market_cap_rank                              *int               `json:"market_cap_rank"`
	Fully_diluted_valuation                      map[string]float64 `json:"fully_diluted_valuation"`
	Market_cap_fdv_ratio                         *float64           `json:"market_cap_fdv_ratio"`
	Total_volume                                 map[string]float64 `json:"total_volume"`
	High_24h                                     map[string]float64 `json:"high_24h"`
	Low_24h                                      map[string]float64 `json:"low_24h"`
	Price_change_24h                             *float64           `json:"price_change_24h"`
	Price_change_percentage_24h                  *float64           `json:"price_change_percentage_24h"`
	Price_change_percentage_7d                   *float64           `json:"price_change_percentage_7d"`
	Price_change_percentage_14d                  *float64           `json:"price_change_percentage_14d"`
	Price_change_percentage_30d                  *float64           `json:"price_change_percentage_30d"`
	Price_change_percentage_60d                  *float64           `json:"price_change_percentage_60d"`
	Price_change_percentage_200d                 *float64           `json:"price_change_percentage_200d"`
	Price_change_percentage_1y                   *float64           `json:"price_change_percentage_1y"`
	Market_cap_change_24h                        *float64           `json:"market_cap_change_24h"`
	Market_cap_change_percentage_24h             *float64           `json:"market_cap_change_percentage_24h"`
	Price_change_percentage_1h_in_currency       map[string]float64 `json:"price_change_percentage_1h_in_currency"`
	Price_change_percentage_24h_in_currency      map[string]float64 `json:"price_change_percentage_24h_in_currency"`
	Price_change_percentage_7d_in_currency       map[string]float64 `json:"price_change_percentage_7d_in_currency"`
	Price_change_percentage_14d_in_currency      map[string]float64 `json:"price_change_percentage_14d_in_currency"`
	Price_change_percentage_30d_in_currency      map[string]float64 `json:"price_change_percentage_30d_in_currency"`
	Price_change_percentage_60d_in_currency      map[string]float64 `json:"price_change_percentage_60d_in_currency"`
	Price_change_percentage_200d_in_currency     map[string]float64 `json:"price_change_percentage_200d_in_currency"`
	Price_change_percentage_1y_in_currency       map[string]float64 `json:"price_change_percentage_1y_in_currency"`
	Market_cap_change_24h_in_currency            map[string]float64 `json:"market_cap_change_24h_in_currency"`
	Market_cap_change_percentage_24h_in_currency map[string]float64 `json:"market_cap_change_percentage_24h_in_currency"`
	Total_supply                                 *float64           `json:"total_supply"`
	Max_supply                                   *float64           `json:"max_supply"`
	Circulating_supply                           *float64           `json:"circulating_supply"`
	Last_updated                                 *string            `json:"last_updated"`
}

type developerDataStruct struct {
	Forks                               *int           `json:"forks"`
	Stars                               *int           `json:"stars"`
	Subscribers                         *int           `json:"subscribers"`
	Total_issues                        *int           `json:"total_issues"`
	Closed_issues                       *int           `json:"closed_issues"`
	Pull_requests_merged                *int           `json:"pull_requests_merged"`
	Pull_request_contributors           *int           `json:"pull_request_contributors"`
	Core_additions_deletions_4_weeks    map[string]int `json:"core_additions_deletions_4_weeks"`
	Commit_count_4_weeks                int            `json:"commit_count_4_weeks"`
	Last_4_weeks_commit_activity_series []int          `json:"last_4_weeks_commit_activity_series"`
}
type tickersMarketStruct struct {
	Name                  *string `json:"name"`
	Identifier            *string `json:"identifier"`
	Has_trading_incentive *bool   `json:"has_trading_incentive"`
}

type tickersStruct struct {
	Base                      *string             `json:"base"`
	Target                    *string             `json:"target"`
	Market                    tickersMarketStruct `json:"market"`
	Last                      *float64            `json:"last"`
	Volume                    *float64            `json:"volume"`
	Converted_last            map[string]float64  `json:"converted_last"`
	Converted_volume          map[string]float64  `json:"converted_volume"`
	Trust_score               *string             `json:"trust_score"`
	Bid_ask_spread_percentage *float64            `json:"bid_ask_spread_percentage"`
	Timestamp                 *string             `json:"timestamp"`
	Last_traded_at            *string             `json:"last_traded_at"`
	Last_fetch_at             *string             `json:"last_fetch_at"`
	Is_anomaly                *bool               `json:"is_anomaly"`
	Is_stale                  *bool               `json:"is_stale"`
	Trade_url                 *string             `json:"trade_url"`
	Token_info_url            *string             `json:"token_info_url"`
	Coin_id                   *string             `json:"coin_id"`
	Target_coin_id            *string             `json:"target_coin_id"`
}

type idStruct struct {
	Id                              *string                      `json:"id"`
	Symbol                          *string                      `json:"symbol"`
	Name                            *string                      `json:"name"`
	WebSlug                         *string                      `json:"web_slug"`
	Asset_platform_id               *string                      `json:"asset_platform_id"`
	Plaftorms                       map[string]string            `json:"plaftorms"`
	Detail_platforms                map[string]map[string]string `json:"detail_platforms"`
	Block_time_in_minutes           *float64                     `json:"block_time_in_minutes"`
	Hashing_algorithm               *string                      `json:"hashing_algorithm"`
	Categories                      []string                     `json:"categories"`
	Preview_listing                 bool                         `json:"preview_listing"`
	Public_notice                   *string                      `json:"public_notice"`
	Additional_notices              []string                     `json:"additional_notices"`
	Localization                    map[string]string            `json:"localization"`
	Description                     map[string]string            `json:"description"`
	Links                           linksStruct                  `json:"links"`
	Image                           map[string]string            `json:"image"`
	Country_origin                  *string                      `json:"country_origin"`
	Genesis_date                    *string                      `json:"genesis_date"`
	Sentiment_votes_up_percentage   float64                      `json:"sentiment_votes_up_percentage"`
	Sentiment_votes_down_percentage float64                      `json:"sentiment_votes_down_percentage"`
	Watchlist_portfolio_users       int                          `json:"watchlist_portfolio_users"`
	Market_cap_rank                 int                          `json:"market_cap_rank"`
	Market_data                     marketDataStruct             `json:"market_data"`
	Community_data                  map[string]float64           `json:"community_data"`
	Developer_data                  developerDataStruct          `json:"developer_data"`
	Status_updates                  []string                     `json:"status_updates"`
	Last_updated                    *string                      `json:"last_updated"`
	Tickers                         []tickersStruct              `json:"tickers"`
}

func Id(
	id string,
	localization bool,
	tickers bool,
	market_data bool,
	community_data bool,
	developer_data bool,
	sparkline bool,
	dex_pair_format string,
) (*http.Response, error) {
	urlAddress := ""
	if localization {
		urlAddress += "localization=true"
	} else {
		urlAddress += "localization=false"
	}
	if tickers {
		urlAddress += "&tickers=true"
	}
	if market_data {
		urlAddress += "&market_data=true"
	}
	if community_data {
		urlAddress += "&community_data=true"
	}
	if developer_data {
		urlAddress += "&developer_data=true"
	}
	if sparkline {
		urlAddress += "&sparkline=true"
	}
	if len(dex_pair_format) > 0 {
		urlAddress += "&dex_pair_format=" + dex_pair_format
	}
	urlAddress = CoingeckoAPIURL + "/" + id + "?" + url.PathEscape(urlAddress)
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

func IdStruct(
	id string,
	localization bool,
	tickers bool,
	market_data bool,
	community_data bool,
	developer_data bool,
	sparkline bool,
	dex_pair_format string,
) (*idStruct, error) {
	identifier, err := Id(id, localization, tickers, market_data, community_data, developer_data, sparkline, dex_pair_format)
	if err != nil {
		return nil, err
	}
	defer identifier.Body.Close()
	if identifier.StatusCode != http.StatusOK {
		return nil, errors.New("status code not OK, " + identifier.Status)
	}
	bodyBytes, err := io.ReadAll(identifier.Body)
	if err != nil {
		return nil, err
	}
	var i idStruct
	if err = json.Unmarshal(bodyBytes, &i); err != nil {
		return nil, err
	}
	return &i, nil
}
