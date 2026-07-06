package dto

type HotelDto struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Code                 string `json:"code"`
	Timezone             string `json:"timezone"`
	ReceptionistWhatsapp string `json:"receptionist_whatsapp"`
	LogoURL              string `json:"logo_url"`
	WifiSSID             string `json:"wifi_ssid"`
	WifiUsername         string `json:"wifi_username"`
	WifiPassword         string `json:"wifi_password"`
	BreakfastTime        string `json:"breakfast_time"`
	BreakfastLocation    string `json:"breakfast_location"`
	CheckoutTime         string `json:"checkout_time"`
	LateCheckoutPolicy   string `json:"late_checkout_policy"`
	CreatedAt            int64  `json:"created_at"`
	UpdatedAt            int64  `json:"updated_at"`
}

type CreateHotelDto struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
