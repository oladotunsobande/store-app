package services

import (
	"fmt"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

// BusinessRequest Struct definition for business row
type BusinessRequest struct {
	UID          string `json:"uid"`
	Name         string `json:"name"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Status       string `json:"status"`
	Subscription string `json:"subscription"`
}

// SubscriptionRequest Struct definition for subscription row
type SubscriptionRequest struct {
	UID      string  `json:"uid" validate:"required_without=Name Price Duration IsTrial"`
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"min=0.00,required"`
	Duration uint16  `json:"duration" validate:"min=1,required"`
	IsTrial  bool    `json:"is_trial" validate:"required"`
}

// PayloadValidation validates request payload
func PayloadValidation(_request interface{}) error {
	validate = validator.New()

	err := validate.Struct(_request)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("Validation error: ", err)
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}
	}

	return nil
}
