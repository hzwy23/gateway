package eureka

import (
	"encoding/json"
	"github.com/wisrc/gateway/core/discovery/register"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/wisrc/gateway/config"
	"github.com/wisrc/gateway/logger"
)

const eurekaApp = "/eureka/apps"

type eurekaRegister struct {
	register *register.ApplicationRegisterCenter
}

func NewEurekaRegister() *register.ApplicationRegisterCenter {
	reg := register.NewApplicationRegisterCenter()
	r := &eurekaRegister{
		register: reg,
	}
	r.enableEurekaClient()
	return reg
}

type eurekaApps struct {
	Applications struct {
		VersionsDelta string `json:"versions__delta"`
		AppsHashcode  string `json:"apps__hashcode"`
		Application   []struct {
			Name     string `json:"name"`
			Instance []struct {
				InstanceId string `json:"instanceId"`
				HostName   string `json:"hostName"`
				App        string `json:"app"`
				IpAddr     string `json:"ipAddr"`
				Status     string `json:"status"`
				Port       struct {
					Port    int    `json:"$"`
					Enabled string `json:"@enabled"`
				} `json:"port"`
				SecurePort struct {
					Port    int    `json:"$"`
					Enabled string `json:"@enabled"`
				} `json:"securePort"`
			} `json:"instance"`
		} `json:"application"`
	} `json:"applications"`
}


func (r *eurekaRegister)enableEurekaClient() {
	conf := config.GetRegisterCenter()
	eurekaRefresh := time.NewTicker(time.Second * conf.RefreshFrequency)
	go func(tick *time.Ticker) {
		defer tick.Stop()
		for {
			logger.Info("同步 Eureka 注册中心")
			for _, url := range conf.EurekaConfig.ServiceUrls {
				remoteUrl := url + eurekaApp
				body, err := r.httpRequest(http.MethodGet, remoteUrl)
				if err != nil {
					logger.Error(err.Error())
					continue
				}
				rst := eurekaApps{}
				err = json.Unmarshal(body, &rst)
				if err != nil {
					logger.Error(err.Error())
					continue
				}
				go r.updateRegister(&rst)
			}
			<-tick.C
		}
	}(eurekaRefresh)
}

func (r *eurekaRegister)httpRequest(method, url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (r *eurekaRegister)updateRegister(apps *eurekaApps) {
	for _, app := range apps.Applications.Application {
		instances := make([]*register.AppInstance, 0)
		for _, inst := range app.Instance {
			if inst.Status != register.UP {
				logger.Error("DOWN 掉的实例", inst.App, ",", inst.InstanceId, ",", inst.IpAddr, ",", inst.Status)
				continue
			}
			port := 0
			secure := false
			if inst.Port.Enabled == "true" {
				port = inst.Port.Port
			} else if inst.SecurePort.Enabled == "true" {
				port = inst.SecurePort.Port
				secure = true
			} else {
				logger.Error("服务已经关闭了")
			}

			instance := &register.AppInstance{
				InstanceId: inst.InstanceId,
				IpAddr:     inst.IpAddr,
				Status:     inst.Status,
				Port:       port,
				Secure:     secure,
			}
			instances = append(instances, instance)
		}

		appService := &register.AppService{
			ServiceId:  app.Name,
			Instances:  instances,
			UpdateTime: time.Now().Unix(),
		}
		r.register.UpdateApplication(appService)
	}
}
