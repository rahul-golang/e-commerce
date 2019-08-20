package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	endpoint "gokit/ecommerse/carts/pkg/endpoint"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
)

// makeAddToCartHandler creates the handler logic
func makeAddToCartHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/carts").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.AddToCartEndpoint, decodeAddToCartRequest, encodeAddToCartResponse, options...)))
}

// decodeAddToCartRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddToCartRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.AddToCartRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	fmt.Println("Add to card request ", req)
	return req, err
}

// encodeAddToCartResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddToCartResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetFromAllHandler creates the handler logic
func makeGetFromAllHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/carts").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetFromAllEndpoint, decodeGetFromAllRequest, encodeGetFromAllResponse, options...)))
}

// decodeGetFromAllRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetFromAllRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetFromAllRequest{}

	return req, nil
}

// encodeGetFromAllResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetFromAllResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateInCartHandler creates the handler logic
func makeUpdateInCartHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/carts").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateInCartEndpoint, decodeUpdateInCartRequest, encodeUpdateInCartResponse, options...)))
}

// decodeUpdateInCartRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateInCartRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateInCartRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateInCartResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateInCartResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteFromCartHandler creates the handler logic
func makeDeleteFromCartHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE").Path("/carts/{id}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"DELETE"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteFromCartEndpoint, decodeDeleteFromCartRequest, encodeDeleteFromCartResponse, options...)))
}

// decodeDeleteFromCartRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteFromCartRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteFromCartRequest{}
	vars := mux.Vars(r)
	req.Id = vars["id"]
	return req, nil
}

// encodeDeleteFromCartResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteFromCartResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetFromCartHandler creates the handler logic
func makeGetFromCartHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/carts/{id}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetFromCartEndpoint, decodeGetFromCartRequest, encodeGetFromCartResponse, options...)))
}

// decodeGetFromCartRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetFromCartRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetFromCartRequest{}

	vars := mux.Vars(r)
	req.Id = vars["id"]
	fmt.Println(req)
	return req, nil
}

// encodeGetFromCartResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetFromCartResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
