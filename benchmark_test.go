package main

import (
	"github.com/wisrc/gateway/core"
	"github.com/wisrc/gateway/core/discovery/eureka"
	"testing"
)

func BenchmarkGateway(b *testing.B)  {
	eureka.EnableEurekaClient()
	server := core.NewAPIGatewayServer()
	server.Start()
}