package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product Schema definition of product table
type Product struct {
	gorm.Model
	UID            string  `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Brand          string  `gorm:"type:varchar(45);not null;index" binding:"required" json:"brand"`
	BarCode        string  `gorm:"type:varchar(25);not null;index" binding:"required" json:"barcode"`
	Name           string  `gorm:"type:varchar(45);not null;index" binding:"required" json:"name"`
	Price          float64 `gorm:"type:decimal(15,2);not null" binding:"required" json:"price"`
	TotalStock     uint32  `gorm:"type:int;not null" binding:"required" json:"total_stock"`
	AvailableStock uint32  `gorm:"type:int;not null" binding:"required" json:"available_stock"`
	LockedStock    uint32  `gorm:"type:int;default:0" binding:"required" json:"locked_stock"`
	Status         string  `gorm:"type:enum('available','unavailable','exhausted');not null" binding:"required" json:"status"`
	CategoryUID    string  `binding:"required" json:"category_uid"`
	BusinessUID    string  `binding:"required" json:"business_uid"`
}

// BeforeCreate Hook for generating UUID
func (product *Product) BeforeCreate(tx *gorm.DB) {
	product.UID = uuid.New().String()
}

// MigrateProductSchema Create table and relationships (if any)
func MigrateProductSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{})
	return db
}
