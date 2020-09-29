package services

import (
	database "store-app/src/data"
	"store-app/src/data/models"

	"github.com/google/uuid"
)

// CreateNewSubscription Add new subscription
func (_payload SubscriptionPayloadType) CreateNewSubscription() (interface{}, error) {
	var isTrial uint8 = 0

	if _payload.IsTrial == true {
		isTrial = 1
	}

	newSubscription := &models.Subscription{
		UID:      uuid.New().String(),
		Name:     _payload.Name,
		Price:    _payload.Price,
		Duration: _payload.Duration,
		IsTrial:  isTrial,
	}

	Construct := database.RepositoryConstruct{
		Payload: newSubscription,
	}

	response, err := Construct.AddOne()
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetSingleSubscription Get single subscription record
func (_payload SubscriptionPayloadType) GetSingleSubscription() (interface{}, error) {
	var schema models.Subscription

	Construct := database.RepositoryConstruct{
		Schema:     schema,
		Clause:     "uid = ?",
		Parameters: _payload.UID,
	}

	response, err := Construct.GetOne()
	if err != nil {
		return nil, err
	}

	return response, nil
}