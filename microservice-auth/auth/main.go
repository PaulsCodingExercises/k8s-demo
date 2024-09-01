package main

import (
	"context"
	"fmt"

	"go-micro.dev/v4"
)

type Greeter struct{}

type HelloRequest struct {
	name string
}

type HelloResponse struct {
	reply string
}

func (g *Greeter) Hello(ctx context.Context, req *HelloRequest, rsp *HelloResponse) error {
	fmt.Println("Hello service was called")
	// (*rsp).reply = fmt.Sprintf("Hello %v. Thank you for interfacing with me.\n", req.name)
	// Business logic goes here...
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hello"),
		micro.WrapHandler(AuthMiddleware()),
	)

	service.Init()

	if err := micro.RegisterHandler(service.Server(), new(Greeter)); err != nil {
		fmt.Println(err)
		return
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
