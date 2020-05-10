package main

import (
	"github.com/wisrc/gateway/core"
	"github.com/wisrc/gateway/core/discovery/eureka"
	_ "net/http/pprof"
)

func main() {
	eureka.EnableEurekaClient()
	server := core.NewAPIGatewayServer()
	server.Start()
}
