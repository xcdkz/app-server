package simple

import "net/http"

func Price(vs_currencies string,
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
	url := CoingeckoAPIURL + "/price?" + "vs_currencies=" + vs_currencies
	if len(ids) != 0 {
		url += "&ids=" + ids
	}
	if len(names) != 0 {
		url += "&names=" + names
	}
	if len(symbols) != 0 {
		url += "&symbols=" + symbols
	}
	if len(include_tokens) != 0 {
		url += "&include_tokens=" + include_tokens
	}
	if include_market_cap {
		url += "&include_market_cap=true"
	}
	if include_24hr_vol {
		url += "&include_24hr_vol=true"
	}
	if include_24hr_change {
		url += "&include_24hr_change=true"
	}
	if include_last_updated_at {
		url += "&include_last_updated_at=true"
	}
	if len(precision) != 0 {
		url += "&precision=" + precision
	}
	println(url)
	return nil, nil
}
