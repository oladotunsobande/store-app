package company

import (
	"store-app/src/data/models"
)

// BusinessPayloadType Struct definition for business payload
type BusinessPayloadType struct {
	Name         string `json:"name"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Status       string `json:"status"`
	Subscription string `json:"subscription"`
}

// CreateNewBusiness Add new business account
func (_payload BusinessPayloadType) CreateNewBusiness() (response, error) {

	data := models.Business{
		Name:            _payload.Name,
		Province:        _payload.Province,
		Country:         _payload.Country,
		Status:          _payload.Status,
		SubscriptionUID: _payload.Subscription,
	}

}
