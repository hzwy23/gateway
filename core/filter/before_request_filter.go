package filter

import (
	"github.com/wisrc/gateway/core/context"
	"sync"

	"github.com/wisrc/gateway/logger"
)

type BeforeRequest func(ctx *context.GatewayContext)  error

var beforeRequestFunc []BeforeRequest
var beforeRequestLock = &sync.RWMutex{}

// RegisterBeforeRequest 注册请求前过滤器
func RegisterBeforeRequest(handle BeforeRequest, filterName string) {
	beforeRequestLock.Lock()
	defer beforeRequestLock.Unlock()
	logger.Info("注册过滤器，过滤器名称是：", filterName)
	beforeRequestFunc = append(beforeRequestFunc, handle)
}

func BeforeRequestFilter(ctx *context.GatewayContext)  error {
	beforeRequestLock.RLock()
	defer beforeRequestLock.RUnlock()

	for _, handle := range beforeRequestFunc {
		err := handle(ctx)
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	return  nil
}

func init() {
	RegisterBeforeRequest(func(ctx *context.GatewayContext)  error {
		//logger.Info("接收到请求：", ctx.Request.RequestURI)
		return nil
	}, "注册请求前过滤器")
}
