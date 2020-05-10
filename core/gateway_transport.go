package core

import (
	"github.com/wisrc/gateway/core/router"
	"net"
	"net/http"
	"sync"
	"time"
)

type GatewayTransport struct {
	Transport map[*router.Router]*http.Transport
	lock *sync.RWMutex
}

var defaultGatewayTransport = &GatewayTransport{
	Transport : make(map[*router.Router]*http.Transport),
	lock : &sync.RWMutex{},
}

func (r *GatewayTransport)GetTransport(route *router.Router) *http.Transport{
	r.lock.RLock()
	if v, ok:= r.Transport[route];ok {
		defer r.lock.RUnlock()
		return v
	} else {
		r.lock.RUnlock()
		return r.createTransport(route)
	}
}

func (r *GatewayTransport)createTransport(route *router.Router) *http.Transport{
	r.lock.Lock()
	defer r.lock.Unlock()
	tp:= &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   route.Details.Timeout * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConnsPerHost:  100,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	r.Transport[route] = tp
	return tp
}
