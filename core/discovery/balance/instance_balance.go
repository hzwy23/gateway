package balance

import (
	"errors"
	"github.com/wisrc/gateway/core/discovery/register"
	"math/rand"
	"strings"
	"time"
)

type InstanceBalance struct {
	register *register.ApplicationRegisterCenter
}

func NewInstanceBalance(register *register.ApplicationRegisterCenter) *InstanceBalance {
	return &InstanceBalance{
		register: register,
	}
}

// GetService 根据微服务名，获取微服务 Scheme，IP，Port 信息
func (r *InstanceBalance)GetService(serviceId string) (*register.AppInstance, error) {
	id := strings.ToUpper(serviceId)
	if app, ok := r.register.Services[id]; ok {
		return r.getInstance(app)
	}
	return nil, errors.New(serviceId + ",应用服务不存在")
}

// getInstance 获取有效的实例
func (r *InstanceBalance)getInstance(app *register.AppService) (*register.AppInstance, error) {
	r.register.Lock.RLock()
	// 生成随机数，从现有节点随机选取一个节点
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(app.Instances))
	inst := app.Instances[idx]
	if inst.Status == register.DOWN {
		// 服务已过期
		r.register.Lock.RUnlock()

		// 剔除宕机的服务
		r.register.Lock.Lock()
		app.Instances = append(app.Instances[:idx], app.Instances[idx+1:]...)
		r.register.Lock.Unlock()

		// 如果可用节点为空，则服务所有节点已宕机
		if len(app.Instances) == 0 {
			return nil, errors.New("无有效的实例")
		}
		// 在剩余正常实例中查找可用的节点
		return r.getInstance(app)
	}
	r.register.Lock.RUnlock()
	return inst, nil
}
