package services

import (
	"errors"
	database "store-app/src/data"
	repository "store-app/src/data/repository"
)

// CreateNewBusiness Add new business account
func (_payload BusinessRequest) CreateNewBusiness() error {
	// Get subscription
	record := new(SubscriptionRequest)
	record.UID = _payload.Subscription

	result, err := record.GetSingleSubscription()
	if err != nil {
		return err
	}

	if result == nil {
		return errors.New("Subscription record does not exist")
	}

	data := &database.Business{
		Name:            _payload.Name,
		Province:        _payload.Province,
		Country:         _payload.Country,
		Status:          _payload.Status,
		SubscriptionUID: _payload.Subscription,
	}

	Construct := repository.DataConstruct{
		Payload: data,
	}

	err = Construct.AddOne()
	if err != nil {
		return err
	}

	return nil
}
