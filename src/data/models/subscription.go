package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Subscription Schema definition of subscription table
type Subscription struct {
	gorm.Model
	UID      string  `gorm:"type:varchar(45);primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name     string  `gorm:"type:varchar(45)" binding:"required" json:"name"`
	Price    float64 `gorm:"type:decimal(15,2)" binding:"required" json:"price"`
	Duration uint16  `gorm:"type:int" binding:"required" json:"duration"`
	IsTrial  uint8   `binding:"required" json:"is_trial"`
}

// BeforeCreate Hook for generating UUID
func (subscription *Subscription) BeforeCreate(tx *gorm.DB) (err error) {
	subscription.UID = uuid.New().String()
	return
}

// MigrateSubscriptionSchema Create table and relationships (if any)
func MigrateSubscriptionSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Subscription{})
	return db
}
