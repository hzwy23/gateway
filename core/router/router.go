package router

import (
	"errors"
	"github.com/wisrc/gateway/core/discovery"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/wisrc/gateway/config"
)

var router = make(map[string]*Router)

type Router struct {
	// 路由名称
	SubPath string
	// 是否叶子
	isLeaf bool
	// 路由详细信息
	Details *routerDetails
	// 节点所有
	node map[string]*Router
}

type routerDetails struct {
	// 路由地址
	Path string
	// 微服务,驼峰命令属性，需要使用 Tag 标签指定名称
	ServiceId string
	// 目标路由
	Url *url.URL
	// 是否过滤前缀
	StripPrefix bool
	// 超时时间
	Timeout time.Duration
}


// match 路由匹配
func Match(path string) (*url.URL, *Router , error) {
	route := match(path)
	if route == nil {
		return nil, route, errors.New(path + ", 未匹配上路由")
	}

	if route.Details.Url != nil {
		if route.Details.StripPrefix {
			remoteUrl := &url.URL{
				Host:   route.Details.Url.Host,
				Scheme: route.Details.Url.Scheme,
				Path:   strings.TrimPrefix(path, strings.TrimRight(route.Details.Path, "**")),
			}
			return remoteUrl,route, nil
		} else {
			remoteUrl := &url.URL{
				Host:   route.Details.Url.Host,
				Scheme: route.Details.Url.Scheme,
				Path:   path,
			}
			return remoteUrl,route, nil
		}
	} else {
		inst, err := discovery.GetServiceInstance(route.Details.ServiceId)
		if err != nil {
			return nil, route, err
		}

		scheme := "http"
		if inst.Secure {
			scheme = "https"
		}
		if route.Details.StripPrefix {
			remoteUrl := &url.URL{
				Host:   inst.IpAddr + ":" + strconv.Itoa(inst.Port),
				Scheme: scheme,
				Path:   strings.TrimPrefix(path, strings.TrimRight(route.Details.Path, "**")),
			}
			return remoteUrl, route, nil
		} else {
			remoteUrl := &url.URL{
				Host:   inst.IpAddr + ":" + strconv.Itoa(inst.Port),
				Scheme: scheme,
				Path:   path,
			}
			return remoteUrl, route, nil
		}
	}
}


func match(path string) *Router {
	paths := strings.Split(strings.TrimLeft(path, "/"), "/")
	var head *Router
	for idx, key := range paths {
		if idx == 0 {
			if m, ok := router[key]; ok {
				head = m
			} else {
				return nil
			}
		} else {
			if u, yes := head.node[key]; yes {
				head = u
			} else if v, all := head.node["**"]; all {
				head = v
			} else if idx == len(paths)-1 {
				return head
			} else {
				return nil
			}
		}
		if head.isLeaf {
			return head
		}
	}
	return nil
}

func init() {

	routerConfig := config.GetGatewayRouter()

	for _, v := range routerConfig.Routers {

		subList := strings.Split(strings.TrimLeft(v.Path, "/"), "/")
		if len(subList) == 0 {
			continue
		}

		var head *Router
		for idx, key := range subList {
			if idx == 0 {
				if u, ok := router[key]; ok {
					head = u
				} else {
					vu :=strings.Split(v.Url,"://")
					var ur *url.URL = nil
					if len(vu) == 2  {
						ur = &url.URL{
							Scheme: vu[0],
							Host: vu[1],
						}
					}
					n := &Router{
						SubPath: key,
						Details:  &routerDetails {
							Path: v.Path,
							ServiceId:v.ServiceId,
							Url:ur,
							StripPrefix: v.StripPrefix,
							Timeout: v.Timeout,
						},
						isLeaf:  true,
						node:    make(map[string]*Router),
					}
					router[key] = n
					head = n
				}
			} else {
				if m, ok := head.node[key]; ok {
					head = m
				} else {
					vu :=strings.Split(v.Url,"://")
					var ur *url.URL = nil
					if len(vu) == 2  {
						ur = &url.URL{
							Scheme: vu[0],
							Host: vu[1],
						}
					}
					n := &Router{
						SubPath: key,
						Details:  &routerDetails {
							Path: v.Path,
							ServiceId:v.ServiceId,
							Url: ur,
							StripPrefix: v.StripPrefix,
							Timeout: v.Timeout,
						},
						isLeaf:  true,
						node:    make(map[string]*Router),
					}
					head.isLeaf = false
					if head.node == nil {
						head.node = make(map[string]*Router)
					}
					head.node[key] = n
					head = n
				}
			}

		}
	}
}
