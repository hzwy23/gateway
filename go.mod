module "github.com/wisrc/gateway"

replace (
	golang.org/x/sys latest => github.com/golang/sys latest
	golang.org/x/tools latest => github.com/golang/tools latest
	golang.org/x/sys/unix latest => github.com/golang/sys/unix latest
)
