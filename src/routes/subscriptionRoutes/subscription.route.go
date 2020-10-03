package subscriptionroutes

import (
	"github.com/labstack/echo"

	subscriptioncontroller "store-app/src/controllers/subscriptionController"
)

// Index configures all subscription routes
func Index(root string, router *echo.Group) {
	// Create new subscription
	router.POST(root+"/create", subscriptioncontroller.CreateSubscription)
}
