package cmcapi

// CryptoCurrency

type CmcRestApiCryptocurrencyListingsLatestResRow struct {
	Id                            int      `json:"id"`                               // The unique CoinMarketCap ID for this cryptocurrency.
	Name                          string   `json:"name"`                             // The name of this cryptocurrency.
	Symbol                        string   `json:"symbol"`                           // The ticker symbol for this cryptocurrency.
	Slug                          string   `json:"slug"`                             // The web URL friendly shorthand version of this cryptocurrency name.
	CmcRank                       int      `json:"cmc_rank"`                         // The cryptocurrency's CoinMarketCap rank by market cap.
	NumMarketPairs                int      `json:"num_market_pairs"`                 // The number of active trading pairs available for this cryptocurrency across supported exchanges.
	CirculatingSupply             float64  `json:"circulating_supply"`               // The approximate number of coins circulating for this cryptocurrency.
	TotalSupply                   float64  `json:"total_supply"`                     // The approximate total amount of coins in existence right now (minus any coins that have been verifiably burned).
	MarketCapByTotalSupply        float64  `json:"market_cap_by_total_supply"`       // The market cap by total supply. This field is only returned if requested through the aux request parameter.
	MaxSupply                     float64  `json:"max_supply"`                       // The expected maximum limit of coins ever to be available for this cryptocurrency.
	InfiniteSupply                bool     `json:"infinite_supply"`                  // The cryptocurrency is known to have an infinite supply.
	LastUpdated                   string   `json:"last_updated"`                     // Timestamp (ISO 8601) of the last time this cryptocurrency's market data was updated.
	DateAdded                     string   `json:"date_added"`                       // Timestamp (ISO 8601) of when this cryptocurrency was added to CoinMarketCap.
	Tags                          []string `json:"tags"`                             // Array of tags associated with this cryptocurrency. Currently only a mineable tag will be returned if the cryptocurrency is mineable. Additional tags will be returned in the future.
	SelfReportedCirculatingSupply float64  `json:"self_reported_circulating_supply"` // The self reported number of coins circulating for this cryptocurrency.
	SelfReportedMarketCap         float64  `json:"self_reported_market_cap"`         // The self reported market cap for this cryptocurrency.
	TvlRatio                      float64  `json:"tvl_ratio"`                        // Percentage of Total Value Locked
	Platform                      struct {
		Id           int    `json:"id"`            // The unique CoinMarketCap ID for this parent cryptocurrency platform.
		Name         string `json:"name"`          // The name of the parent cryptocurrency platform.
		Symbol       string `json:"symbol"`        // The ticker symbol of the parent cryptocurrency platform.
		Slug         string `json:"slug"`          // The web URL friendly shorthand version of the parent cryptocurrency platform name.
		TokenAddress string `json:"token_address"` // The token address of the parent cryptocurrency platform.
	} `json:"platform"` // Metadata about the parent cryptocurrency platform this cryptocurrency belongs to if it is a token, otherwise null.
	Quote map[string]struct {
		Price                 float64 `json:"price"`                    // Price in the specified currency for this historical.
		Volume24h             float64 `json:"volume_24h"`               // Rolling 24 hour adjusted volume in the specified currency.
		VolumeChange24h       float64 `json:"volume_change_24h"`        // 24 hour change in the specified currencies volume.
		Volume24hReported     float64 `json:"volume_24h_reported"`      // Rolling 24 hour reported volume in the specified currency. This field is only returned if requested through the aux request parameter.
		Volume7d              float64 `json:"volume_7d"`                // Rolling 7 day adjusted volume in the specified currency. This field is only returned if requested through the aux request parameter.
		Volume7dReported      float64 `json:"volume_7d_reported"`       // Rolling 7 day reported volume in the specified currency. This field is only returned if requested through the aux request parameter.
		Volume30d             float64 `json:"volume_30d"`               // Rolling 30 day adjusted volume in the specified currency. This field is only returned if requested through the aux request parameter.
		Volume30dReported     float64 `json:"volume_30d_reported"`      // Rolling 30 day reported volume in the specified currency. This field is only returned if requested through the aux request parameter.
		MarketCap             float64 `json:"market_cap"`               // Market cap in the specified currency.
		MarketCapDominance    float64 `json:"market_cap_dominance"`     // Market cap dominance in the specified currency.
		FullyDilutedMarketCap float64 `json:"fully_diluted_market_cap"` // Fully diluted market cap in the specified currency.
		Tvl                   float64 `json:"tvl"`                      // Total Value Locked
		PercentChange1h       float64 `json:"percent_change_1h"`        // 1 hour change in the specified currency.
		PercentChange24h      float64 `json:"percent_change_24h"`       // 24 hour change in the specified currency.
		PercentChange7d       float64 `json:"percent_change_7d"`        // 7 day change in the specified currency.
		LastUpdated           string  `json:"last_updated"`             // Timestamp (ISO 8601) of when the conversion currency's current value was referenced.
	} `json:"quote"` // A map of market quotes in different currency conversions. The default map included is USD.

}
type CmcRestApiCryptocurrencyListingsLatestRes []CmcRestApiCryptocurrencyListingsLatestResRow
