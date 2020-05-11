package main

import (
	"github.com/wisrc/gateway/core"
	"github.com/wisrc/gateway/core/discovery/register/eureka"
	_ "net/http/pprof"
	"testing"
)

func BenchmarkGateway(b *testing.B)  {
	eureka.EnableEurekaClient()
	server := core.NewAPIGatewayServer()
	server.Start()
}