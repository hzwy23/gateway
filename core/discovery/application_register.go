package discovery

import (
	"github.com/wisrc/gateway/config"
	"github.com/wisrc/gateway/logger"
	"strings"
	"sync"
	"time"
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

const (
	UP   = "UP"
	DOWN = "DOWN"
)

type ApplicationRegisterCenter struct {
	Services map[string]*AppService
	lock *sync.RWMutex
}

var serviceRegister = &ApplicationRegisterCenter{
	Services: make(map[string]*AppService),
	lock: &sync.RWMutex{},
}

func UpdateApplication(app *AppService)  {
	serviceRegister.UpdateApplication(app)
}

func (r *ApplicationRegisterCenter)UpdateApplication(app *AppService) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.Services[strings.ToUpper(app.ServiceId)] = app
}

func (r *ApplicationRegisterCenter)refresh() {
	registerCenter := config.GetRegisterCenter()
	ticker := time.NewTicker(time.Second * registerCenter.RefreshFrequency)
	go func(ticker *time.Ticker) {
		for {
			<-ticker.C
			logger.Info("服务状态检测程序更新...")
			for key, app := range r.Services {
				if app.UpdateTime-time.Now().Unix() > registerCenter.RefreshFrequency.Nanoseconds()*2 {
					logger.Error(key, "，服务服务DOWN")
					r.lock.Lock()
					delete(r.Services, key)
					r.lock.Unlock()
				}
			}
		}
	}(ticker)
}

func init() {
	serviceRegister.refresh()
}
