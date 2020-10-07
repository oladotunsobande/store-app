package repository

import (
	"log"
)

// AddOne Repository function to add a single row to table
func (_dataConstruct DataConstruct) AddOne() (bool, error) {
	result := connection.Create(_dataConstruct.Payload)

	if result.Error != nil {
		return false, result.Error
	}

	log.Println("Rows: ", result.RowsAffected)

	return true, nil
}
