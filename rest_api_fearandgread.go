package cmcapi

// Tools

// cmc RestApiFearAndGreedLatest REST接口 GET 获取恐惧和贪婪指数最新值
func (client *CmcRestClient) NewCmcRestApiFearAndGreedLatest() *CmcRestApiFearAndGreedLatestAPI {
	return &CmcRestApiFearAndGreedLatestAPI{
		client: client,
		req:    &CmcRestApiFearAndGreedLatestReq{},
	}
}

func (api *CmcRestApiFearAndGreedLatestAPI) Do() (*CmcRestRes[CmcRestApiFearAndGreedLatestRes], error) {
	url := cmcHandlerRequestAPIWithPathQueryParam(NowNetType, api.req, RestApiMap[RestApiFearAndGreedLatest])
	return cmcCallAPIWithSecret[CmcRestApiFearAndGreedLatestRes](api.client.c, url, NIL_REQBODY, GET)
}
