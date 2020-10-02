package main

import (
	"log"
	"store-app/src/services"
)

func main() {
	subscription := services.SubscriptionPayloadType{
		UID: "d6f3d9e7-cdf3-40d4-acf8-ada631a1a402",
	}

	result, err := subscription.GetSingleSubscription()
	if err != nil {
		log.Fatal(err)
	}
	/*subscription := services.SubscriptionPayloadType{
		Name:     "Basic",
		Price:    2000.00,
		Duration: 28,
		IsTrial:  false,
	}

	result, err := subscription.CreateNewSubscription()
	if err != nil {
		log.Fatal(err)
	}*/

	log.Println("Subscription: ", result)
}
