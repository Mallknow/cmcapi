package cmcapi

// CryptoCurrency

type CmcRestApiCryptocurrencyListingsLatestReq struct {
	Start                *int     `json:"start"`                  //integer >= 1 1 Optionally offset the start (1-based index) of the paginated list of items to return.
	Limit                *int     `json:"limit"`                  //integer [ 1 .. 5000 ] 100 Optionally specify the number of results to return. Use this parameter and the "start" parameter to determine your own pagination size.
	PriceMin             *float64 `json:"price_min"`              //number [ 0 .. 100000000000000000 ] Optionally specify a threshold of minimum USD price to filter results by.
	PriceMax             *float64 `json:"price_max"`              //number [ 0 .. 100000000000000000 ] Optionally specify a threshold of maximum USD price to filter results by.
	MarketCapMin         *float64 `json:"market_cap_min"`         //number [ 0 .. 100000000000000000 ] Optionally specify a threshold of minimum market cap to filter results by.
	MarketCapMax         *float64 `json:"market_cap_max"`         //number [ 0 .. 100000000000000000 ] Optionally specify a threshold of maximum market cap to filter results by.
	Volume24hMin         *float64 `json:"volume_24h_min"`         //number [ 0 .. 100000000000000000 ] Optionally specify a threshold of minimum 24 hour USD volume to filter results by.
	Volume24hMax         *float64 `json:"volume_24h_max"`         //number [ 0 .. 100000000000000000 ] Optionally specify a threshold of maximum 24 hour USD volume to filter results by.
	CirculatingSupplyMin *float64 `json:"circulating_supply_min"` //number [ 0 .. 100000000000000000 ] Optionally specify a threshold of minimum circulating supply to filter results by.
	CirculatingSupplyMax *float64 `json:"circulating_supply_max"` //number [ 0 .. 100000000000000000 ] Optionally specify a threshold of maximum circulating supply to filter results by.
	PercentChange24hMin  *float64 `json:"percent_change_24h_min"` //number >= -100 Optionally specify a threshold of minimum 24 hour percent change to filter results by.
	PercentChange24hMax  *float64 `json:"percent_change_24h_max"` //number >= -100 Optionally specify a threshold of maximum 24 hour percent change to filter results by.
	Convert              *string  `json:"convert"`                //string Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found here. Each conversion is returned in its own "quote" object.
	ConvertId            *string  `json:"convert_id"`             //string Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to convert outside of ID format. Ex: convert_id=1,2781 would replace convert=BTC,USD in your query. This parameter cannot be used when convert is used.
	Sort                 *string  `json:"sort"`                   //string "market_cap" "name""symbol""date_added""market_cap""market_cap_strict""price""circulating_supply""total_supply""max_supply""num_market_pairs""volume_24h""percent_change_1h""percent_change_24h""percent_change_7d""market_cap_by_total_supply_strict""volume_7d""volume_30d" What field to sort the list of cryptocurrencies by.
	SortDir              *string  `json:"sort_dir"`               //string "asc""desc" The direction in which to order cryptocurrencies against the specified sort.
	CryptocurrencyType   *string  `json:"cryptocurrency_type"`    //string "all" "all""coins""tokens" The type of cryptocurrency to include.
	Tag                  *string  `json:"tag"`                    //string "all" "all""defi""filesharing" The tag of cryptocurrency to include.
	Aux                  *string  `json:"aux"`                    //string "num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply" Optionally specify a comma-separated list of supplemental data fields to return. Pass num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_market_cap_included_in_calc to include all auxiliary fields.
}

type CmcRestApiCryptocurrencyListingsLatestAPI struct {
	client *CmcRestClient
	req    *CmcRestApiCryptocurrencyListingsLatestReq
}

// integer >= 1 1 Optionally offset the start (1-based index) of the paginated list of items to return.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) Start(start int) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.Start = GetPointer(start)
	return api
}

