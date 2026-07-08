package handler

import (
	"net/http"
	"room-service-msc/dto"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// Parse ID param
func GetIDParam(c echo.Context) (int64, error) {
	return strconv.ParseInt(c.Param("id"), 10, 64)
}

// Bind + validate DTO
func BindAndValidate(c echo.Context, v interface{}, validate *validator.Validate) error {
	if err := c.Bind(v); err != nil {
		return err
	}
	return validate.Struct(v)
}

func Success(c echo.Context, data interface{}) error {

	httpCode := http.StatusOK
	resp := &dto.Response{Result: true, Data: data}

	return createResponse(c, &httpCode, resp)
}

func Failed(c echo.Context, message string) error {

	httpCode := http.StatusInternalServerError
	resp := &dto.Response{Result: false, Message: message}

	return createResponse(c, &httpCode, *resp)
}

func BadRequest(c echo.Context, message string) error {

	httpCode := http.StatusBadRequest
	resp := &dto.Response{Result: false, Message: message}

	return createResponse(c, &httpCode, *resp)
}

func NotFound(c echo.Context) error {

	httpCode := http.StatusNotFound
	resp := dto.Response{Result: false}

	return createResponse(c, &httpCode, resp)
}

func createResponse(c echo.Context, httpCode *int, response interface{}) error {

	c.Response().WriteHeader(*httpCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	return c.JSONPretty(*httpCode, response, "  ")
}
