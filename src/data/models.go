package database

import (
	"time"

	"gorm.io/gorm"
)

// Subscription Schema definition of subscription table
type Subscription struct {
	gorm.Model
	UID      string     `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name     string     `gorm:"type:varchar(45);not null;index" binding:"required" json:"name"`
	Price    float64    `gorm:"type:decimal(15,2);not null:default:0.00" binding:"required" json:"price"`
	Duration uint16     `gorm:"type:int;not null" binding:"required" json:"duration"`
	IsTrial  uint8      `gorm:"type:tinyint;not null:default:0" binding:"required" json:"is_trial"`
	Business []Business `gorm:"foreignKey:SubscriptionUID;references:UID" json:"business"`
}

// Business Schema definition of business table
type Business struct {
	gorm.Model
	UID                string        `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name               string        `gorm:"type:varchar(60);not null;index" binding:"required" json:"name"`
	Province           string        `gorm:"type:varchar(30);not null" binding:"required" json:"province"`
	Country            string        `gorm:"type:varchar(30);not null" binding:"required" json:"country"`
	SubscriptionExpiry time.Time     `gorm:"type:datetime;not null" binding:"required" json:"subscription_expiry"`
	Status             string        `gorm:"type:enum('active','inactive');not null" binding:"required" json:"status"`
	SubscriptionUID    string        `gorm:"not null" binding:"required" json:"subscription_uid"`
	Products           []Product     `gorm:"foreignKey:BusinessUID;references:UID" json:"products"`
	Collections        []Collection  `gorm:"foreignKey:BusinessUID;references:UID" json:"collections"`
	Vendors            []Vendor      `gorm:"foreignKey:BusinessUID;references:UID" json:"vendors"`
	Transactions       []Transaction `gorm:"foreignKey:BusinessUID;references:UID" json:"transactions"`
	Personnels         []Personnel   `gorm:"foreignKey:BusinessUID;references:UID" json:"personnel"`
}

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

// Transaction Schema definition of transaction table
type Transaction struct {
	gorm.Model
	UID         string  `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Category    string  `gorm:"type:enum('DEBIT','CREDIT');not null" binding:"required" json:"category"`
	Reference   string  `gorm:"type:varchar(45);not null;index" binding:"required" json:"reference"`
	Narration   string  `gorm:"type:varchar(75);not null" binding:"required" json:"narration"`
	Amount      float64 `gorm:"type:decimal(15,2);not null" binding:"required" json:"amount"`
	BusinessUID string  `gorm:"not null" binding:"required" json:"business_uid"`
}

// Vendor Schema definition of vendor table
type Vendor struct {
	gorm.Model
	UID         string `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name        string `gorm:"type:varchar(50);not null;index" binding:"required" json:"name"`
	Address     string `gorm:"type:varchar(65);not null" binding:"required" json:"address"`
	Province    string `gorm:"type:varchar(20);not null" binding:"required" json:"province"`
	Country     string `gorm:"type:varchar(45);not null" binding:"required" json:"country"`
	Email       string `gorm:"type:varchar(55);not null;index" binding:"required" json:"email"`
	PhoneNumber string `gorm:"type:varchar(20);not null;index" binding:"required" json:"phone_number"`
	Status      string `gorm:"type:enum('active','inactive');not null" binding:"required" json:"status"`
	BusinessUID string `gorm:"not null" binding:"required" json:"business_uid"`
}

