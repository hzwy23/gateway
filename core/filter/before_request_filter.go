package filter

import (
	"github.com/wisrc/gateway/core/context"
	"sync"

	"github.com/wisrc/gateway/logger"
)

var beforeRequestFunc []Handler
var beforeRequestLock = &sync.RWMutex{}

// RegisterBeforeRequest 注册请求前过滤器
func registerBeforeRequest(handle Handler) {
	beforeRequestLock.Lock()
	defer beforeRequestLock.Unlock()
	logger.Info("注册过滤器，过滤器名称是：", handle.Name)

	result := make([]Handler, len(beforeRequestFunc) + 1)

	if len(beforeRequestFunc) == 0 {
		result = append(beforeRequestFunc, handle)
	} else {
		for idx, h := range beforeRequestFunc {
			if h.Priority > handle.Priority {
				if idx == 0 {
					// 第一个元素
					f := []Handler{handle}
					result = append(f, beforeRequestFunc[0])
				} else if idx + 1 == len(beforeRequestFunc) {
					// 最后一个元素
					last := beforeRequestFunc[idx]
					v := append(beforeRequestFunc[:idx], handle)
					result = append(v, last)
				} else {
					// 中间元素
					v := append(beforeRequestFunc[:idx], handle)
					result = append(v, beforeRequestFunc[idx:]...)
				}
				break
			}
			result = append(beforeRequestFunc, handle)
		}
	}
	beforeRequestFunc = result
}

func BeforeRequestFilter(ctx *context.GatewayContext)  error {
	beforeRequestLock.RLock()
	defer beforeRequestLock.RUnlock()

	for _, f := range beforeRequestFunc {
		err := f.Handle(ctx)
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	return  nil
}
