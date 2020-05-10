package discovery

import (
	"github.com/wisrc/gateway/config"
	"github.com/wisrc/gateway/logger"
	"time"
)

func refresh() {
	registerCenter := config.GetRegisterCenter()
	ticker := time.NewTicker(time.Second * registerCenter.RefreshFrequency)
	go func(ticker *time.Ticker) {
		for {
			<-ticker.C
			logger.Info("服务状态检测程序更新...")
			for key, app := range serviceRegister {
				if app.UpdateTime-time.Now().Unix() > registerCenter.RefreshFrequency.Nanoseconds()*2 {
					logger.Error(key, "，服务服务DOWN")
					lock.Lock()
					delete(serviceRegister, key)
					lock.Unlock()
				}
			}
		}
	}(ticker)
}

func init() {
	refresh()
}
