package service

import (
	"context"
	"fmt"
	"gokit/ecommerse/products/pkg/models"
	"gokit/ecommerse/products/pkg/repository"
)

// ProductsService describes the service.
type ProductsService interface {
	CreateProduct(ctx context.Context, createReq models.CreateProductReq) (createResp *models.CreateProductResp, err error)
	GetAllProduct(ctx context.Context) (allRecordResp []*models.Product, err error)
	UpdateProduct(ctx context.Context, upadteReq models.Product) (updateResp *models.Product, err error)
	DeleteProduct(ctx context.Context, id string) (deleteResp *models.DeleteProductResp, err error)
	GetProduct(ctx context.Context, id string) (createResp *models.Product, err error)
	UpdateProductStock(ctx context.Context, updateStockReq models.UpdateStockReq) (updateStockResp *models.UpdateStockResp, err error)
}

type basicProductsService struct {
	productRepository repository.ProductRepositoryInterface
}

func (b *basicProductsService) CreateProduct(ctx context.Context, createReq models.CreateProductReq) (*models.CreateProductResp, error) {
	fmt.Println(createReq)
	createResp := &models.CreateProductResp{}
	productRepo, err := b.productRepository.Create(ctx, createReq.Product)
	if err != nil {
		createResp.Product = nil
		createResp.Message = err.Error()
		return createResp, err
	}
	fmt.Println(productRepo)
	createResp.Product = productRepo
	createResp.Message = "Records"
	return createResp, nil
}
func (b *basicProductsService) GetAllProduct(ctx context.Context) ([]*models.Product, error) {
	resp, err := b.productRepository.All(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}
func (b *basicProductsService) UpdateProduct(ctx context.Context, upadteReq models.Product) (*models.Product, error) {
	resp, err := b.productRepository.Update(ctx, upadteReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (b *basicProductsService) DeleteProduct(ctx context.Context, id string) (*models.DeleteProductResp, error) {

	resp, err := b.productRepository.Delete(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}
func (b *basicProductsService) GetProduct(ctx context.Context, id string) (*models.Product, error) {

	resp, err := b.productRepository.Get(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}

// NewBasicProductsService returns a naive, stateless implementation of ProductsService.
func NewBasicProductsService(productRepository repository.ProductRepositoryInterface) ProductsService {
	return &basicProductsService{productRepository: productRepository}
}

// New returns a ProductsService with all of the expected middleware wired in.
func New(middleware []Middleware, productRepository repository.ProductRepositoryInterface) ProductsService {
	var svc ProductsService = NewBasicProductsService(productRepository)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicProductsService) UpdateProductStock(ctx context.Context, updateStockReq models.UpdateStockReq) (updateStockResp *models.UpdateStockResp, err error) {

	updateStockResp, err = b.productRepository.UpdateProductStock(ctx, updateStockReq)

	return
}
