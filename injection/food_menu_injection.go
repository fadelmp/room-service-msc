package injection

import (
	"room-service-msc/handler"
	"room-service-msc/mapper"
	"room-service-msc/repository"
	"room-service-msc/service"
	"room-service-msc/usecase"
)

func FoodMenuInjection() *handler.FoodMenuHandler {

	// Repository Layer
	foodMenuRepo := repository.NewFoodMenuRepository()
	hotelRepo := repository.NewHotelRepository()

	// Service Layer
	hotelService := service.NewHotelService(hotelRepo)

	// Mapper Layer
	mapper := mapper.NewFoodMenuMapper()

	// Usecase Layer
	foodMenuUsecase := usecase.NewFoodMenuUsecase(mapper, hotelService, foodMenuRepo)

	// Handler Layer
	handler := handler.NewFoodMenuHandler(foodMenuUsecase)

	return &handler
}
