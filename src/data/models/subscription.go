package models

import (
	"gorm.io/gorm"
)

// Subscription Schema definition of subscription table
type Subscription struct {
	gorm.Model
	UID      string   `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name     string   `gorm:"type:varchar(45);not null;index" binding:"required" json:"name"`
	Price    float64  `gorm:"type:decimal(15,2);not null:default:0.00" binding:"required" json:"price"`
	Duration uint16   `gorm:"type:int;not null" binding:"required" json:"duration"`
	IsTrial  uint8    `gorm:"type:tinyint;not null:default:0" binding:"required" json:"is_trial"`
	Business Business `gorm:"foreignKey:SubscriptionUID;references:UID" json:"business"`
}

// BeforeCreate Hook for generating UUID
/*func (subscription *Subscription) BeforeCreate(tx *gorm.DB) error {
	subscription.UID = uuid.New().String()
	return nil
}*/

// SubscriptionSchema Get subscription schema interface
func SubscriptionSchema() *Subscription {
	return &Subscription{}
}
