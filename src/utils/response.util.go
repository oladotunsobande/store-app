package utils

import (
	"encoding/json"
	"log"
)

// result defines api response interface
type result struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Token   string      `json:"token,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Count   uint64      `json:"count,omitempty"`
}

// SetResponse converts result struct to map (JSON equivalent)
func (_result result) SetResponse() map[string]interface{} {
	jsonObject, _ := json.Marshal(_result)

	vals := make(map[string]interface{})
	err := json.Unmarshal(jsonObject, &vals)
	if err != nil {
		log.Fatal(err)
	}

	return vals
}

// SuccessAuthResult returns successful authentication interface
func SuccessAuthResult(data interface{}, message string, token string) map[string]interface{} {
	var _resultPointer = result{
		Success: true,
		Message: message,
		Token:   token,
		Data:    data,
	}

	return _resultPointer.SetResponse()
}

// SuccessResult returns data for successful api call
func SuccessResult(data interface{}) map[string]interface{} {
	var _resultPointer = result{
		Success: true,
		Data:    data,
	}

	return _resultPointer.SetResponse()
}

// SuccessResultWithMessage returns data for successful api call
func SuccessResultWithMessage(data interface{}, message string) map[string]interface{} {
	var _resultPointer = result{
		Success: true,
		Message: message,
		Data:    data,
	}

	return _resultPointer.SetResponse()
}

// SuccessResultWithCount returns list and count
func SuccessResultWithCount(data interface{}, count uint64) map[string]interface{} {
	var _resultPointer = result{
		Success: true,
		Data:    data,
		Count:   count,
	}

	return _resultPointer.SetResponse()
}

// ErrorMessage returns error from failed api call
func ErrorMessage(err string) map[string]interface{} {
	var _resultPointer = result{
		Success: false,
		Error:   err,
	}

	return _resultPointer.SetResponse()
}

// SuccessMessage returns message for successful api call
func SuccessMessage(message string) map[string]interface{} {
	var _resultPointer = result{
		Success: true,
		Message: message,
	}

	return _resultPointer.SetResponse()
}
