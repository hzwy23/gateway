package discovery

import (
	"github.com/wisrc/gateway/core/discovery/balance"
	"github.com/wisrc/gateway/core/discovery/register"
	"github.com/wisrc/gateway/core/discovery/register/eureka"
	"sync"
)

type Discovery struct {
	instBalance *balance.InstanceBalance
	lock *sync.RWMutex
}

var clientDiscovery = &Discovery{
	lock: &sync.RWMutex{},
}

func EnableDiscovery()  {
	er := eureka.NewEurekaRegister()
	clientDiscovery.lock.Lock()
	defer clientDiscovery.lock.Unlock()
	clientDiscovery.instBalance = balance.NewInstanceBalance(er)
}


func GetServiceInstance(serviceId string)(*register.AppInstance, error)  {
	return clientDiscovery.instBalance.GetService(serviceId)
}