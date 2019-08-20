package models

import (
	"github.com/jinzhu/gorm"
)

//Cart Feilds
type Cart struct {
	gorm.Model
	UserID        string `json:"user_id"`
	ProductID     string `json:"product_id"`
	CartsQuantity string `json:"carts_quantity"`
	CartsPrice    string `json:"carts_price"`
	CartsStatus   string `json:"carts_status"`
}

type GetResp struct {
	Message string `json:"message"`
	Cart    *Cart  `json:"cart"`
}

type AddReq struct {
	Cart Cart `json:"cart"`
}

type AddResp struct {
	Message string `json:"message"`
	Cart    *Cart  `json:"cart"`
}

type DeleteResp struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

type UpdateReq struct {
	Cart Cart `json:"cart"`
}

type UpdateCartResp struct {
	Message string `json:"message"`
	Cart    *Cart  `json:"cart"`
}

type GetAllResp struct {
	Message string  `json:"message"`
	Cart    []*Cart `json:"cart"`
}
