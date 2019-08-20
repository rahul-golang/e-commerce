package service

import (
	"context"
	"fmt"
	"gokit/ecommerse/users/pkg/models"
	"gokit/ecommerse/users/pkg/repository"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	CreateUser(ctx context.Context, createReq models.CreateUserReq) (createResp *models.CreateUserResp, err error)
	GetAllUser(ctx context.Context) (allRecordResp []*models.User, err error)
	UpdateUser(ctx context.Context, upadteReq models.User) (updateResp *models.User, err error)
	DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error)
	GetUser(ctx context.Context, id string) (createResp *models.User, err error)
}

type basicUsersService struct {
	userRepositoryInterface repository.UserRepositoryInterface
}

func (b *basicUsersService) CreateUser(ctx context.Context, createReq models.CreateUserReq) (*models.CreateUserResp, error) {
	user, err := b.userRepositoryInterface.Create(ctx, createReq.User)
	if err != nil {
		return nil, err
	}

	return &models.CreateUserResp{
		Message: "record created sucessfully",
		User:    user,
	}, err
}
func (b *basicUsersService) GetAllUser(ctx context.Context) (allRecordResp []*models.User, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetAllUsers")
	defer span.Finish()
	span.LogFields(
		log.String("event", "soft error"),
		log.String("type", "cache timeout"),
		log.Int("waited.millis", 1500))
	allRecordResp, err = b.userRepositoryInterface.All(ctx)

	return allRecordResp, err
}
func (b *basicUsersService) UpdateUser(ctx context.Context, upadteReq models.User) (updateResp *models.User, err error) {
	updateResp, err = b.userRepositoryInterface.Update(ctx, upadteReq)
	return updateResp, err
}
func (b *basicUsersService) DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error) {
	deleteResp, err = b.userRepositoryInterface.Delete(ctx, id)
	return deleteResp, err
}
func (b *basicUsersService) GetUser(ctx context.Context, id string) (createResp *models.User, err error) {

	fmt.Println("id", id)

	span, ctx := opentracing.StartSpanFromContext(ctx, "Get User by Id")
	defer span.Finish()
	span.LogFields(
		log.String("event", "soft error"),
		log.String("type", "cache timeout"),
		log.Int("waited.millis", 1500))
	createResp, err = b.userRepositoryInterface.Get(ctx, id)

	return createResp, err

}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService(userRepositoryInterface repository.UserRepositoryInterface) UsersService {
	return &basicUsersService{userRepositoryInterface: userRepositoryInterface}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware, userRepositoryInterface repository.UserRepositoryInterface) UsersService {
	var svc UsersService = NewBasicUsersService(userRepositoryInterface)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
