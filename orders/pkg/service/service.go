package service

import (
	"context"
	"fmt"
	"gokit/ecommerse/orders/pkg/models"
	"gokit/ecommerse/orders/pkg/repository"
)

// OrdersService describes the service.
type OrdersService interface {
	CreateOrder(ctx context.Context, orderReq models.CrateOrderReq) (orderResp *models.CrateOrderResp, err error)
}

type basicOrdersService struct {
	repoInterface repository.OrderRepositoryInterface
}

func (b *basicOrdersService) CreateOrder(ctx context.Context, orderReq models.CrateOrderReq) (*models.CrateOrderResp, error) {
	orderResp, err := b.repoInterface.CreateOrders(ctx, orderReq.Orders)
	if err != nil {
		fmt.Println("CreateOrder() in service layer ", err)
		return &models.CrateOrderResp{
			Message: "Error in Order placed : " + err.Error(),
			Orders:  nil,
		}, nil
	}

	return &models.CrateOrderResp{
		Message: "Order placed Sucsessfully",
		Orders:  orderResp,
	}, nil
}

// NewBasicOrdersService returns a naive, stateless implementation of OrdersService.
func NewBasicOrdersService(orderRepositoryInterface repository.OrderRepositoryInterface) OrdersService {
	return &basicOrdersService{repoInterface: orderRepositoryInterface}
}

// New returns a OrdersService with all of the expected middleware wired in.
func New(middleware []Middleware, orderRepositoryInterface repository.OrderRepositoryInterface) OrdersService {
	var svc OrdersService = NewBasicOrdersService(orderRepositoryInterface)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
