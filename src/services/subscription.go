package services

import (
	database "store-app/src/data"
	repository "store-app/src/data/repository"

	"github.com/google/uuid"
)

// CreateNewSubscription Add new subscription
func (_payload SubscriptionPayloadType) CreateNewSubscription() (interface{}, error) {
	var isTrial uint8 = 0

	if _payload.IsTrial == true {
		isTrial = 1
	}

	newSubscription := &database.Subscription{
		UID:      uuid.New().String(),
		Name:     _payload.Name,
		Price:    _payload.Price,
		Duration: _payload.Duration,
		IsTrial:  isTrial,
	}

	Construct := repository.DataConstruct{
		Payload: newSubscription,
	}

	_, err := Construct.AddOne()
	if err != nil {
		return nil, err
	}

	// Get new record
	record := SubscriptionPayloadType{
		UID: newSubscription.UID,
	}

	response, err := record.GetSingleSubscription()
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetSingleSubscription Get single subscription record
func (_payload SubscriptionPayloadType) GetSingleSubscription() (interface{}, error) {
	Construct := repository.DataConstruct{
		Model:      &database.Subscription{},
		Clause:     "uid = ?",
		Parameters: _payload.UID,
	}

	response, err := Construct.GetOne()
	if err != nil {
		return nil, err
	}

	return response, nil
}
