package grpc

import (
	"context"
	"errors"
	endpoint "gokit/ecommerse/products/pkg/endpoint"
	pb "gokit/ecommerse/products/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/mitchellh/mapstructure"
	context1 "golang.org/x/net/context"
)

// makeCreateProductHandler creates the handler logic
func makeCreateProductHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateProductEndpoint, decodeCreateProductRequest, encodeCreateProductResponse, options...)
}

// decodeCreateProductResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeCreateProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Products' Decoder is not impelemented")
}

// encodeCreateProductResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateProductResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Products' Encoder is not impelemented")
}
func (g *grpcServer) CreateProduct(ctx context1.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	_, rep, err := g.createProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateProductReply), nil
}

// makeGetAllProductHandler creates the handler logic
func makeGetAllProductHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetAllProductEndpoint, decodeGetAllProductRequest, encodeGetAllProductResponse, options...)
}

// decodeGetAllProductResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeGetAllProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Products' Decoder is not impelemented")
}

// encodeGetAllProductResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetAllProductResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Products' Encoder is not impelemented")
}
func (g *grpcServer) GetAllProduct(ctx context1.Context, req *pb.GetAllProductRequest) (*pb.GetAllProductReply, error) {
	_, rep, err := g.getAllProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllProductReply), nil
}

// makeUpdateProductHandler creates the handler logic
func makeUpdateProductHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateProductEndpoint, decodeUpdateProductRequest, encodeUpdateProductResponse, options...)
}

// decodeUpdateProductResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeUpdateProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Products' Decoder is not impelemented")
}

// encodeUpdateProductResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateProductResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Products' Encoder is not impelemented")
}
func (g *grpcServer) UpdateProduct(ctx context1.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductReply, error) {
	_, rep, err := g.updateProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateProductReply), nil
}

// makeDeleteProductHandler creates the handler logic
func makeDeleteProductHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteProductEndpoint, decodeDeleteProductRequest, encodeDeleteProductResponse, options...)
}

// decodeDeleteProductResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeDeleteProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Products' Decoder is not impelemented")
}

// encodeDeleteProductResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteProductResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Products' Encoder is not impelemented")
}
func (g *grpcServer) DeleteProduct(ctx context1.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductReply, error) {
	_, rep, err := g.deleteProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteProductReply), nil
}

// makeGetProductHandler creates the handler logic
func makeGetProductHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetProductEndpoint, decodeGetProductRequest, encodeGetProductResponse, options...)
}

// decodeGetProductResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.

func decodeGetProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(pb.UpdateProductStockRequest)
	endpointReq := endpoint.UpdateProductStockRequest{}
	mapstructure.Decode(req, &endpointReq)

	return &endpointReq, nil
}

// encodeGetProductResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.

func encodeGetProductResponse(_ context.Context, r interface{}) (interface{}, error) {

	resp := r.(endpoint.UpdateProductStockResponse)
	endpointResp := pb.UpdateProductStockReply{
		Message: resp.UpdateStockResp.Message,
	}
	return &endpointResp, nil
}
func (g *grpcServer) GetProduct(ctx context1.Context, req *pb.GetProductRequest) (*pb.GetProductReply, error) {
	_, rep, err := g.getProduct.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetProductReply), nil
}

// makeUpdateProductStockHandler creates the handler logic
func makeUpdateProductStockHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateProductStockEndpoint, decodeUpdateProductStockRequest, encodeUpdateProductStockResponse, options...)
}

// decodeUpdateProductStockResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeUpdateProductStockRequest(_ context.Context, r interface{}) (interface{}, error) {

	req := r.(*pb.UpdateProductStockRequest)
	endpointReq := endpoint.UpdateProductStockRequest{}

	mapstructure.Decode(req, &endpointReq)

	return &endpointReq, nil
}

// encodeUpdateProductStockResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeUpdateProductStockResponse(_ context.Context, r interface{}) (interface{}, error) {

	resp := r.(endpoint.UpdateProductStockResponse)
	grpcHandleResp := pb.UpdateProductStockReply{}
	grpcHandleResp.Message = resp.UpdateStockResp.Message
	return &grpcHandleResp, nil
}
func (g *grpcServer) UpdateProductStock(ctx context1.Context, req *pb.UpdateProductStockRequest) (*pb.UpdateProductStockReply, error) {
	_, rep, err := g.updateProductStock.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateProductStockReply), nil
}
