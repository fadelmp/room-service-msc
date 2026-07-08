package dto

type FoodMenuDto struct {
	ID          string `json:"id"`
	HotelID     string `json:"hotel_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	IsAvailable bool   `json:"is_available"`
}
