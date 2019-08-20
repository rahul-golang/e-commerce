package endpoint

import (
	"context"
	models "gokit/ecommerse/users/pkg/models"
	service "gokit/ecommerse/users/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateUserRequest collects the request parameters for the CreateUser method.
type CreateUserRequest struct {
	CreateReq models.CreateUserReq `json:"create_req"`
}

// CreateUserResponse collects the response parameters for the CreateUser method.
type CreateUserResponse struct {
	CreateResp *models.CreateUserResp `json:"create_resp"`
	Err        error                  `json:"err"`
}

// MakeCreateUserEndpoint returns an endpoint that invokes CreateUser on the service.
func MakeCreateUserEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		createResp, err := s.CreateUser(ctx, req.CreateReq)
		return &CreateUserResponse{
			CreateResp: createResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateUserResponse) Failed() error {
	return r.Err
}

// GetAllUserRequest collects the request parameters for the GetAllUser method.
type GetAllUserRequest struct{}

// GetAllUserResponse collects the response parameters for the GetAllUser method.
type GetAllUserResponse struct {
	AllRecordResp []*models.User `json:"all_record_resp"`
	Err           error          `json:"err"`
}

// MakeGetAllUserEndpoint returns an endpoint that invokes GetAllUser on the service.
func MakeGetAllUserEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		allRecordResp, err := s.GetAllUser(ctx)
		return GetAllUserResponse{
			AllRecordResp: allRecordResp,
			Err:           err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAllUserResponse) Failed() error {
	return r.Err
}

// UpdateUserRequest collects the request parameters for the UpdateUser method.
type UpdateUserRequest struct {
	UpadteReq models.User `json:"upadte_req"`
}

// UpdateUserResponse collects the response parameters for the UpdateUser method.
type UpdateUserResponse struct {
	UpdateResp *models.User `json:"update_resp"`
	Err        error        `json:"err"`
}

// MakeUpdateUserEndpoint returns an endpoint that invokes UpdateUser on the service.
func MakeUpdateUserEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		updateResp, err := s.UpdateUser(ctx, req.UpadteReq)
		return UpdateUserResponse{
			Err:        err,
			UpdateResp: updateResp,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserResponse) Failed() error {
	return r.Err
}

// DeleteUserRequest collects the request parameters for the DeleteUser method.
type DeleteUserRequest struct {
	Id string `json:"id"`
}

// DeleteUserResponse collects the response parameters for the DeleteUser method.
type DeleteUserResponse struct {
	DeleteResp *models.DeleteUserResp `json:"delete_resp"`
	Err        error                  `json:"err"`
}

// MakeDeleteUserEndpoint returns an endpoint that invokes DeleteUser on the service.
func MakeDeleteUserEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		deleteResp, err := s.DeleteUser(ctx, req.Id)
		return DeleteUserResponse{
			DeleteResp: deleteResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserResponse) Failed() error {
	return r.Err
}

// GetUserRequest collects the request parameters for the GetUser method.
type GetUserRequest struct {
	Id string `json:"id"`
}

// GetUserResponse collects the response parameters for the GetUser method.
type GetUserResponse struct {
	CreateResp *models.User `json:"create_resp"`
	Err        error        `json:"err"`
}

// MakeGetUserEndpoint returns an endpoint that invokes GetUser on the service.
func MakeGetUserEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		createResp, err := s.GetUser(ctx, req.Id)
		return GetUserResponse{
			CreateResp: createResp,
			Err:        err,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateUser implements Service. Primarily useful in a client.
func (e Endpoints) CreateUser(ctx context.Context, createReq models.CreateUserReq) (createResp *models.CreateUserResp, err error) {
	request := CreateUserRequest{CreateReq: createReq}
	response, err := e.CreateUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateUserResponse).CreateResp, response.(CreateUserResponse).Err
}

// GetAllUser implements Service. Primarily useful in a client.
func (e Endpoints) GetAllUser(ctx context.Context) (allRecordResp []*models.User, err error) {
	request := GetAllUserRequest{}
	response, err := e.GetAllUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllUserResponse).AllRecordResp, response.(GetAllUserResponse).Err
}

// UpdateUser implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUser(ctx context.Context, upadteReq models.User) (updateResp *models.User, err error) {
	request := UpdateUserRequest{UpadteReq: upadteReq}
	response, err := e.UpdateUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateUserResponse).UpdateResp, response.(UpdateUserResponse).Err
}

// DeleteUser implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error) {
	request := DeleteUserRequest{Id: id}
	response, err := e.DeleteUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteUserResponse).DeleteResp, response.(DeleteUserResponse).Err
}

// GetUser implements Service. Primarily useful in a client.
func (e Endpoints) GetUser(ctx context.Context, id string) (createResp *models.User, err error) {
	request := GetUserRequest{Id: id}
	response, err := e.GetUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserResponse).CreateResp, response.(GetUserResponse).Err
}
