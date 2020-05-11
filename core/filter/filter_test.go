package filter

import (
	"fmt"
	"github.com/wisrc/gateway/core/context"
	"testing"
)

func TestRegisterFilter(t *testing.T) {
	RegisterFilter(BeforeRequest,  Handler{
		Name: "demo2",
		Priority: 2,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeRequestFunc)
	RegisterFilter(BeforeRequest,  Handler{
		Name: "demo1",
		Priority: 1,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeRequestFunc)
	RegisterFilter(BeforeRequest,  Handler{
		Name: "demo5",
		Priority: 5,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeRequestFunc)

	RegisterFilter(BeforeRequest,  Handler{
		Name: "demo22",
		Priority: 2,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeRequestFunc)

	RegisterFilter(BeforeRequest,  Handler{
		Name: "demo3",
		Priority: 3,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeRequestFunc)
}


func TestRegisterFilter2(t *testing.T) {
	RegisterFilter(BeforeResponse,  Handler{
		Name: "demo2",
		Priority: 2,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeResponseFunc)
	RegisterFilter(BeforeResponse,  Handler{
		Name: "demo1",
		Priority: 1,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeResponseFunc)
	RegisterFilter(BeforeResponse,  Handler{
		Name: "demo5",
		Priority: 5,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeResponseFunc)

	RegisterFilter(BeforeResponse,  Handler{
		Name: "demo22",
		Priority: 2,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeResponseFunc)

	RegisterFilter(BeforeResponse,  Handler{
		Name: "demo3",
		Priority: 3,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(beforeResponseFunc)
}

func TestRegisterFilter3(t *testing.T) {
	RegisterFilter(AfterResponse,  Handler{
		Name: "demo2",
		Priority: 2,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(afterResponseFunc)
	RegisterFilter(AfterResponse,  Handler{
		Name: "demo1",
		Priority: 1,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(afterResponseFunc)
	RegisterFilter(AfterResponse,  Handler{
		Name: "demo5",
		Priority: 5,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(afterResponseFunc)

	RegisterFilter(AfterResponse,  Handler{
		Name: "demo22",
		Priority: 2,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(afterResponseFunc)

	RegisterFilter(AfterResponse,  Handler{
		Name: "demo3",
		Priority: 3,
		Handle: func(ctx *context.GatewayContext) error {
			fmt.Println("hello world")
			return nil
		},
	})
	fmt.Println(afterResponseFunc)
}