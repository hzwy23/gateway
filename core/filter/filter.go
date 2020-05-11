package filter

import (
	"github.com/wisrc/gateway/core/context"
)

type Handler struct {
	Name string
	Priority int
	Handle HandleFilter
}

type HandleFilter func(ctx *context.GatewayContext)  error

const (
	BeforeRequest = iota
	BeforeResponse
	AfterResponse
)

func RegisterFilter(code int, handle Handler) {
	switch code {
		case BeforeRequest:
			registerBeforeRequest(handle)
		case BeforeResponse:
			registerBeforeResponse(handle)
		case AfterResponse:
			registerAfterResponse(handle)
		default:
			panic("无效过滤器类型")
	}
}