package discovery

import (
	"errors"
	"math/rand"
	"strings"
)

func GetService(serviceId string) (*AppInstance, error) {
	id := strings.ToUpper(serviceId)
	if app, ok := serviceRegister[id]; ok {
		lock.RLock()
		defer lock.RUnlock()
		return getInstance(app)
	}
	return nil, errors.New(serviceId + ",应用服务不存在")
}

func getInstance(app *AppService) (*AppInstance, error) {
	idx := rand.Intn(len(app.Instances))
	inst := app.Instances[idx]
	return inst, nil
}
