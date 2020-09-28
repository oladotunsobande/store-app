package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Business Schema definition of business table
type Business struct {
	gorm.Model
	UID                string        `gorm:"type:varchar(45);primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name               string        `gorm:"type:varchar(60)" binding:"required" json:"name"`
	Province           string        `gorm:"type:varchar(30)" binding:"required" json:"province,omitempty"`
	Country            string        `gorm:"type:varchar(30)" binding:"required" json:"country"`
	SubscriptionExpiry time.Time     `binding:"required" json:"subscription_expiry,omitempty"`
	Status             string        `gorm:"type:enum('active','inactive')" binding:"required" json:"status"`
	SubscriptionUID    string        `binding:"required" json:"subscription_uid"`
	Product            []Product     `gorm:"foreignKey:BusinessUID;references:UID" json:"products"`
	Collection         []Collection  `gorm:"foreignKey:BusinessUID;references:UID" json:"collections"`
	Vendor             []Vendor      `gorm:"foreignKey:BusinessUID;references:UID" json:"vendors"`
	Transaction        []Transaction `gorm:"foreignKey:BusinessUID;references:UID" json:"transactions"`
	Personnel          []Personnel   `gorm:"foreignKey:BusinessUID;references:UID" json:"personnel"`
}

// BeforeCreate Hook for generating UUID and checking status
func (business *Business) BeforeCreate(tx *gorm.DB) (err error) {
	business.UID = uuid.New().String()

	if business.Status != "active" && business.Status != "inactive" {
		err = errors.New("Invalid business status")
	}

	return
}

// MigrateBusinessSchema Create table and relationships (if any)
func MigrateBusinessSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Business{})
	return db
}
