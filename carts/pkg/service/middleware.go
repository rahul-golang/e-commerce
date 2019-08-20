package service

import (
	"context"
	models "gokit/ecommerse/carts/pkg/models"

	log "github.com/go-kit/kit/log"
)

type Middleware func(CartsService) CartsService

type loggingMiddleware struct {
	logger log.Logger
	next   CartsService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next CartsService) CartsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) AddToCart(ctx context.Context, addReq models.AddReq) (addResp *models.AddResp, err error) {
	defer func() {
		l.logger.Log("method", "AddToCart", "addReq", addReq, "addResp", addResp, "err", err)
	}()
	return l.next.AddToCart(ctx, addReq)
}
func (l loggingMiddleware) GetFromAll(ctx context.Context) (allRecordResp *models.GetAllResp, err error) {
	defer func() {
		l.logger.Log("method", "GetFromAll", "allRecordResp", allRecordResp, "err", err)
	}()
	return l.next.GetFromAll(ctx)
}
func (l loggingMiddleware) UpdateInCart(ctx context.Context, upadteReq models.UpdateReq) (updateResp *models.UpdateCartResp, err error) {
	defer func() {
		l.logger.Log("method", "UpdateInCart", "upadteReq", upadteReq, "updateResp", updateResp, "err", err)
	}()
	return l.next.UpdateInCart(ctx, upadteReq)
}
func (l loggingMiddleware) DeleteFromCart(ctx context.Context, id string) (deleteResp *models.DeleteResp, err error) {
	defer func() {
		l.logger.Log("method", "DeleteFromCart", "id", id, "deleteResp", deleteResp, "err", err)
	}()
	return l.next.DeleteFromCart(ctx, id)
}
func (l loggingMiddleware) GetFromCart(ctx context.Context, id string) (createResp *models.GetResp, err error) {
	defer func() {
		l.logger.Log("method", "GetFromCart", "id", id, "createResp", createResp, "err", err)
	}()
	return l.next.GetFromCart(ctx, id)
}
