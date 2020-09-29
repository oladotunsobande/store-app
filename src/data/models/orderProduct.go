package models

import (
	"gorm.io/gorm"
)

// OrderProduct Schema definition of order_products jointable
type OrderProduct struct {
	gorm.Model
	OrderUID   string `gorm:"not null" binding:"required" json:"order_uid"`
	ProductUID string `gorm:"not null" binding:"required" json:"product_uid"`
	Quantity   uint16 `gorm:"type:int;not null" binding:"required" json:"quantity"`
}

// OrderProductSchema Get orderproduct schema interface
func OrderProductSchema() *OrderProduct {
	return &OrderProduct{}
}
