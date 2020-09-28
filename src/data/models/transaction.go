package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Transaction Schema definition of transaction table
type Transaction struct {
	gorm.Model
	UID         string  `gorm:"type:varchar(45);primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Category    string  `gorm:"type:enum('DEBIT','CREDIT')" binding:"required" json:"category"`
	Reference   string  `gorm:"type:varchar(45)" binding:"required" json:"reference"`
	Narration   string  `gorm:"type:varchar(75)" binding:"required" json:"narration"`
	Amount      float64 `binding:"required" json:"amount"`
	BusinessUID string  `binding:"required" json:"business_uid"`
}

// BeforeCreate Hook for generating UUID
func (transaction *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	transaction.UID = uuid.New().String()
	return
}

// MigrateTransactionSchema Create table and relationships (if any)
func MigrateTransactionSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Transaction{})
	return db
}
