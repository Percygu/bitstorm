package api

import "bitstorm/internal/pkg/constant"

// 这个通用的，不需要修改
// HttpResponse http独立请求返回结构体
type HttpResponse struct {
	Code constant.ErrCode `json:"code"`
	Msg  string           `json:"msg"`
	Data interface{}      `json:"data"`
}
