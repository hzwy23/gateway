package config

import (
	"time"
)

type Routers struct {
	Routers          map[string]RouterDetails
	IgnoredPatterns  []string `yaml:"ignoredPatterns"`
	SensitiveHeaders []string `yaml:"sensitiveHeaders"`
}

type RouterDetails struct {
	// 路由地址
	Path string
	// 微服务,驼峰命令属性，需要使用 Tag 标签指定名称
	ServiceId string `yaml:"serviceId"`
	// 目标路由
	Url string
	// 是否过滤前缀
	StripPrefix bool `yaml:"stripPrefix"`
	// 超时时间
	Timeout time.Duration
}
