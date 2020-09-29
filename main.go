package main

import (
	"log"
	"store-app/src/services"
)

func main() {
	subscription := services.SubscriptionPayloadType{
		Name:     "Premium",
		Price:    20000.00,
		Duration: 365,
		IsTrial:  false,
	}

	result, err := subscription.CreateNewSubscription()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("New Subscription: ", result)
}
