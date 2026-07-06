package domain

import "gorm.io/gorm"

// Database Design
type Room struct {
	ID         string         `gorm:"column:id;primaryKey;Index"`
	HotelID    string         `gorm:"column:hotel_id;type:VARCHAR(255);NOT NULL"`
	RoomNumber string         `gorm:"column:room_number;type:VARCHAR(255);NOT NULL"`
	Floor      string         `gorm:"column:floor;type:VARCHAR(10);NOT NULL;Index"`
	CreatedAt  int64          `gorm:"column:created_at;type:BIGINT;autoCreateTime"`
	UpdatedAt  int64          `gorm:"column:updated_at;type:BIGINT;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index"`

	Hotel Hotel `gorm:"foreignKey:HotelID"` // belongs-to: room ini milik satu hotel
}
