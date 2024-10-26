package cmcapi

import (
	"fmt"
)

type CmcStatusRes struct {
	Status struct {
		Timestamp    string      `json:"timestamp"`
		ErrorCode    interface{} `json:"error_code"`
		ErrorMessage string      `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
	} `json:"status"`
}

type CmcRestRes[T any] struct {
	CmcStatusRes   // 状态信息
	Data         T `json:"data"` // 请求结果
}

func handlerCommonRest[T any](data []byte) (*CmcRestRes[T], error) {
	res := &CmcRestRes[T]{}
	log.Warn(string(data))
	err := json.Unmarshal(data, &res)
	if err != nil {
		log.Error("rest返回值获取失败: ", err)
	}
	return res, err
}
func (s *CmcStatusRes) handlerError() error {
	// API Response Bug: Status ErrorCode could be int or string
	errorCode := fmt.Sprintf("%v", s.Status.ErrorCode)
	if errorCode != "0" {
		return fmt.Errorf("request error:[code:%v] message: %v", s.Status.ErrorCode, s.Status.ErrorMessage)
	}
	return nil
}
