package routes

import (
	"room-service-msc/handler"

	"github.com/labstack/echo/v4"
)

func WebFoodMenuRoute(routes *echo.Group, handler handler.FoodMenuHandler) {

	route := routes.Group("/food-menu")
	{
		route.POST("/create", handler.Create)
	}
}
