package service

import (
	"context"
	models "gokit/ecommerse/products/pkg/models"

	log "github.com/go-kit/kit/log"
)

type Middleware func(ProductsService) ProductsService

type loggingMiddleware struct {
	logger log.Logger
	next   ProductsService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ProductsService) ProductsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateProduct(ctx context.Context, createReq models.CreateProductReq) (createResp *models.CreateProductResp, err error) {
	defer func() {
		l.logger.Log("method", "CreateProduct", "createReq", createReq, "createResp", createResp, "err", err)
	}()
	return l.next.CreateProduct(ctx, createReq)
}
func (l loggingMiddleware) GetAllProduct(ctx context.Context) (allRecordResp []*models.Product, err error) {
	defer func() {
		l.logger.Log("method", "GetAllProduct", "allRecordResp", allRecordResp, "err", err)
	}()
	return l.next.GetAllProduct(ctx)
}
func (l loggingMiddleware) UpdateProduct(ctx context.Context, upadteReq models.Product) (updateResp *models.Product, err error) {
	defer func() {
		l.logger.Log("method", "UpdateProduct", "upadteReq", upadteReq, "updateResp", updateResp, "err", err)
	}()
	return l.next.UpdateProduct(ctx, upadteReq)
}
func (l loggingMiddleware) DeleteProduct(ctx context.Context, id string) (deleteResp *models.DeleteProductResp, err error) {
	defer func() {
		l.logger.Log("method", "DeleteProduct", "id", id, "deleteResp", deleteResp, "err", err)
	}()
	return l.next.DeleteProduct(ctx, id)
}
func (l loggingMiddleware) GetProduct(ctx context.Context, id string) (createResp *models.Product, err error) {
	defer func() {
		l.logger.Log("method", "GetProduct", "id", id, "createResp", createResp, "err", err)
	}()
	return l.next.GetProduct(ctx, id)
}

func (l loggingMiddleware) UpdateProductStock(ctx context.Context, updateStockReq models.UpdateStockReq) (updateStockResp *models.UpdateStockResp, err error) {
	defer func() {
		l.logger.Log("method", "UpdateProductStock", "updateStockReq", updateStockReq, "updateStockResp", updateStockResp, "err", err)
	}()
	return l.next.UpdateProductStock(ctx, updateStockReq)
}