// Product Schema definition of product table
type Product struct {
	gorm.Model
	UID            string  `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Brand          string  `gorm:"type:varchar(45);not null;index" binding:"required" json:"brand"`
	BarCode        string  `gorm:"type:varchar(25);not null;index" binding:"required" json:"barcode"`
	Name           string  `gorm:"type:varchar(45);not null;index" binding:"required" json:"name"`
	Price          float64 `gorm:"type:decimal(15,2);not null" binding:"required" json:"price"`
	TotalStock     uint32  `gorm:"type:int;not null" binding:"required" json:"total_stock"`
	AvailableStock uint32  `gorm:"type:int;not null" binding:"required" json:"available_stock"`
	LockedStock    uint32  `gorm:"type:int;default:0" binding:"required" json:"locked_stock"`
	Status         string  `gorm:"type:enum('available','unavailable','exhausted');not null" binding:"required" json:"status"`
	CategoryUID    string  `gorm:"not null" binding:"required" json:"category_uid"`
	BusinessUID    string  `gorm:"not null" binding:"required" json:"business_uid"`
}

// Collection Schema definition of product collection table
type Collection struct {
	gorm.Model
	UID            string    `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Brand          string    `gorm:"type:varchar(45);not null;index" binding:"required" json:"brand"`
	BarCode        string    `gorm:"type:varchar(25);not null;index" binding:"required" json:"barcode"`
	Name           string    `gorm:"type:varchar(45);not null;index" binding:"required" json:"name"`
	Price          float64   `gorm:"type:decimal(15,2);not null" binding:"required" json:"price"`
	ContentCount   uint32    `gorm:"type:int;not null" binding:"required" json:"content_count"`
	ExpiryDate     time.Time `gorm:"type:datetime;not null" binding:"required" json:"expiry_date"`
	TotalStock     uint32    `gorm:"type:int;not null" binding:"required" json:"total_stock"`
	AvailableStock uint32    `gorm:"type:int;not null" binding:"required" json:"available_stock"`
	LockedStock    uint32    `gorm:"type:int;default:0" binding:"required" json:"locked_stock"`
	Status         string    `gorm:"type:enum('available','unavailable','exhausted');not null;default:'available'" binding:"required" json:"status"`
	CategoryUID    string    `gorm:"not null" binding:"required" json:"category_uid"`
	BusinessUID    string    `gorm:"not null" binding:"required" json:"business_uid"`
}

// Category Schema definition of product categories table
type Category struct {
	gorm.Model
	UID        string       `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	Name       string       `gorm:"type:varchar(45);not null;index" binding:"required" json:"name"`
	Product    []Product    `gorm:"foreignKey:CategoryUID;references:UID" json:"products"`
	Collection []Collection `gorm:"foreignKey:CategoryUID;references:UID" json:"collections"`
}

// Customer Schema definition of customer table
type Customer struct {
	gorm.Model
	UID         string  `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	FirstName   string  `gorm:"type:varchar(20);not null;index" binding:"required" json:"first_name"`
	LastName    string  `gorm:"type:varchar(20);not null;index" binding:"required" json:"last_name"`
	Email       string  `gorm:"type:varchar(65);index" binding:"required" json:"email"`
	PhoneNumber string  `gorm:"type:varchar(20);not null;index" binding:"required" json:"phone_number"`
	Order       []Order `gorm:"foreignKey:CustomerUID;references:UID" json:"orders"`
}

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

// OrderProduct Schema definition of order_products jointable
type OrderProduct struct {
	gorm.Model
	OrderUID   string `gorm:"not null" binding:"required" json:"order_uid"`
	ProductUID string `gorm:"not null" binding:"required" json:"product_uid"`
	Quantity   uint16 `gorm:"type:int;not null" binding:"required" json:"quantity"`
}

// OrderCollection Schema definition of order_collections jointable
type OrderCollection struct {
	gorm.Model
	OrderUID      string `gorm:"not null" binding:"required" json:"order_uid"`
	CollectionUID string `gorm:"not null" binding:"required" json:"collection_uid"`
	Quantity      uint16 `gorm:"type:int;not null" binding:"required" json:"quantity"`
}

// Payment Schema definition of payment table
type Payment struct {
	gorm.Model
	UID           string `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	InvoiceNumber string `gorm:"type:varchar(20);not null;index" binding:"required" json:"invoice_number"`
	Medium        string `gorm:"type:enum('cash','debit-card','bank-transfer','other');not null" binding:"required" json:"medium"`
	Status        string `gorm:"type:enum('pending','completed');not null" binding:"required" json:"status"`
	OrderUID      string `gorm:"not null" binding:"required" json:"order_uid"`
}

// Parameter Schema definition of parameter table
type Parameter struct {
	gorm.Model
	UID         string      `gorm:"type:varchar(45);not null;primaryKey;uniqueIndex" binding:"required" json:"uid"`
	DisplayName string      `gorm:"type:varchar(65);not null" binding:"required" json:"display_name"`
	Name        string      `gorm:"type:varchar(45);not null;index" binding:"required" json:"name"`
	Value       interface{} `gorm:"type:json;not null" binding:"required" json:"value"`
}
