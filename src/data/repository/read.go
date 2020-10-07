package repository

// GetOne Repository to retrieve a single row
func (_dataConstruct DataConstruct) GetOne() (map[string]interface{}, error) {
	record := map[string]interface{}{}

	result := connection.Model(_dataConstruct.Model).Where(_dataConstruct.Clause, _dataConstruct.Parameters).First(record)
	if result.Error != nil {
		return nil, result.Error
	}

	return record, nil
}
