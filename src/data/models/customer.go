package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Customer Schema definition of customer table
type Customer struct {
	gorm.Model
	UID         string  `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	FirstName   string  `gorm:"type:varchar(20);not null;index" binding:"required" json:"first_name"`
	LastName    string  `gorm:"type:varchar(20);not null;index" binding:"required" json:"last_name"`
	Email       string  `gorm:"type:varchar(65);index" binding:"required" json:"email"`
	PhoneNumber string  `gorm:"type:varchar(20);not null;index" binding:"required" json:"phone_number"`
	Order       []Order `gorm:"foreignKey:CustomerUID;references:UID" json:"orders"`
}

// BeforeCreate Hook for generating UUID
func (customer *Customer) BeforeCreate(tx *gorm.DB) {
	customer.UID = uuid.New().String()
}

// MigrateCustomerSchema Create table and relationships (if any)
func MigrateCustomerSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Customer{})
	return db
}
