package models

import "time"

type Orders struct {
	OrdersID                  uint64          `json:"orders_id" gorm:"primary_key;AUTO_INCREMENT"`
	TotalTax                  int64           `json:"total_tax"`
	CustomersID               uint64          `json:"customers_id"`
	CustomersName             string          `json:"customers_name"`
	CustomersCompany          string          `json:"customers_company"`
	CustomersStreetAddress    string          `json:"customers_street_address"`
	CustomersSuburb           string          `json:"customers_suburb"`
	CustomersCity             string          `json:"customers_city"`
	CustomersPostCode         string          `json:"customers_postcode"`
	CustomersState            string          `json:"customers_state"`
	CustomersCountry          string          `json:"customers_country"`
	CustomersTelephone        string          `json:"customers_telephone"`
	Email                     string          `json:"email"`
	CustomersAddressFormatID  string          `json:"customers_address_format_id"`
	DeliveryName              string          `json:"delivery_name"`
	DeliveryCompany           string          `json:"delivery_company"`
	DeliveryStreetAddress     string          `json:"delivery_street_address"`
	DeliverySuburb            string          `json:"delivery_suburb"`
	DeliveryCity              string          `json:"delivery_city"`
	DeliveryPostcode          string          `json:"delivery_postcode"`
	DeliveryState             string          `json:"delivery_state"`
	DeliveryCountry           string          `json:"delivery_country"`
	DeliveryAddressFormatID   uint64          `json:"delivery_address_format_id"`
	BillingName               string          `json:"billing_name"`
	BillingCompany            string          `json:"billing_company"`
	BillingStreetAddress      string          `json:"billing_street_address"`
	BillingSuburb             string          `json:"billing_suburb"`
	BillingCity               string          `json:"billing_city"`
	BillingPostCode           string          `json:"billing_postcode"`
	BillingState              string          `json:"billing_state"`
	BillingCountry            string          `json:"billing_country"`
	BillingAddresFormatID     uint64          `json:"billing_address_format_id"`
	PaymentMethod             string          `json:"payment_method"`
	CcType                    string          `json:"cc_type"`
	CcOwner                   string          `json:"cc_owner"`
	CcNumber                  string          `json:"cc_number"`
	CcExpires                 string          `json:"cc_expires"`
	LastModified              time.Time       `json:"last_modified"`
	DatePurchased             time.Time       `json:"date_purchased"`
	OrdersDateFinished        time.Time       `json:"orders_date_finished"`
	Currency                  string          `json:"currency"`
	CurrencyValue             float64         `json:"currency_value"`
	OrderPrice                float64         `json:"order_price"`
	ShippingCost              float64         `json:"shipping_cost"`
	ShippingMethod            string          `json:"shipping_method"`
	ShippingDuration          time.Time       `json:"shipping_duration"`
	OrderInformation          string          `json:"order_information"`
	IsSeen                    bool            `json:"is_seen"`
	CouponCode                string          `json:"coupon_code"`
	CouponAmount              float64         `json:"coupon_amount"`
	ExcludePoductIDs          uint64          `json:"exclude_product_ids"`
	ProductCategories         string          `json:"product_categories"`
	ExcludedProductCategories string          `json:"excluded_product_categories"`
	FreeShipping              bool            `json:"free_shipping"`
	OrderedSource             string          `json:"ordered_source"`
	DeliveryPhone             string          `json:"delivery_phone"`
	BillingPhone              string          `json:"billing_phone"`
	TransactionID             string          `json:"transaction_id"`
	OrderProducts             []OrderProducts `json:"products" gorm:"foreignkey:orders_id"`
	OrderStatus               uint64          `json:"order_status" gorm:"DEFAULT=1;foreignkey:order_status"`
}

//OrderProducts is
type OrderProducts struct {
	OrderProductID uint64  `json:"order_product_id" gorm:"primary_key;AUTO_INCREMENT"`
	OrdersID       uint64  `json:"order_id"`
	ProductID      uint64  `json:"product_id"`
	ProductModel   string  `json:"product_model"`
	ProductName    string  `json:"product_name"`
	ProductPrice   float64 `json:"product_price"`
	ProductTax     float64 `json:"product_tax"`
	ProductQuntity uint64  `json:"product_quntity"`
}

type OrderStatus struct {
	OrderStatusID uint64 `json:"order_status_id" gorm:"primary_key;AUTO_INCREMENT"`
	OrderStatus   string `json:"order_status"`
}

type CrateOrderReq struct {
	Orders Orders `json:"order"`
}

type CrateOrderResp struct {
	Message string  `json:"message"`
	Orders  *Orders `json:"order"`
}
