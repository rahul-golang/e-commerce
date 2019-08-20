package models

import (
	"github.com/jinzhu/gorm"
)

//Product Feilds
type Product struct {
	gorm.Model
	ProductName          string `json:"products_name"`
	ProductsQuantity     string `json:"products_quantity"`
	ProductsModel        string `json:"products_model"`
	ProductsImage        string `json:"products_image"`
	ProductsPrice        string `json:"products_price"`
	ProductWeight        string `json:"products_weight"`
	ProductWeightUnit    string `json:"products_weight_unit"`
	ProductsStatus       string `json:"products_status"`
	ProductLike          string `json:"products_like"`
	ProductManufacturers string `json:"product_manufacturers"`
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
