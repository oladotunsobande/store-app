package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Parameter Schema definition of parameter table
type Parameter struct {
	gorm.Model
	UID         string      `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	DisplayName string      `gorm:"type:varchar(65);not null" binding:"required" json:"display_name"`
	Name        string      `gorm:"type:varchar(45);not null;index" binding:"required" json:"name"`
	Value       interface{} `gorm:"type:json;not null" binding:"required" json:"value"`
}

// BeforeCreate Hook for generating UUID
func (parameter *Parameter) BeforeCreate(tx *gorm.DB) {
	parameter.UID = uuid.New().String()
}

// MigrateParameterSchema Create table and relationships (if any)
func MigrateParameterSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Parameter{})
	return db
}
