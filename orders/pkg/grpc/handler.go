package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "gokit/ecommerse/orders/pkg/endpoint"
	pb "gokit/ecommerse/orders/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeCreateOrderHandler creates the handler logic
func makeCreateOrderHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateOrderEndpoint, decodeCreateOrderRequest, encodeCreateOrderResponse, options...)
}

// decodeCreateOrderResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeCreateOrderRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Orders' Decoder is not impelemented")
}

// encodeCreateOrderResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateOrderResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Orders' Encoder is not impelemented")
}
func (g *grpcServer) CreateOrder(ctx context1.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	_, rep, err := g.createOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateOrderReply), nil
}
