package database

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	secrets "store-app/src/config"
)

// GetSchemas Return slice of all table schema interfaces
func GetSchemas() []interface{} {
	var schemas []interface{}

	schemas = append(schemas, &Subscription{})
	schemas = append(schemas, &Business{})
	schemas = append(schemas, &Category{})
	schemas = append(schemas, &Vendor{})
	schemas = append(schemas, &Product{})
	schemas = append(schemas, &Collection{})
	schemas = append(schemas, &Transaction{})
	schemas = append(schemas, &Personnel{})
	schemas = append(schemas, &Customer{})
	schemas = append(schemas, &Order{})
	schemas = append(schemas, &OrderProduct{})
	schemas = append(schemas, &OrderCollection{})
	schemas = append(schemas, &Payment{})
	schemas = append(schemas, &Parameter{})

	return schemas
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
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

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

		// Migrate all table schemas and relationships
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(GetSchemas()...)
	}

	return db
}
