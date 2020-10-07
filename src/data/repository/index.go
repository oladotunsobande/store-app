package repository

import (
	database "store-app/src/data"

	"gorm.io/gorm"
)

// DataConstruct Struct definition for data input
type DataConstruct struct {
	Model      interface{}
	Payload    interface{}
	Clause     string
	Parameters interface{}
}

var connection *gorm.DB

// Connect to SQL database
func init() {
	connection = database.ConnectSQL()
}
