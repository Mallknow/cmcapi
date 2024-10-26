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
