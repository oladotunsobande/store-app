package database

import (
	"log"

	"gorm.io/gorm"
)

var connection *gorm.DB

// RepositoryConstruct Struct definition for data input
type RepositoryConstruct struct {
	Schema     interface{}
	Payload    interface{}
	Clause     string
	Parameters interface{}
}

// Connect to SQL database
func init() {
	connection = ConnectSQL()
}

// GetOne Repository to retrieve a single row
func (_repositoryConstruct RepositoryConstruct) GetOne() (interface{}, error) {
	result := connection.Where(_repositoryConstruct.Clause, _repositoryConstruct.Parameters).First(_repositoryConstruct.Schema)

	if result.Error != nil {
		return nil, result.Error
	}

	log.Println("GetOne - Result: ", result)

	return result, nil
}

// AddOne Repository function to add a single row to table
func (_repositoryConstruct RepositoryConstruct) AddOne() (interface{}, error) {
	result := connection.Create(_repositoryConstruct.Payload)

	if result.Error != nil {
		return nil, result.Error
	}

	log.Println("AddOne - Result: ", result)

	return result, nil
}

//func AddMany() {}

//func UpdateOne() {}

//func UpdateMany() {}