// integer [ 1 .. 5000 ] 100 Optionally specify the number of results to return. Use this parameter and the "start" parameter to determine your own pagination size.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) Limit(limit int) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// number [ 0 .. 100000000000000000 ] Optionally specify a threshold of minimum USD price to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) PriceMin(priceMin float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.PriceMin = GetPointer(priceMin)
	return api
}

// number [ 0 .. 100000000000000000 ] Optionally specify a threshold of maximum USD price to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) PriceMax(priceMax float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.PriceMax = GetPointer(priceMax)
	return api
}

// number [ 0 .. 100000000000000000 ] Optionally specify a threshold of minimum market cap to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) MarketCapMin(marketCapMin float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.MarketCapMin = GetPointer(marketCapMin)
	return api
}

// number [ 0 .. 100000000000000000 ] Optionally specify a threshold of maximum market cap to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) MarketCapMax(marketCapMax float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.MarketCapMax = GetPointer(marketCapMax)
	return api
}

// number [ 0 .. 100000000000000000 ] Optionally specify a threshold of minimum 24 hour USD volume to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) Volume24hMin(volume24hMin float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.Volume24hMin = GetPointer(volume24hMin)
	return api
}

// number [ 0 .. 100000000000000000 ] Optionally specify a threshold of maximum 24 hour USD volume to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) Volume24hMax(volume24hMax float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.Volume24hMax = GetPointer(volume24hMax)
	return api
}

// number [ 0 .. 100000000000000000 ] Optionally specify a threshold of minimum circulating supply to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) CirculatingSupplyMin(circulatingSupplyMin float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.CirculatingSupplyMin = GetPointer(circulatingSupplyMin)
	return api
}

// number [ 0 .. 100000000000000000 ] Optionally specify a threshold of maximum circulating supply to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) CirculatingSupplyMax(circulatingSupplyMax float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.CirculatingSupplyMax = GetPointer(circulatingSupplyMax)
	return api
}

// number >= -100 Optionally specify a threshold of minimum 24 hour percent change to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) PercentChange24hMin(percentChange24hMin float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.PercentChange24hMin = GetPointer(percentChange24hMin)
	return api
}

// number >= -100 Optionally specify a threshold of maximum 24 hour percent change to filter results by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) PercentChange24hMax(percentChange24hMax float64) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.PercentChange24hMax = GetPointer(percentChange24hMax)
	return api
}

// string Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found here. Each conversion is returned in its own "quote" object.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) Convert(convert string) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.Convert = GetPointer(convert)
	return api
}

// string Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to convert outside of ID format. Ex: convert_id=1,2781 would replace convert=BTC,USD in your query. This parameter cannot be used when convert is used.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) ConvertId(convertId string) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.ConvertId = GetPointer(convertId)
	return api
}

// string "market_cap" "name""symbol""date_added""market_cap""market_cap_strict""price""circulating_supply""total_supply""max_supply""num_market_pairs""volume_24h""percent_change_1h""percent_change_24h""percent_change_7d""market_cap_by_total_supply_strict""volume_7d""volume_30d" What field to sort the list of cryptocurrencies by.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) Sort(sort string) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.Sort = GetPointer(sort)
	return api
}

// string "asc""desc" The direction in which to order cryptocurrencies against the specified sort.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) SortDir(sortDir string) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.SortDir = GetPointer(sortDir)
	return api
}

// string "all" "all""coins""tokens" The type of cryptocurrency to include.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) CryptocurrencyType(cryptocurrencyType string) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.CryptocurrencyType = GetPointer(cryptocurrencyType)
	return api
}

// string "all" "all""defi""filesharing" The tag of cryptocurrency to include.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) Tag(tag string) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.Tag = GetPointer(tag)
	return api
}

// string "num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply" Optionally specify a comma-separated list of supplemental data fields to return. Pass num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_market_cap_included_in_calc to include all auxiliary fields.
func (api *CmcRestApiCryptocurrencyListingsLatestAPI) Aux(aux string) *CmcRestApiCryptocurrencyListingsLatestAPI {
	api.req.Aux = GetPointer(aux)
	return api
}
