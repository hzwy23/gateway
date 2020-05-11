package balance

import (
	"fmt"
	"github.com/wisrc/gateway/core/discovery/register"
	"testing"
	"time"
)

func TestUpdateApplication(t *testing.T) {
	instance := []*register.AppInstance{
		&register.AppInstance{
			InstanceId: "1",
			IpAddr:     "https://gitbook.cn",
			Port:       443,
			Status:     register.UP,
			Secure:     true,
		},
		&register.AppInstance{
			InstanceId: "2",
			IpAddr:     "https://gitbook.cn",
			Port:       443,
			Status:     register.DOWN,
			Secure:     true,
		},
		&register.AppInstance{
			InstanceId: "3",
			IpAddr:     "https://gitbook.cn",
			Port:       443,
			Status:     register.UP,
			Secure:     true,
		},
		&register.AppInstance{
			InstanceId: "4",
			IpAddr:     "https://gitbook.cn",
			Port:       443,
			Status:     register.UP,
			Secure:     true,
		},
	}

	reg := register.NewApplicationRegisterCenter()

	reg.UpdateApplication(&register.AppService{
		ServiceId: "demo",
		Instances: instance,
		UpdateTime: time.Now().Unix(),
	})

	bal := NewInstanceBalance(reg)

	inst, err := bal.GetService("demo")
	fmt.Println(inst, err, bal)

	inst, err = bal.GetService("demo")
	fmt.Println(inst, err, bal)

	inst, err = bal.GetService("demo")
	fmt.Println(inst, err, bal)

	inst, err = bal.GetService("demo")
	fmt.Println(inst, err, bal)
}