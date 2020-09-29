package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Personnel Schema definition of personnel table
type Personnel struct {
	gorm.Model
	UID            string    `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	FirstName      string    `gorm:"type:varchar(20);not null;index" binding:"required" json:"first_name"`
	LastName       string    `gorm:"type:varchar(20);not null;index" binding:"required" json:"last_name"`
	Email          string    `gorm:"type:varchar(65);not null;index" binding:"required" json:"email"`
	PhoneNumber    string    `gorm:"type:varchar(20)" binding:"required" json:"phone_number,omitempty"`
	Password       string    `gorm:"type:varchar(80);not null" binding:"required" json:"password"`
	ResetKey       string    `gorm:"type:varchar(55)" binding:"required" json:"reset_key,omitempty"`
	ResetKeyExpiry time.Time `json:"reset_key_expiry,omitempty"`
	LastLogin      time.Time `json:"last_login,omitempty"`
	HasAutoSecret  uint8     `gorm:"type:tinyint;not null;default:0" json:"has_auto_secret"`
	Status         string    `gorm:"type:enum('active','inactive');not null;default:'active'" json:"status"`
	BusinessUID    string    `gorm:"not null" binding:"required" json:"business_uid"`
	Order          []Order   `gorm:"foreignKey:PersonnelUID;references:UID" json:"orders"`
}

// BeforeCreate Hook for generating UUID
func (personnel *Personnel) BeforeCreate(tx *gorm.DB) error {
	personnel.UID = uuid.New().String()
	return nil
}

// PersonnelSchema Get personnel schema interface
func PersonnelSchema() *Personnel {
	return &Personnel{}
}
