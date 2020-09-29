package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Vendor Schema definition of vendor table
type Vendor struct {
	gorm.Model
	UID         string `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name        string `gorm:"type:varchar(50);not null;index" binding:"required" json:"name"`
	Address     string `gorm:"type:varchar(65);not null" binding:"required" json:"address"`
	Province    string `gorm:"type:varchar(20);not null" binding:"required" json:"province"`
	Country     string `gorm:"type:varchar(45);not null" binding:"required" json:"country"`
	Email       string `gorm:"type:varchar(55);not null;index" binding:"required" json:"email"`
	PhoneNumber string `gorm:"type:varchar(20);not null;index" binding:"required" json:"phone_number"`
	Status      string `gorm:"type:enum('active','inactive');not null" binding:"required" json:"status"`
	BusinessUID string `gorm:"not null" binding:"required" json:"business_uid"`
}

// BeforeCreate Hook for generating UUID
func (vendor *Vendor) BeforeCreate(tx *gorm.DB) error {
	vendor.UID = uuid.New().String()
	return nil
}

// VendorSchema Get vendor schema interface
func VendorSchema() *Vendor {
	return &Vendor{}
}
