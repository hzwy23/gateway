package core

import (
	"github.com/wisrc/gateway/config"
	"github.com/wisrc/gateway/core/context"
	"github.com/wisrc/gateway/core/filter"
	"github.com/wisrc/gateway/core/response"
	"github.com/wisrc/gateway/core/router"
	"github.com/wisrc/gateway/logger"
	"net/http"
	"net/http/httputil"
)

type GatewayProxy struct {

}

func NewGatewayProxy() *GatewayProxy {
	return &GatewayProxy{}
}


// dispatch API 请求分发
func (r *GatewayProxy)dispatch(w http.ResponseWriter, req *http.Request) {
	ctx := context.NewContext(w, req)

	defer func(){
		if err := recover(); err != nil {
			logger.Error("dispatch:", err)
			r.globalRecover(ctx, err)
		}
	}()

	// 收到消息之后，统一进行处理
	err := filter.BeforeRequestFilter(ctx)
	if err != nil {
		response.ErrorHandle(ctx, err)
		return
	}

	// 请求后端服务
	err = r.httpProxy(ctx)
	if err != nil {
		response.ErrorHandle(ctx, err)
		return
	}

	// 执行后置过滤器
	err = filter.AfterResponseFilter(ctx)
	if err != nil {
		response.ErrorHandle(ctx, err)
		return
	}
}


// httpProxy 发起 http 请求
func (r *GatewayProxy)httpProxy(ctx *context.GatewayContext) error {

	// 匹配路由
	path := ctx.Request.URL.Path
	remoteUrl,route, err := router.Match(path)
	if err != nil {
		logger.Error("route is: ",route,", error is:",err)
		return response.NewError(response.ProxyUrlNotFound, err.Error())
	}
	ctx.RemoteURL = remoteUrl

	// 创建代理对象
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.Header.Add("X-Forwarded-Host", req.Host)
			req.Header.Add("X-Origin-Host", remoteUrl.Host)
			req.Host = remoteUrl.Host
			req.URL.Scheme = remoteUrl.Scheme
			req.URL.Host = remoteUrl.Host
			req.URL.Path = remoteUrl.Path

			r.filterSensitiveHeaders(req)

			if ctx.Request.URL.RawQuery == "" || req.URL.RawQuery == "" {
				req.URL.RawQuery = ctx.Request.URL.RawQuery + req.URL.RawQuery
			} else {
				req.URL.RawQuery = ctx.Request.URL.RawQuery + "&" + req.URL.RawQuery
			}
			ctx.Request = req
			logger.Debug("director: ", remoteUrl.String())
		},
		ModifyResponse: func(resp *http.Response) error {
			logger.Debug("modify response:", remoteUrl.String())
			ctx.Response = resp
			err := filter.BeforeResponseFilter(ctx)
			if err != nil {
				logger.Info("BeforeResponse Stop")
			}
			return nil
		},
		ErrorHandler: r.ErrorHandler,
		Transport: defaultGatewayTransport.GetTransport(route),
	}

	proxy.ServeHTTP(ctx.ResponseWriter, ctx.Request)

	logger.Info("请求完成, 请求地址：", path, "，目标地址：", remoteUrl)
	return nil
}

// globalErrorHandle Proxy 代理处理错误
func (r *GatewayProxy)ErrorHandler(w http.ResponseWriter, request *http.Request, err error) {
	logger.Error(err.Error())
	w.Header().Set("Content-Type","application/json;charset=UTF-8")
	w.WriteHeader(http.StatusBadGateway)
	w.Write([]byte(response.NewError(response.ProxyError, err.Error()).Error()))
}

// globalRecover 捕获全局异常
func (r *GatewayProxy)globalRecover(ctx *context.GatewayContext, errMsg interface{}){
	if ctx.ResponseWriter != nil {
		ctx.ResponseWriter.Header().Set("Content-Type","application/json;charset=UTF-8")
		ctx.ResponseWriter.WriteHeader(http.StatusBadGateway)
		ctx.ResponseWriter.Write([]byte(response.NewError(response.ProxyError, errMsg).Error()))
	}
}

// filterSensitiveHeaders 过滤掉请求 Header 中配置的 Key
func (r *GatewayProxy) filterSensitiveHeaders(req *http.Request)  {
	for _, header := range config.GetGatewayRouter().SensitiveHeaders {
		req.Header.Del(header)
	}
}


//// 创建请求 request
//req, err := http.NewRequest(ctx.Request.Method, remoteUrl.String(), nil)
//if err != nil {
//	logger.Error(err)
//	return NewRestResponse(CreateHttpRequestFailed, StatusStr(CreateHttpRequestFailed), err.Error()), err
//}

//// 发起请求
//resp, err := (&http.Client{
//	// 设置请求超时时间
//	Timeout: time.Second * route.Router.Timeout,
//}).Do(req)

//if err != nil {
//	logger.Error(err.Error())
//	return NewRestResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()), err
//}
//defer resp.Body.Close()
//ctx.resp = resp