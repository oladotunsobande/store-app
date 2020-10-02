package database

import (
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

var connection *gorm.DB

// RepositoryConstruct Struct definition for data input
type RepositoryConstruct struct {
	Model      interface{}
	Payload    interface{}
	Clause     string
	Parameters interface{}
}

// Connect to SQL database
func init() {
	connection = ConnectSQL()
}

// StructToJSON converts struct to JSON
func StructToJSON(data interface{}) ([]byte, error) {
	jsonObject, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonObject, nil
}

// GetOne Repository to retrieve a single row
func (_repositoryConstruct RepositoryConstruct) GetOne() (interface{}, error) {
	record := map[string]interface{}{}

	result := connection.Model(_repositoryConstruct.Model).Where(_repositoryConstruct.Clause, _repositoryConstruct.Parameters).First(record)
	if result.Error != nil {
		return nil, result.Error
	}

	response, err := StructToJSON(record)
	if err != nil {
		return nil, err
	}

	log.Println("GetOne - Result: ", string(response))

	return response, nil
}

// AddOne Repository function to add a single row to table
func (_repositoryConstruct RepositoryConstruct) AddOne() (interface{}, error) {
	result := connection.Create(_repositoryConstruct.Payload)

	if result.Error != nil {
		return nil, result.Error
	}

	log.Println("Rows: ", result.RowsAffected)

	return result, nil
}

//func AddMany() {}

//func UpdateOne() {}

//func UpdateMany() {}
