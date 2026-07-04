package domain

// Database Design
type Hotel struct {
	ID                   string `gorm:"column:id;primaryKey;Index"`
	Name                 string `gorm:"column:name;type:VARCHAR(255);NOT NULL"`
	Code                 string `gorm:"column:code;type:VARCHAR(255);NOT NULL"`
	Timezone             string `gorm:"column:timezone;type:VARCHAR(255);NOT NULL;Index"`
	ReceptionistWhatsapp string `gorm:"column:receptionist_whatsapp;type:VARCHAR(255)"`
	LogoURL              string `gorm:"column:logo_url;type:VARCHAR(255)"`
	WifiSSID             string `gorm:"column:wifi_ssid;type:VARCHAR(255)"`
	WifiUsername         string `gorm:"column:wifi_username;type:VARCHAR(255);"`
	WifiPassword         string `gorm:"column:wifi_password;type:VARCHAR(255)"`
	BreakfastTime        string `gorm:"column:breakfast_time;type:VARCHAR(255)"`
	BreakfastLocation    string `gorm:"column:breakfast_location;type:VARCHAR(255)"`
	CheckoutTime         string `gorm:"column:checkout_time;type:VARCHAR(255)"`
	LateCheckoutPolicy   string `gorm:"column:late_checkout_policy;type:VARCHAR(255)"`
	CreatedAt            int64  `gorm:"column:created_at;type:BIGINT"`
	UpdatedAt            int64  `gorm:"column:updated_at;type:BIGINT"`
	IsDeleted            bool   `gorm:"column:is_deleted;type:BIT"`
}

type HotelQueryModel struct {
	ID     string
	Name   string
	Code   string
	Status string
}
