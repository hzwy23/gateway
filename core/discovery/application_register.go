package discovery

import (
	"sync"
)

type AppService struct {
	ServiceId  string
	Instances  []*AppInstance
	UpdateTime int64
}

type AppInstance struct {
	// 实例ID
	InstanceId string
	// 服务 IP 地址
	IpAddr string
	// 服务端口号
	Port int
	// 服务状态
	Status string
	// 是否启用 SSL
	Secure bool
}

var serviceRegister = make(map[string]*AppService)
var lock = &sync.RWMutex{}

func UpdateApplication(app *AppService) {
	lock.Lock()
	defer lock.Unlock()
	serviceRegister[app.ServiceId] = app
}
