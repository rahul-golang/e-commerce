package grpc

import (
	"context"
	"errors"
	endpoint "gokit/ecommerse/users/pkg/endpoint"
	pb "gokit/ecommerse/users/pkg/grpc/pb"

	"github.com/mitchellh/mapstructure"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeCreateUserHandler creates the handler logic
func makeCreateUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateUserEndpoint, decodeCreateUserRequest, encodeCreateUserResponse, options...)
}

// decodeCreateUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeCreateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeCreateUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) CreateUser(ctx context1.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	_, rep, err := g.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateUserReply), nil
}

// makeGetAllUserHandler creates the handler logic
func makeGetAllUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetAllUserEndpoint, decodeGetAllUserRequest, encodeGetAllUserResponse, options...)
}

// decodeGetAllUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeGetAllUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return endpoint.GetAllUserRequest{}, nil
}

// encodeGetAllUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetAllUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetAllUserResponse)
	// user := []pb.User{
	// 	ID:resp.AllRecordResp
	// }
	var user []*pb.User
	mapstructure.Decode(resp.AllRecordResp, &user)
	return &pb.GetAllUserReply{
		Message: "All User Information",
		User:    user,
		// User: []*pb.User{
		// 	ID: resp.AllRecordResp,
		// },
	}, nil
}
func (g *grpcServer) GetAllUser(ctx context1.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserReply, error) {
	_, rep, err := g.getAllUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAllUserReply), nil
}

// makeUpdateUserHandler creates the handler logic
func makeUpdateUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateUserEndpoint, decodeUpdateUserRequest, encodeUpdateUserResponse, options...)
}

// decodeUpdateUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeUpdateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeUpdateUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) UpdateUser(ctx context1.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	_, rep, err := g.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserReply), nil
}

// makeDeleteUserHandler creates the handler logic
func makeDeleteUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteUserEndpoint, decodeDeleteUserRequest, encodeDeleteUserResponse, options...)
}

// decodeDeleteUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeDeleteUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeDeleteUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) DeleteUser(ctx context1.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	_, rep, err := g.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteUserReply), nil
}

// makeGetUserHandler creates the handler logic
func makeGetUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserEndpoint, decodeGetUserRequest, encodeGetUserResponse, options...)
}

// decodeGetUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.

func decodeGetUserRequest(_ context.Context, r interface{}) (interface{}, error) {

	req := r.(*pb.GetUserRequest)
	return endpoint.GetUserRequest{Id: req.ID}, nil

}

// encodeGetUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeGetUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetUserResponse)
	return &pb.GetUserReply{
		Message: "Created",
		User: &pb.User{
			ID:        string(resp.CreateResp.ID),
			FirstName: resp.CreateResp.FirstName,
		},
	}, nil

}
func (g *grpcServer) GetUser(ctx context1.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	_, rep, err := g.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserReply), nil
}
