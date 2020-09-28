package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Transaction Schema definition of transaction table
type Transaction struct {
	gorm.Model
	UID         string  `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Category    string  `gorm:"type:enum('DEBIT','CREDIT');not null" binding:"required" json:"category"`
	Reference   string  `gorm:"type:varchar(45);not null;index" binding:"required" json:"reference"`
	Narration   string  `gorm:"type:varchar(75);not null" binding:"required" json:"narration"`
	Amount      float64 `gorm:"type:decimal(15,2);not null" binding:"required" json:"amount"`
	BusinessUID string  `binding:"required" json:"business_uid"`
}

// BeforeCreate Hook for generating UUID
func (transaction *Transaction) BeforeCreate(tx *gorm.DB) {
	transaction.UID = uuid.New().String()
}

// MigrateTransactionSchema Create table and relationships (if any)
func MigrateTransactionSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Transaction{})
	return db
}
