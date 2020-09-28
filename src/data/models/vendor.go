package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Vendor Schema definition of vendor table
type Vendor struct {
	gorm.Model
	UID         string `gorm:"type:varchar(45);primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name        string `gorm:"type:varchar(50)" binding:"required" json:"name"`
	Address     string `gorm:"type:varchar(65)" binding:"required" json:"address"`
	Province    string `gorm:"type:varchar(20)" binding:"required" json:"province"`
	Country     string `gorm:"type:varchar(45)" binding:"required" json:"country"`
	Email       string `gorm:"type:varchar(55)" binding:"required" json:"email"`
	PhoneNumber string `gorm:"type:varchar(20)" binding:"required" json:"phone_number"`
	Status      string `gorm:"type:enum('active','inactive')" binding:"required" json:"status"`
	BusinessUID string `binding:"required" json:"business_uid"`
}

// BeforeCreate Hook for generating UUID
func (vendor *Vendor) BeforeCreate(tx *gorm.DB) (err error) {
	vendor.UID = uuid.New().String()
	return
}

// MigrateVendorSchema Create table and relationships (if any)
func MigrateVendorSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Vendor{})
	return db
}
