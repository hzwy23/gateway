package discovery

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

type InstanceBalance struct {
	RegisterCenter *ApplicationRegisterCenter
}

var instanceBalance = &InstanceBalance{
	RegisterCenter: serviceRegister,
}

func GetServiceInstance(serviceId string)(*AppInstance, error)  {
	return instanceBalance.GetService(serviceId)
}

// GetService 根据微服务名，获取微服务 Scheme，IP，Port 信息
func (r *InstanceBalance)GetService(serviceId string) (*AppInstance, error) {
	id := strings.ToUpper(serviceId)
	if app, ok := r.RegisterCenter.Services[id]; ok {
		return r.getInstance(app)
	}
	return nil, errors.New(serviceId + ",应用服务不存在")
}

// getInstance 获取有效的实例
func (r *InstanceBalance)getInstance(app *AppService) (*AppInstance, error) {
	r.RegisterCenter.lock.RLock()
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(app.Instances))
	inst := app.Instances[idx]
	if inst.Status == DOWN {
		// 服务已过期
		r.RegisterCenter.lock.RUnlock()
		r.RegisterCenter.lock.Lock()
		app.Instances = append(app.Instances[:idx], app.Instances[idx+1:]...)
		r.RegisterCenter.lock.Unlock()
		if len(app.Instances) == 0 {
			return nil, errors.New("无有效的实例")
		}
		return r.getInstance(app)
	}
	r.RegisterCenter.lock.RUnlock()
	return inst, nil
}
