package service

import (
	"github.com/alphadecodex/nomise-ragnarok/entity"
	"github.com/alphadecodex/nomise-ragnarok/repository"
)

type OrderService struct {
	OrderRepository *repository.OrderRepository
}

func (orderService *OrderService) CreateOrUpdateOrder(order *entity.Order) error {
	return orderService.OrderRepository.CreateOrder(order)
}
