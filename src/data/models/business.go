package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Business Schema definition of business table
type Business struct {
	gorm.Model
	UID                string        `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name               string        `gorm:"type:varchar(60);not null;index" binding:"required" json:"name"`
	Province           string        `gorm:"type:varchar(30);not null" binding:"required" json:"province"`
	Country            string        `gorm:"type:varchar(30);not null" binding:"required" json:"country"`
	SubscriptionExpiry time.Time     `gorm:"type:datetime;not null" binding:"required" json:"subscription_expiry"`
	Status             string        `gorm:"type:enum('active','inactive');not null" binding:"required" json:"status"`
	SubscriptionUID    string        `binding:"required" json:"subscription_uid"`
	Products           []Product     `gorm:"foreignKey:BusinessUID;references:UID" json:"products"`
	Collections        []Collection  `gorm:"foreignKey:BusinessUID;references:UID" json:"collections"`
	Vendors            []Vendor      `gorm:"foreignKey:BusinessUID;references:UID" json:"vendors"`
	Transactions       []Transaction `gorm:"foreignKey:BusinessUID;references:UID" json:"transactions"`
	Personnels         []Personnel   `gorm:"foreignKey:BusinessUID;references:UID" json:"personnel"`
}

// BeforeCreate Hook for generating UUID
func (business *Business) BeforeCreate(tx *gorm.DB) {
	business.UID = uuid.New().String()
}

// MigrateBusinessSchema Create table and relationships (if any)
func MigrateBusinessSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Business{})
	return db
}
