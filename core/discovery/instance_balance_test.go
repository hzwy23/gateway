package discovery

import (
	"fmt"
	"testing"
	"time"
)

func TestUpdateApplication(t *testing.T) {
	instance := []*AppInstance{
		&AppInstance{
			InstanceId: "1",
			IpAddr: "https://gitbook.cn",
			Port: 443,
			Status: UP,
			Secure: true,
		},
		&AppInstance{
			InstanceId: "2",
			IpAddr: "https://gitbook.cn",
			Port: 443,
			Status: DOWN,
			Secure: true,
		},
		&AppInstance{
			InstanceId: "3",
			IpAddr: "https://gitbook.cn",
			Port: 443,
			Status: UP,
			Secure: true,
		},
		&AppInstance{
			InstanceId: "4",
			IpAddr: "https://gitbook.cn",
			Port: 443,
			Status: UP,
			Secure: true,
		},
	}

	UpdateApplication(&AppService{
		ServiceId: "demo",
		Instances: instance,
		UpdateTime: time.Now().Unix(),
	})

	inst, err := GetServiceInstance("demo")
	fmt.Println(inst, err, serviceRegister.Services["DEMO"].Instances)

	inst, err = GetServiceInstance("demo")
	fmt.Println(inst, err, serviceRegister.Services["DEMO"].Instances)

	inst, err = GetServiceInstance("demo")
	fmt.Println(inst, err, serviceRegister.Services["DEMO"].Instances)

	inst, err = GetServiceInstance("demo")
	fmt.Println(inst, err, serviceRegister.Services["DEMO"].Instances)
}