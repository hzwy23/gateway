package filter

import (
	"github.com/wisrc/gateway/core/context"
	"sync"

	"github.com/wisrc/gateway/logger"
)

type BeforeResponse func(ctx *context.GatewayContext)  error

var beforeResponseFunc []BeforeResponse
var beforeResponseLock = &sync.RWMutex{}

func RegisterBeforeResponse(handle BeforeResponse, filterName string) {
	beforeResponseLock.Lock()
	defer beforeResponseLock.Unlock()
	logger.Info("注册响应前过滤器，过滤器名称是：", filterName)
	beforeResponseFunc = append(beforeResponseFunc, handle)
}

func BeforeResponseFilter(ctx *context.GatewayContext) error {
	beforeResponseLock.RLock()
	defer beforeResponseLock.RUnlock()

	for _, handle := range beforeResponseFunc {
		err := handle(ctx)
		if err != nil {
			logger.Error(err)
			return  err
		}
	}
	return nil
}

func init() {

	//RegisterBeforeResponse(func(ctx *context.GatewayContext)  error {
	//	// 读取返回结果
	//	body, err := ioutil.ReadAll(ctx.Response.Body)
	//	if err != nil {
	//		logger.Error(err)
	//		return response.NewError(response.ParseHttpResponseFailed, err.Error())
	//	}
	//	ctx.Body = body
	//
	//	ctx.Response.Body = ioutil.NopCloser(bytes.NewReader(body))
	//	return errors.New("demo")
	//}, "读取响应值")

	RegisterBeforeResponse(func(ctx *context.GatewayContext) error {
		// 读取返回结果
		logger.Info("请求地址是：",ctx.RemoteURL.String(),", 响应状态吗：", ctx.Response.StatusCode)
		return nil
	}, "读取 http 状态吗")

}
