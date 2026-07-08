package handler

import (
	"room-service-msc/dto"
	"room-service-msc/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// Interface
type FoodMenuHandler interface {
	Create(e echo.Context) error
}

// Class
type foodMenuHandler struct {
	usecase   usecase.FoodMenuUsecase
	validator *validator.Validate
}

// Constructor
func NewFoodMenuHandler(usecase usecase.FoodMenuUsecase) FoodMenuHandler {
	return &foodMenuHandler{
		usecase:   usecase,
		validator: validator.New(),
	}
}

func (h *foodMenuHandler) Create(c echo.Context) error {

	var req *dto.FoodMenuDto
	if err := BindAndValidate(c, req, h.validator); err != nil {
		return BadRequest(c, "invalid request body")
	}

	result, err := h.usecase.Create(req)
	if err != nil {
		return Failed(c, err.Error())
	}

	// Return success
	return Success(c, result)
}
