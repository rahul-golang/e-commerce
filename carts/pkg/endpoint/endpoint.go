package endpoint

import (
	"context"
	models "gokit/ecommerse/carts/pkg/models"
	service "gokit/ecommerse/carts/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// AddToCartRequest collects the request parameters for the AddToCart method.
type AddToCartRequest struct {
	AddReq models.AddReq `json:"add_req"`
}

// AddToCartResponse collects the response parameters for the AddToCart method.
type AddToCartResponse struct {
	AddResp *models.AddResp `json:"add_resp"`
	Err     error           `json:"err"`
}

// MakeAddToCartEndpoint returns an endpoint that invokes AddToCart on the service.
func MakeAddToCartEndpoint(s service.CartsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddToCartRequest)
		addResp, err := s.AddToCart(ctx, req.AddReq)
		return AddToCartResponse{
			AddResp: addResp,
			Err:     err,
		}, nil
	}
}

// Failed implements Failer.
func (r AddToCartResponse) Failed() error {
	return r.Err
}

// GetFromAllRequest collects the request parameters for the GetFromAll method.
type GetFromAllRequest struct{}

// GetFromAllResponse collects the response parameters for the GetFromAll method.
type GetFromAllResponse struct {
	AllRecordResp *models.GetAllResp `json:"all_record_resp"`
	Err           error              `json:"err"`
}

// MakeGetFromAllEndpoint returns an endpoint that invokes GetFromAll on the service.
func MakeGetFromAllEndpoint(s service.CartsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		allRecordResp, err := s.GetFromAll(ctx)
		return GetFromAllResponse{
			AllRecordResp: allRecordResp,
			Err:           err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetFromAllResponse) Failed() error {
	return r.Err
}

// UpdateInCartRequest collects the request parameters for the UpdateInCart method.
type UpdateInCartRequest struct {
	UpadteReq models.UpdateReq `json:"upadte_req"`
}

// UpdateInCartResponse collects the response parameters for the UpdateInCart method.
type UpdateInCartResponse struct {
	UpdateResp *models.UpdateCartResp `json:"update_resp"`
	Err        error                  `json:"err"`
}

// MakeUpdateInCartEndpoint returns an endpoint that invokes UpdateInCart on the service.
func MakeUpdateInCartEndpoint(s service.CartsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateInCartRequest)
		updateResp, err := s.UpdateInCart(ctx, req.UpadteReq)
		return UpdateInCartResponse{
			Err:        err,
			UpdateResp: updateResp,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateInCartResponse) Failed() error {
	return r.Err
}

// DeleteFromCartRequest collects the request parameters for the DeleteFromCart method.
type DeleteFromCartRequest struct {
	Id string `json:"id"`
}

// DeleteFromCartResponse collects the response parameters for the DeleteFromCart method.
type DeleteFromCartResponse struct {
	DeleteResp *models.DeleteResp `json:"delete_resp"`
	Err        error              `json:"err"`
}

// MakeDeleteFromCartEndpoint returns an endpoint that invokes DeleteFromCart on the service.
func MakeDeleteFromCartEndpoint(s service.CartsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteFromCartRequest)
		deleteResp, err := s.DeleteFromCart(ctx, req.Id)
		return DeleteFromCartResponse{
			DeleteResp: deleteResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteFromCartResponse) Failed() error {
	return r.Err
}

// GetFromCartRequest collects the request parameters for the GetFromCart method.
type GetFromCartRequest struct {
	Id string `json:"id"`
}

// GetFromCartResponse collects the response parameters for the GetFromCart method.
type GetFromCartResponse struct {
	CreateResp *models.GetResp `json:"create_resp"`
	Err        error           `json:"err"`
}

// MakeGetFromCartEndpoint returns an endpoint that invokes GetFromCart on the service.
func MakeGetFromCartEndpoint(s service.CartsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFromCartRequest)
		createResp, err := s.GetFromCart(ctx, req.Id)
		return GetFromCartResponse{
			CreateResp: createResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetFromCartResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// AddToCart implements Service. Primarily useful in a client.
func (e Endpoints) AddToCart(ctx context.Context, addReq models.AddReq) (addResp *models.AddResp, err error) {
	request := AddToCartRequest{AddReq: addReq}
	response, err := e.AddToCartEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddToCartResponse).AddResp, response.(AddToCartResponse).Err
}

// GetFromAll implements Service. Primarily useful in a client.
func (e Endpoints) GetFromAll(ctx context.Context) (allRecordResp *models.GetAllResp, err error) {
	request := GetFromAllRequest{}
	response, err := e.GetFromAllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetFromAllResponse).AllRecordResp, response.(GetFromAllResponse).Err
}

// UpdateInCart implements Service. Primarily useful in a client.
func (e Endpoints) UpdateInCart(ctx context.Context, upadteReq models.UpdateReq) (updateResp *models.UpdateCartResp, err error) {
	request := UpdateInCartRequest{UpadteReq: upadteReq}
	response, err := e.UpdateInCartEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateInCartResponse).UpdateResp, response.(UpdateInCartResponse).Err
}

// DeleteFromCart implements Service. Primarily useful in a client.
func (e Endpoints) DeleteFromCart(ctx context.Context, id string) (deleteResp *models.DeleteResp, err error) {
	request := DeleteFromCartRequest{Id: id}
	response, err := e.DeleteFromCartEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteFromCartResponse).DeleteResp, response.(DeleteFromCartResponse).Err
}

// GetFromCart implements Service. Primarily useful in a client.
func (e Endpoints) GetFromCart(ctx context.Context, id string) (createResp *models.GetResp, err error) {
	request := GetFromCartRequest{Id: id}
	response, err := e.GetFromCartEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetFromCartResponse).CreateResp, response.(GetFromCartResponse).Err
}
