package models

import (
	"gorm.io/gorm"
)

// OrderProduct Schema definition of order_products jointable
type OrderProduct struct {
	gorm.Model
	OrderUID   string `binding:"required" json:"order_uid"`
	ProductUID string `binding:"required" json:"product_uid"`
	Quantity   uint16 `gorm:"type:int;not null" binding:"required" json:"quantity"`
}

// MigrateOrderProductSchema Create table and relationships (if any)
func MigrateOrderProductSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&OrderProduct{})
	return db
}
