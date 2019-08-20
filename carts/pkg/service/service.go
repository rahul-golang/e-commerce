package service

import (
	"context"
	"fmt"
	"gokit/ecommerse/carts/pkg/models"
	"gokit/ecommerse/carts/pkg/repository"
)

// CartsService describes the service.
type CartsService interface {
	AddToCart(ctx context.Context, addReq models.AddReq) (addResp *models.AddResp, err error)
	GetFromAll(ctx context.Context) (allRecordResp *models.GetAllResp, err error)
	UpdateInCart(ctx context.Context, upadteReq models.UpdateReq) (updateResp *models.UpdateCartResp, err error)
	DeleteFromCart(ctx context.Context, id string) (deleteResp *models.DeleteResp, err error)
	GetFromCart(ctx context.Context, id string) (createResp *models.GetResp, err error)
}

type basicCartsService struct {
	cartRepositoryInterface repository.CartRepositoryInterface
}

func (b *basicCartsService) AddToCart(ctx context.Context, addReq models.AddReq) (*models.AddResp, error) {
	resp, err := b.cartRepositoryInterface.AddToCart(ctx, addReq.Cart)
	if err != nil {
		return nil, err
	}
	return &models.AddResp{
		Message: "Added Successfully",
		Cart:    resp,
	}, err
}

func (b *basicCartsService) GetFromAll(ctx context.Context) (*models.GetAllResp, error) {
	resp, err := b.cartRepositoryInterface.GetFromAll(ctx)
	if err != nil {
		return nil, err
	}
	return &models.GetAllResp{
		Message: "Records Retrived Successfully",
		Cart:    resp,
	}, err
}
func (b *basicCartsService) UpdateInCart(ctx context.Context, upadteReq models.UpdateReq) (*models.UpdateCartResp, error) {
	resp, err := b.cartRepositoryInterface.UpdateInCart(ctx, upadteReq.Cart)
	if err != nil {
		return nil, err
	}
	return &models.UpdateCartResp{
		Message: "Updated Successfully",
		Cart:    resp,
	}, err
}
func (b *basicCartsService) DeleteFromCart(ctx context.Context, id string) (*models.DeleteResp, error) {
	resp, err := b.cartRepositoryInterface.DeleteFromCart(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.DeleteResp{
		Message: "Deleted Successfully",
		ID:      string(resp.ID),
	}, err
}
func (b *basicCartsService) GetFromCart(ctx context.Context, id string) (*models.GetResp, error) {
	resp, err := b.cartRepositoryInterface.GetFromCart(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return &models.GetResp{
		Message: "Record Retrived Successfully",
		Cart:    resp,
	}, err
}

// NewBasicCartsService returns a naive, stateless implementation of CartsService.
func NewBasicCartsService(repoInterface repository.CartRepositoryInterface) CartsService {
	return &basicCartsService{repoInterface}
}

// New returns a CartsService with all of the expected middleware wired in.
func New(middleware []Middleware, repoInterface repository.CartRepositoryInterface) CartsService {
	var svc CartsService = NewBasicCartsService(repoInterface)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
