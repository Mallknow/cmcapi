package cmcapi

// CryptoCurrency
// cmc RestApiCryptocurrencyListingsLatest REST接口 GET 列出所有加密货币清单
func (client *CmcRestClient) NewCmcRestApiCryptocurrencyListingsLatest() *CmcRestApiCryptocurrencyListingsLatestAPI {
	return &CmcRestApiCryptocurrencyListingsLatestAPI{
		client: client,
		req:    &CmcRestApiCryptocurrencyListingsLatestReq{},
	}
}

func (api *CmcRestApiCryptocurrencyListingsLatestAPI) Do() (*CmcRestRes[CmcRestApiCryptocurrencyListingsLatestRes], error) {
	url := cmcHandlerRequestAPIWithPathQueryParam(NowNetType, api.req, RestApiMap[RestApiCryptocurrencyListingsLatest])
	return cmcCallAPIWithSecret[CmcRestApiCryptocurrencyListingsLatestRes](api.client.c, url, NIL_REQBODY, GET)
}

// cmc RestApiCryptocurrencyQuotesLatest REST接口 GET 获取最新的加密货币报价
func (client *CmcRestClient) NewCmcRestApiCryptocurrencyQuotesLatest() *CmcRestApiCryptocurrencyQuotesLatestAPI {
	return &CmcRestApiCryptocurrencyQuotesLatestAPI{
		client: client,
		req:    &CmcRestApiCryptocurrencyQuotesLatestReq{},
	}
}

func (api *CmcRestApiCryptocurrencyQuotesLatestAPI) Do() (*CmcRestRes[CmcRestApiCryptocurrencyQuotesLatestRes], error) {
	url := cmcHandlerRequestAPIWithPathQueryParam(NowNetType, api.req, RestApiMap[RestApiCryptocurrencyQuotesLatest])
	return cmcCallAPIWithSecret[CmcRestApiCryptocurrencyQuotesLatestRes](api.client.c, url, NIL_REQBODY, GET)
}
