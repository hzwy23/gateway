package main

import (
	"github.com/wisrc/gateway/core"
	"github.com/wisrc/gateway/core/discovery/eureka"
)

func main() {
	eureka.EnableEurekaClient()
	server := core.NewAPIGatewayServer()
	server.Start()
}
