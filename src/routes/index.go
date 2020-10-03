package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	subscriptionroutes "store-app/src/routes/subscriptionRoutes"
)

// APIIndex sets default response sent when the base path is called
func APIIndex(ctx echo.Context) error {
	return ctx.String(200, "Welcome to Store Manager API V1")
}

// RegisterRoutes register all app routes
func RegisterRoutes() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())

	rootRouter := router.Group("/v1")

	// account routes group
	subscriptionroutes.Index("/subscription", rootRouter)

	return router
}
