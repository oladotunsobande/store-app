package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Personnel Schema definition of personnel table
type Personnel struct {
	gorm.Model
	UID            string    `gorm:"type:varchar(45);primaryKey;uniqueIndex" binding:"required" json:"uid"`
	FirstName      string    `gorm:"type:varchar(20)" binding:"required" json:"first_name"`
	LastName       string    `gorm:"type:varchar(20)" binding:"required" json:"last_name"`
	Email          string    `gorm:"type:varchar(65)" binding:"required" json:"email"`
	PhoneNumber    string    `gorm:"type:varchar(20)" binding:"required" json:"phone_number"`
	Password       string    `gorm:"type:varchar(80)" binding:"required" json:"password"`
	ResetKey       string    `gorm:"type:varchar(55)" binding:"required" json:"reset_key"`
	ResetKeyExpiry time.Time `json:"reset_key_expiry"`
	LastLogin      time.Time `json:"last_login"`
	HasAutoSecret  uint8     `gorm:"type:tinyint" json:"has_auto_secret"`
	Status         string    `gorm:"type:enum('active','inactive')" json:"status"`
	BusinessUID    string    `binding:"required" json:"business_uid"`
}

// BeforeCreate Hook for generating UUID and checking status
func (personnel *Personnel) BeforeCreate(tx *gorm.DB) (err error) {
	personnel.UID = uuid.New().String()
	return
}

// MigratePersonnelSchema Create table and relationships (if any)
func MigratePersonnelSchema(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Personnel{})
	return db
}
