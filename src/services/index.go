package services

// BusinessPayloadType Struct definition for business row
type BusinessPayloadType struct {
	UID          string `json:"uid"`
	Name         string `json:"name"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Status       string `json:"status"`
	Subscription string `json:"subscription"`
}

// SubscriptionPayloadType Struct definition for subscription row
type SubscriptionPayloadType struct {
	UID      string  `json:"uid"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Duration uint16  `json:"duration"`
	IsTrial  bool    `json:"is_trial"`
}
