package endpoint

import (
	"context"
	models "gokit/ecommerse/products/pkg/models"
	service "gokit/ecommerse/products/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateProductRequest collects the request parameters for the CreateProduct method.
type CreateProductRequest struct {
	CreateReq models.CreateProductReq `json:"create_req"`
}

// CreateProductResponse collects the response parameters for the CreateProduct method.
type CreateProductResponse struct {
	CreateResp *models.CreateProductResp `json:"create_resp"`
	Err        error                     `json:"err"`
}

// MakeCreateProductEndpoint returns an endpoint that invokes CreateProduct on the service.
func MakeCreateProductEndpoint(s service.ProductsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		createResp, err := s.CreateProduct(ctx, req.CreateReq)
		return CreateProductResponse{
			CreateResp: createResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateProductResponse) Failed() error {
	return r.Err
}

// GetAllProductRequest collects the request parameters for the GetAllProduct method.
type GetAllProductRequest struct{}

// GetAllProductResponse collects the response parameters for the GetAllProduct method.
type GetAllProductResponse struct {
	AllRecordResp []*models.Product `json:"all_record_resp"`
	Err           error             `json:"err"`
}

// MakeGetAllProductEndpoint returns an endpoint that invokes GetAllProduct on the service.
func MakeGetAllProductEndpoint(s service.ProductsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		allRecordResp, err := s.GetAllProduct(ctx)
		return GetAllProductResponse{
			AllRecordResp: allRecordResp,
			Err:           err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAllProductResponse) Failed() error {
	return r.Err
}

// UpdateProductRequest collects the request parameters for the UpdateProduct method.
type UpdateProductRequest struct {
	UpadteReq models.Product `json:"upadte_req"`
}

// UpdateProductResponse collects the response parameters for the UpdateProduct method.
type UpdateProductResponse struct {
	UpdateResp *models.Product `json:"update_resp"`
	Err        error           `json:"err"`
}

// MakeUpdateProductEndpoint returns an endpoint that invokes UpdateProduct on the service.
func MakeUpdateProductEndpoint(s service.ProductsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateProductRequest)
		updateResp, err := s.UpdateProduct(ctx, req.UpadteReq)
		return UpdateProductResponse{
			Err:        err,
			UpdateResp: updateResp,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateProductResponse) Failed() error {
	return r.Err
}

// DeleteProductRequest collects the request parameters for the DeleteProduct method.
type DeleteProductRequest struct {
	Id string `json:"id"`
}

// DeleteProductResponse collects the response parameters for the DeleteProduct method.
type DeleteProductResponse struct {
	DeleteResp *models.DeleteProductResp `json:"delete_resp"`
	Err        error                     `json:"err"`
}

// MakeDeleteProductEndpoint returns an endpoint that invokes DeleteProduct on the service.
func MakeDeleteProductEndpoint(s service.ProductsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProductRequest)
		deleteResp, err := s.DeleteProduct(ctx, req.Id)
		return DeleteProductResponse{
			DeleteResp: deleteResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteProductResponse) Failed() error {
	return r.Err
}

// GetProductRequest collects the request parameters for the GetProduct method.
type GetProductRequest struct {
	Id string `json:"id"`
}

// GetProductResponse collects the response parameters for the GetProduct method.
type GetProductResponse struct {
	CreateResp *models.Product `json:"create_resp"`
	Err        error           `json:"err"`
}

// MakeGetProductEndpoint returns an endpoint that invokes GetProduct on the service.
func MakeGetProductEndpoint(s service.ProductsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)
		createResp, err := s.GetProduct(ctx, req.Id)
		return GetProductResponse{
			CreateResp: createResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetProductResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateProduct implements Service. Primarily useful in a client.
func (e Endpoints) CreateProduct(ctx context.Context, createReq models.CreateProductReq) (createResp *models.CreateProductResp, err error) {
	request := CreateProductRequest{CreateReq: createReq}
	response, err := e.CreateProductEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateProductResponse).CreateResp, response.(CreateProductResponse).Err
}

// GetAllProduct implements Service. Primarily useful in a client.
func (e Endpoints) GetAllProduct(ctx context.Context) (allRecordResp []*models.Product, err error) {
	request := GetAllProductRequest{}
	response, err := e.GetAllProductEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllProductResponse).AllRecordResp, response.(GetAllProductResponse).Err
}

// UpdateProduct implements Service. Primarily useful in a client.
func (e Endpoints) UpdateProduct(ctx context.Context, upadteReq models.Product) (updateResp *models.Product, err error) {
	request := UpdateProductRequest{UpadteReq: upadteReq}
	response, err := e.UpdateProductEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateProductResponse).UpdateResp, response.(UpdateProductResponse).Err
}

// DeleteProduct implements Service. Primarily useful in a client.
func (e Endpoints) DeleteProduct(ctx context.Context, id string) (deleteResp *models.DeleteProductResp, err error) {
	request := DeleteProductRequest{Id: id}
	response, err := e.DeleteProductEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteProductResponse).DeleteResp, response.(DeleteProductResponse).Err
}

// GetProduct implements Service. Primarily useful in a client.
func (e Endpoints) GetProduct(ctx context.Context, id string) (createResp *models.Product, err error) {
	request := GetProductRequest{Id: id}
	response, err := e.GetProductEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetProductResponse).CreateResp, response.(GetProductResponse).Err
}
