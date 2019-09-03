package models

import (
	"github.com/jinzhu/gorm"
)

//Product Feilds
type Product struct {
	gorm.Model
	ProductsQuantity       uint64             `json:"products_quantity" gorm:"default:0"`
	ProductsModel          string             `json:"products_model" `
	ProductsImageURL       string             `json:"products_image" gorm:"default:0"`
	ProductsPrice          float64            `json:"products_price"  gorm:"not null"`
	ProductWeight          float64            `json:"products_weight"   gorm:"not null"`
	ProductWeightUnit      string             `json:"products_weight_unit"  gorm:"not null"`
	ProductsStatus         string             `json:"products_status"  gorm:"not null"`
	ProductLikes           uint               `json:"products_likes" `
	ProductManufacturersID uint               `json:"product_manufacturers_id"`
	ProductsOrdered        uint32             `json:"products_ordered" gorm:"default:0" `
	LowLimit               uint16             `json:"low_limit" gorm:"default:0"`
	IsFeature              bool               `json:"is_feature" gorm:"default:0"`
	ProductsSlug           string             `json:"products_slug" gorm:"not null"`
	ProductsType           string             `json:"products_type" `
	ProductsMinOrder       uint64             `json:"products_min_order" gorm:"default:1"`
	ProductsMaxStock       uint64             `json:"products_max_stock" gorm:"default:0"`
	AdminApproval          string             `json:"admin_approval" gorm:"default:'pending'"`
	RejectReason           string             `json:"reject_reason"`
	AddedBy                string             `json:"added_by"  gorm:"not null"`
	ProductDescription     ProductDescription `json:"products_description" gorm:"foreignkey:products_id"`
	ProductImages          []ProductImages    `json:"products" gorm:"foreignkey:products_id"`
}

type ProductDescription struct {
	ProductsID          uint   `json:"products_id" gorm:"UNIQUE"`
	LanguagID           uint   `json:"language_id"`
	ProductsName        string `json:"products_name"`
	ProductsDescription string `json:"products_description"`
	ProductsURL         string `json:"products_url"`
	ProductsViewed      string `json:"products_viewed"`
	ProductsBanner      string `json:"products_left_banner"`
}

type ProductImages struct {
	gorm.Model
	ImageURL   string `gorm:"size:255" json:"image_url"`
	ImageDesc  string `gorm:"size:255" json:"image_desc"`
	ProductsID uint   `json:"products_id"` //gorm:"ForeignKey:products_id"
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
	AddStock    bool            `json:"add_stock"`
	RemoveStock bool            `json:"remove_stock"`
	Products    []*ProductStock `json:"products"`
}
type ProductStock struct {
	ProductID        uint64 `json:"product_id"`
	ProductsQuantity int64  `json:"product_quantity"`
}

type UpdateStockResp struct {
	Message string `json:"message"`
}
