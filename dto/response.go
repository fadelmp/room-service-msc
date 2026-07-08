package dto

// @Summary Response object
// @Description Represents a Response Payload Data
type Response struct {
	Data    interface{} `json:"data"`
	Result  bool        `json:"result"`
	Message string      `json:"message"`
}
