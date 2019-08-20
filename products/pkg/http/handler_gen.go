// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	endpoint "gokit/ecommerse/products/pkg/endpoint"
	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeCreateProductHandler(m, endpoints, options["CreateProduct"])
	makeGetAllProductHandler(m, endpoints, options["GetAllProduct"])
	makeUpdateProductHandler(m, endpoints, options["UpdateProduct"])
	makeDeleteProductHandler(m, endpoints, options["DeleteProduct"])
	makeGetProductHandler(m, endpoints, options["GetProduct"])
	return m
}
