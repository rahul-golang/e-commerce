package service

import (
	"context"
	models "gokit/ecommerse/users/pkg/models"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UsersService) UsersService

type loggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UsersService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UsersService) UsersService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateUser(ctx context.Context, createReq models.CreateUserReq) (createResp *models.CreateUserResp, err error) {
	defer func() {
		l.logger.Log("method", "CreateUser", "createReq", createReq, "createResp", createResp, "err", err)
	}()
	return l.next.CreateUser(ctx, createReq)
}
func (l loggingMiddleware) GetAllUser(ctx context.Context) (allRecordResp []*models.User, err error) {
	defer func() {
		l.logger.Log("method", "GetAllUser", "allRecordResp", allRecordResp, "err", err)
	}()
	return l.next.GetAllUser(ctx)
}
func (l loggingMiddleware) UpdateUser(ctx context.Context, upadteReq models.User) (updateResp *models.User, err error) {
	defer func() {
		l.logger.Log("method", "UpdateUser", "upadteReq", upadteReq, "updateResp", updateResp, "err", err)
	}()
	return l.next.UpdateUser(ctx, upadteReq)
}
func (l loggingMiddleware) DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error) {
	defer func() {
		l.logger.Log("method", "DeleteUser", "id", id, "deleteResp", deleteResp, "err", err)
	}()
	return l.next.DeleteUser(ctx, id)
}
func (l loggingMiddleware) GetUser(ctx context.Context, id string) (createResp *models.User, err error) {
	defer func() {
		l.logger.Log("method", "GetUser", "id", id, "createResp", createResp, "err", err)
	}()
	return l.next.GetUser(ctx, id)
}
