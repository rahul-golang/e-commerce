// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "gokit/ecommerse/products/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateProductEndpoint endpoint.Endpoint
	GetAllProductEndpoint endpoint.Endpoint
	UpdateProductEndpoint endpoint.Endpoint
	DeleteProductEndpoint endpoint.Endpoint
	GetProductEndpoint    endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.ProductsService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		CreateProductEndpoint: MakeCreateProductEndpoint(s),
		DeleteProductEndpoint: MakeDeleteProductEndpoint(s),
		GetAllProductEndpoint: MakeGetAllProductEndpoint(s),
		GetProductEndpoint:    MakeGetProductEndpoint(s),
		UpdateProductEndpoint: MakeUpdateProductEndpoint(s),
	}
	for _, m := range mdw["CreateProduct"] {
		eps.CreateProductEndpoint = m(eps.CreateProductEndpoint)
	}
	for _, m := range mdw["GetAllProduct"] {
		eps.GetAllProductEndpoint = m(eps.GetAllProductEndpoint)
	}
	for _, m := range mdw["UpdateProduct"] {
		eps.UpdateProductEndpoint = m(eps.UpdateProductEndpoint)
	}
	for _, m := range mdw["DeleteProduct"] {
		eps.DeleteProductEndpoint = m(eps.DeleteProductEndpoint)
	}
	for _, m := range mdw["GetProduct"] {
		eps.GetProductEndpoint = m(eps.GetProductEndpoint)
	}
	return eps
}
