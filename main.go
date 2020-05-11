package main

import (
	"github.com/wisrc/gateway/core"
	"github.com/wisrc/gateway/core/discovery"
)

func main() {
	discovery.EnableDiscovery()
	server := core.NewAPIGatewayServer()
	server.Start()
}
