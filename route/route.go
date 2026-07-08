package routes

import (
	"room-service-msc/injection"

	"github.com/labstack/echo/v4"
)

func Init(routes *echo.Echo) *echo.Echo {

	// Swagger Documentation Route
	//SwaggerRoute(routes)

	// Injection
	foodMenu := injection.FoodMenuInjection()

	web := routes.Group("/web")
	{
		WebFoodMenuRoute(web, *foodMenu)
	}

	// mobile := routes.Group("/mobile")
	// {

	// }

	return routes
}
