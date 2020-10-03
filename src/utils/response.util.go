package utils

// Result defines api response interface
type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Token   string      `json:"token,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Count   uint64      `json:"count,omitempty"`
}

// SuccessAuthResult returns successful authentication interface
func SuccessAuthResult(data interface{}, message string, token string) interface{} {
	return Result{
		Success: true,
		Message: message,
		Token:   token,
		Data:    data,
	}
}

// SuccessResult returns data for successful api call
func SuccessResult(data interface{}) interface{} {
	return Result{
		Success: true,
		Data:    data,
	}
}

// SuccessResultWithMessage returns data for successful api call
func SuccessResultWithMessage(data interface{}, message string) interface{} {
	return Result{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// SuccessResultWithCount returns list and count
func SuccessResultWithCount(data interface{}, count uint64) interface{} {
	return Result{
		Success: true,
		Data:    data,
		Count:   count,
	}
}

// ErrorMessage returns error from failed api call
func ErrorMessage(err string) interface{} {
	return Result{
		Success: false,
		Error:   err,
	}
}

// SuccessMessage returns message for successful api call
func SuccessMessage(message string) interface{} {
	return Result{
		Success: true,
		Message: message,
	}
}
