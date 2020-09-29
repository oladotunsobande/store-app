package models

import (
	"gorm.io/gorm"
)

// OrderCollection Schema definition of order_collections jointable
type OrderCollection struct {
	gorm.Model
	OrderUID      string `gorm:"not null" binding:"required" json:"order_uid"`
	CollectionUID string `gorm:"not null" binding:"required" json:"collection_uid"`
	Quantity      uint16 `gorm:"type:int;not null" binding:"required" json:"quantity"`
}

// OrderCollectionSchema Get ordercollection schema interface
func OrderCollectionSchema() *OrderCollection {
	return &OrderCollection{}
}
