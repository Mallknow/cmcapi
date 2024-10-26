package cmcapi

import "fmt"

type CmcStatusRes struct {
	Status struct {
		Timestamp    string `json:"timestamp"`
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		Elapsed      int    `json:"elapsed"`
		CreditCount  int    `json:"credit_count"`
	} `json:"status"`
}

type CmcRestRes[T any] struct {
	Data         T `json:"data"` // 请求结果
	CmcStatusRes   // 状态信息
}

func handlerCommonRest[T any](data []byte) (*CmcRestRes[T], error) {
	res := &CmcRestRes[T]{}
	log.Warn(string(data))
	err := json.Unmarshal(data, &res)
	if err != nil {
		log.Error("rest返回值获取失败", err)
	}
	return res, err
}
func (status *CmcStatusRes) handlerError() error {
	if status.Status.ErrorCode != 0 {
		return fmt.Errorf("request error:[code:%v] message: %v", status.Status.ErrorCode, status.Status.ErrorMessage)
	}
	return nil
}
