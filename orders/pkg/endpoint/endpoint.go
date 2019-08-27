package endpoint

import (
	"context"
	models "gokit/ecommerse/orders/pkg/models"
	service "gokit/ecommerse/orders/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateOrderRequest collects the request parameters for the CreateOrder method.
type CreateOrderRequest struct {
	OrderReq models.CrateOrderReq `json:"order_req"`
}

// CreateOrderResponse collects the response parameters for the CreateOrder method.
type CreateOrderResponse struct {
	OrderResp *models.CrateOrderResp `json:"order_resp"`
	Err       error                  `json:"err"`
}

// MakeCreateOrderEndpoint returns an endpoint that invokes CreateOrder on the service.
func MakeCreateOrderEndpoint(s service.OrdersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateOrderRequest)
		orderResp, err := s.CreateOrder(ctx, req.OrderReq)
		return CreateOrderResponse{
			Err:       err,
			OrderResp: orderResp,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateOrderResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateOrder implements Service. Primarily useful in a client.
func (e Endpoints) CreateOrder(ctx context.Context, orderReq models.CrateOrderReq) (orderResp *models.CrateOrderResp, err error) {
	request := CreateOrderRequest{OrderReq: orderReq}
	response, err := e.CreateOrderEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateOrderResponse).OrderResp, response.(CreateOrderResponse).Err
}
