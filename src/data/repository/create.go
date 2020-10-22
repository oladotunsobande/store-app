package repository

import (
	"log"
)

// AddOne Repository function to add a single row to table
func (_dataConstruct DataConstruct) AddOne() error {
	result := connection.Create(_dataConstruct.Payload)

	if result.Error != nil {
		return result.Error
	}

	log.Println("Rows: ", result.RowsAffected)

	return nil
}
