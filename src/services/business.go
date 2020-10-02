package services

import (
	"errors"
	database "store-app/src/data"
)

// CreateNewBusiness Add new business account
func (_payload BusinessPayloadType) CreateNewBusiness() (interface{}, error) {
	// Get subscription
	record := new(SubscriptionPayloadType)
	record.UID = _payload.Subscription

	result, err := record.GetSingleSubscription()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("Subscription record does not exist")
	}

	data := &database.Business{
		Name:            _payload.Name,
		Province:        _payload.Province,
		Country:         _payload.Country,
		Status:          _payload.Status,
		SubscriptionUID: _payload.Subscription,
	}

	Construct := database.RepositoryConstruct{
		Payload: data,
	}

	response, err := Construct.AddOne()
	if err != nil {
		return nil, err
	}

	return response, nil
}
