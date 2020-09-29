package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order Schema definition of order table
type Order struct {
	gorm.Model
	UID              string            `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	OrderNumber      string            `gorm:"type:varchar(20);not null;autoIncrement" binding:"required" json:"order_number"`
	Amount           string            `gorm:"type:decimal(15,2);not null" binding:"required" json:"amount"`
	Quantity         uint32            `gorm:"type:int;not null" binding:"required" json:"quantity"`
	Status           string            `gorm:"type:enum('pending','confirmed','cancelled');not null" binding:"required" json:"status"`
	CustomerUID      string            `gorm:"not null" binding:"required" json:"business_uid"`
	PersonnelUID     string            `gorm:"not null" binding:"required" json:"personnel_uid"`
	Payment          Payment           `gorm:"foreignKey:OrderUID;references:UID" json:"payment"`
	OrderProducts    []OrderProduct    `gorm:"manytomany:order_products" json:"order_products"`
	OrderCollections []OrderCollection `gorm:"manytomany:order_collections" json:"order_collections"`
}

// BeforeCreate Hook for generating UUID
func (order *Order) BeforeCreate(tx *gorm.DB) error {
	order.UID = uuid.New().String()
	return nil
}

// OrderSchema Get order schema interface
func OrderSchema() *Order {
	return &Order{}
}
