package filter

import (
	"github.com/wisrc/gateway/core/context"
	"sync"

	"github.com/wisrc/gateway/logger"
)

/**
* body 相应结果
* w 相应
* r 请求
**/
type AfterResponse func(ctx *context.GatewayContext)  error

var afterResponseFunc []AfterResponse
var afterResponseLock = &sync.RWMutex{}

func RegisterAfterResponse(handle AfterResponse, filterName string) {
	afterResponseLock.Lock()
	defer afterResponseLock.Unlock()
	logger.Info("注册响应后过滤器，过滤器名称是：", filterName)
	afterResponseFunc = append(afterResponseFunc, handle)
}

func AfterResponseFilter(ctx *context.GatewayContext)  error {
	afterResponseLock.RLock()
	defer afterResponseLock.RUnlock()

	for _, handle := range afterResponseFunc {
		 err := handle(ctx)
		if err != nil {
			logger.Error(err)
			return  err
		}
	}

	return  nil
}

func init() {
	RegisterAfterResponse(func(ctx *context.GatewayContext)  error {
		logger.Info("请求处理完成: ", ctx.RemoteURL.String())
		return  nil
	}, "注册响应后处理过滤器")
}
