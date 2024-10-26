package cmcapi

type CmcRestApi int

const (
	// CryptoCurrency
	RestApiCryptocurrencyListingsLatest CmcRestApi = iota // GET 列出所有加密货币清单
	// Tools
	RestApiFearAndGreedLatest CmcRestApi = iota //获取恐惧和贪婪指数最新值
)

var RestApiMap = map[CmcRestApi]string{
	// CryptoCurrency
	RestApiCryptocurrencyListingsLatest: "/v1/cryptocurrency/listings/latest", // GET 列出所有加密货币清单
	// Tools
	RestApiFearAndGreedLatest: "/v3/fear-and-greed/latest", // GET 获取恐惧和贪婪指数最新值
}
