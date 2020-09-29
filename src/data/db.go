package database

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	secrets "store-app/src/config"
	"store-app/src/data/models"
)

// GetSchemas Return slice of all table schema interfaces
func GetSchemas() []interface{} {
	var schemas []interface{}

	schemas = append(schemas, models.SubscriptionSchema())
	schemas = append(schemas, models.BusinessSchema())
	schemas = append(schemas, models.CategorySchema())
	schemas = append(schemas, models.VendorSchema())
	schemas = append(schemas, models.ProductSchema())
	schemas = append(schemas, models.CollectionSchema())
	schemas = append(schemas, models.TransactionSchema())
	schemas = append(schemas, models.PersonnelSchema())
	schemas = append(schemas, models.CustomerSchema())
	schemas = append(schemas, models.OrderSchema())
	schemas = append(schemas, models.OrderProductSchema())
	schemas = append(schemas, models.OrderCollectionSchema())
	schemas = append(schemas, models.PaymentSchema())
	schemas = append(schemas, models.ParameterSchema())

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
