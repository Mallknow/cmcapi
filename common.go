package cmcapi

import (
	"bytes"
	"compress/gzip"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_32 = 32
	BIT_SIZE_64 = 64
)

var NIL_REQBODY = []byte{}

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

func GetPointer[T any](v T) *T {
	return &v
}

func Request(url string, reqBody []byte, method string, isGzip bool) ([]byte, error) {
	return RequestWithHeader(url, reqBody, method, map[string]string{}, isGzip)
}

func RequestWithHeader(url string, reqBody []byte, method string, headerMap map[string]string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accepts", "application/json")

	for k, v := range headerMap {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}

	req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	log.Debug("reqURL: ", req.URL.String())
	if reqBody != nil && len(reqBody) > 0 {
		log.Debug("reqBody: ", string(reqBody))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}

	data, err := io.ReadAll(body)
	return data, err
}

type Cmc struct{}

const (
	CMC_API_HTTP      = "pro-api.coinmarketcap.com"
	CMC_TEST_API_HTTP = "sandbox-api.coinmarketcap.com"
	IS_GZIP           = true
)

var NowNetType NetType = MAIN_NET

type NetType int

const (
	MAIN_NET NetType = iota
	TEST_NET
)

func SetNetType(netType NetType) {
	NowNetType = netType
}

type APIType int

const (
	REST APIType = iota
)

type Client struct {
	APIKey string
}

type CmcRestClient struct {
	c *Client
}

func (*Cmc) NewRestClient(apiKey string) *CmcRestClient {
	return &CmcRestClient{
		c: &Client{
			APIKey: apiKey,
		},
	}
}

// cmcCallAPI is a common function to call CMC API
func cmcCallAPI[T any](client *Client, url url.URL, reqBody []byte, method string) (*CmcRestRes[T], error) {
	body, err := Request(url.String(), reqBody, method, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// cmcCallAPIWithSecret is a common function to call CMC API with secret key
func cmcCallAPIWithSecret[T any](client *Client, url url.URL, reqBody []byte, method string) (*CmcRestRes[T], error) {
	body, err := RequestWithHeader(url.String(), reqBody, method, map[string]string{
		"X-CMC_PRO_API_KEY": client.APIKey,
	}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

func cmcGetRestHostByNetType(netType NetType) string {
	switch netType {
	case MAIN_NET:
		return CMC_API_HTTP
	case TEST_NET:
		return CMC_TEST_API_HTTP
	}
	return ""
}

func cmcHandlerRequestAPIWithPathQueryParam[T any](netType NetType, request *T, name string) url.URL {
	query := cmcHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     cmcGetRestHostByNetType(netType),
		Path:     name,
		RawQuery: query,
	}
	return u
}

func cmcHandlerReq[T any](req *T) string {
	var argBuffer bytes.Buffer

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		argName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			argBuffer.WriteString(argName + "=" + v.Field(i).Elem().String() + "&")
		case reflect.Int, reflect.Int64:
			argBuffer.WriteString(argName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			argBuffer.WriteString(argName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			argBuffer.WriteString(argName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			args := make([]reflect.Value, 0)
			result := ToStringMethod.Call(args)
			argBuffer.WriteString(argName + "=" + result[0].String() + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			argBuffer.WriteString(argName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(argBuffer.String(), "&")
}
