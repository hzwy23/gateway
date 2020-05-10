package response

import (
	"encoding/json"
	"errors"
	"github.com/wisrc/gateway/core/context"
)

const (
	CreateHttpRequestFailed = 1000
	ParseHttpResponseFailed = 1001
	ProxyUrlNotFound = 1004
	ProxyError = 1502
)

var respStatus = map[int]string{
	CreateHttpRequestFailed: "创建 HTTP Request 请求失败",
	ParseHttpResponseFailed: "解析 HTTP Response Body 失败",
	ProxyUrlNotFound: "API Not Found",
	ProxyError: "Bad Gateway, HTTP 请求后台服务失败",
}

func StatusText(code int) string {
	return respStatus[code]
}

type RestResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewRestResponse 创建 http 响应对象
func NewRestResponse(code int, msg string, data interface{}) *RestResponse {
	return &RestResponse{
		code,
		msg,
		data,
	}
}

// NewError 创建标准错误信息
func NewError(code int, data interface{}) error{
	ret := &RestResponse{
		code,
		StatusText(code),
		data,
	}
	v,_ := json.Marshal(ret)
	return errors.New(string(v))
}

// errorHandle 请求失败
func ErrorHandle(ctx *context.GatewayContext, err error) {
	copyHeader(ctx)
	ctx.ResponseWriter.Write([]byte(err.Error()))
}

func copyHeader(ctx *context.GatewayContext){
	if ctx.Response != nil && ctx.Response.Header != nil {
		for k,v := range ctx.Response.Header {
			for _, val := range v {
				ctx.ResponseWriter.Header().Set(k, val)
			}
		}
	} else {
		ctx.ResponseWriter.Header().Set("Content-type","application/json;charset=UTF-8")
	}
}