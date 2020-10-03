package subscriptioncontroller

import (
	"github.com/labstack/echo"

	"store-app/src/services"
	"store-app/src/utils"
)

// CreateSubscription processes request for the creation of a new subscription
func CreateSubscription(ctx echo.Context) (err error) {
	subscription := new(services.SubscriptionPayloadType)

	if err := ctx.Bind(subscription); err != nil {
		return err
	}

	// [TODO] Request payload validation

	result, err := subscription.CreateNewSubscription()
	if err != nil {
		return ctx.JSON(400, utils.ErrorMessage(err.Error()))
	}

	return ctx.JSON(200, utils.SuccessResultWithMessage(result, "Subscription created successfully"))
}
