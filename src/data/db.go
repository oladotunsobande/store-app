package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	secrets "store-app/src/config"
)

// Connect Connect to SQL database
func Connect() (db *sql.DB) {
	db, err := sql.Open(secrets.GetSecrets().DatabaseProvider, secrets.GetSecrets().DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Set connection pool options
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Println("Failed to establish database connection")
		panic(err.Error())
	} else {
		log.Println("Database connected!")
	}

	return db
}

// ConnectSQL Connect to MySQL database
func ConnectSQL() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       secrets.GetSecrets().DatabaseURL,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	// Set connection pool options
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		log.Println("Failed to establish database connection")
		panic(err.Error())
	} else {
		log.Println("[MySQL] - Database connected!")
	}

	return db
}
