package config

import (
	"io/ioutil"

	"github.com/wisrc/gateway/logger"
	"gopkg.in/yaml.v2"
)

var appConfig = &AppConfig{}

type AppConfig struct {
	// eureka 配置信息
	RegisterCenter *RegisterCenter `yaml:"registerCenter"`

	// 路由配置信息
	GatewayRouter *Routers `yaml:"gateway"`

	// web 服务配置信息
	Server *ServerConfig `yaml:"server"`
}

func GetServerConfig() *ServerConfig {
	return appConfig.Server
}

func GetGatewayRouter() *Routers {
	return appConfig.GatewayRouter
}

func GetRegisterCenter() *RegisterCenter {
	return appConfig.RegisterCenter
}

func init() {
	logger.Info("加载路由配置信息")
	data, err := ioutil.ReadFile("conf/properties.yml")
	if err != nil {
		logger.Error("从 conf/properties.yml 配置文件中加载配置信息失败")
	}
	yaml.Unmarshal(data, appConfig)
	logger.Info("路由配置信息解析完成")
}
