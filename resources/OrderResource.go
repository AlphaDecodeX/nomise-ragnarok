package resources

import (
	"github.com/alphadecodex/nomise-ragnarok/entity"
	"github.com/alphadecodex/nomise-ragnarok/service"
	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	order entity.Order
}

type OrderResource struct {
	OrderService *service.OrderService
}

func NewOrderResource(orderService *service.OrderService) *OrderResource {
	return &OrderResource{
		OrderService: orderService,
	}
}

func (OrderResource *OrderResource) CreateOrder(ctx *gin.Context) error {
	var request OrderRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		return err
	}
	err = OrderResource.OrderService.CreateOrUpdateOrder(&request.order)
	if err != nil {
		// Handle the error
		return err
	}
	return nil

}
