package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Payment Schema definition of payment table
type Payment struct {
	gorm.Model
	UID           string `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	InvoiceNumber string `gorm:"type:varchar(20);not null;index" binding:"required" json:"invoice_number"`
	Medium        string `gorm:"type:enum('cash','debit-card','bank-transfer','other');not null" binding:"required" json:"medium"`
	Status        string `gorm:"type:enum('pending','completed');not null" binding:"required" json:"status"`
	OrderUID      string `binding:"required" json:"order_uid"`
}

// BeforeCreate Hook for generating UUID
func (payment *Payment) BeforeCreate(tx *gorm.DB) {
	payment.UID = uuid.New().String()
}

// MigratePaymentSchema Create table and relationships (if any)
func MigratePaymentSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Payment{})
	return db
}
