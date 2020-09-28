package models

import (
	"gorm.io/gorm"
)

// OrderCollection Schema definition of order_collections jointable
type OrderCollection struct {
	gorm.Model
	OrderUID      string `binding:"required" json:"order_uid"`
	CollectionUID string `binding:"required" json:"collection_uid"`
	Quantity      uint16 `gorm:"type:int;not null" binding:"required" json:"quantity"`
}

// MigrateOrderCollectionSchema Create table and relationships (if any)
func MigrateOrderCollectionSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&OrderCollection{})
	return db
}
