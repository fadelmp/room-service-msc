package domain

import "gorm.io/gorm"

// Database Design
type FoodMenu struct {
	ID          string         `gorm:"column:id;primaryKey"`
	HotelID     string         `gorm:"column:hotel_id;type:VARCHAR(50);NOT NULL"`
	Name        string         `gorm:"column:name;type:VARCHAR(100);NOT NULL"`
	Description string         `gorm:"column:description;type:VARCHAR(255)"`
	Price       int            `gorm:"column:price;type:INT"`
	Category    string         `gorm:"column:category;type:VARCHAR(20)"`
	IsAvailable bool           `gorm:"column:is_available;type:BIT"`
	CreatedAt   int64          `gorm:"column:created_at;type:BIGINT;autoCreateTime"`
	UpdatedAt   int64          `gorm:"column:updated_at;type:BIGINT;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`

	Hotel Hotel `gorm:"foreignKey:HotelID"` // belongs-to
}
