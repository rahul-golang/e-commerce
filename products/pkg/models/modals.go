package models

import (
	"github.com/jinzhu/gorm"
)

//Product Feilds
type Product struct {
	gorm.Model
	ProductName            string  `json:"products_name"`
	ProductsQuantity       uint64  `json:"products_quantity"`
	ProductsModel          string  `json:"products_model"`
	ProductsImage          string  `json:"products_image"`
	ProductsPrice          float64 `json:"products_price"`
	ProductWeight          float64 `json:"products_weight"`
	ProductWeightUnit      string  `json:"products_weight_unit"`
	ProductsStatus         string  `json:"products_status"`
	ProductLikes           uint    `json:"products_likes"`
	ProductManufacturers   string  `json:"product_manufacturers"`
	ProductManufacturersID uint    `json:"product_manufacturers_id"`
}

type GetProductResp struct {
	Message string  `json:"message"`
	Product Product `json:"product"`
}

type CreateProductReq struct {
	Product Product `json:"product"`
}

type CreateProductResp struct {
	Message string   `json:"message"`
	Product *Product `json:"product"`
}

type DeleteProductResp struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

type UpdateProductReq struct {
	Product Product `json:"product"`
}

type UpdateProductResp struct {
	Message string  `json:"message"`
	Product Product `json:"product"`
}

type GetAllProductResp struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Product []*Product `json:"product"`
}
type UpdateStockReq struct {
	AddStock    uint64 `json:"add_stock"`
	RemoveStock uint64 `json:"remove_stock"`
	ProductID   uint64 `json:"product_id"`
}

type UpdateStockResp struct {
	Message string `json:"message"`
}
