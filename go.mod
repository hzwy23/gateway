module "github.com/wisrc/gateway"

replace {
    "golang.org/x/sys" => "github.com/golang/x/sys" latest
    "golang.org/x/tools" => "github.com/golang/x/tools" latest
}
