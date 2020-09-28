package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Category Schema definition of product categories table
type Category struct {
	gorm.Model
	UID        string       `gorm:"type:varchar(45);primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name       string       `gorm:"type:varchar(45)" binding:"required" json:"name"`
	Product    []Product    `gorm:"foreignKey:CategoryUID;references:UID" json:"products"`
	Collection []Collection `gorm:"foreignKey:CategoryUID;references:UID" json:"collections"`
}

// BeforeCreate Hook for generating UUID
func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	category.UID = uuid.New().String()
	return
}

// MigrateCategorySchema Create table and relationships (if any)
func MigrateCategorySchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Category{})
	return db
}
