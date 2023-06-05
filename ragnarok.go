package main

import (
	"github.com/alphadecodex/nomise-ragnarok/resources"
	"github.com/alphadecodex/nomise-ragnarok/service"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	var OrderResource resources.OrderResource
	err := container.Invoke(func(orderService *service.OrderService) {
		OrderResource = *resources.NewOrderResource(orderService)
	})
	if err != nil {
		// Handle the error
	}
}
