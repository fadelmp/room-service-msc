package domain

import "gorm.io/gorm"

// Database Design
type HotelFoodMenu struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement"`
	Prefix    string         `gorm:"column:prefix;type:VARCHAR(3);NOT NULL"`
	Code      string         `gorm:"column:code;type:VARCHAR(10);NOT NULL"`
	Sequence  int            `gorm:"column:sequence;type:INT;NOT NULL"`
	CreatedBy string         `gorm:"column:created_by;type:VARCHAR(100)"`
	CreatedAt int64          `gorm:"column:created_at;type:BIGINT;autoCreateTime"`
	UpdatedBy string         `gorm:"column:created_by;type:VARCHAR(100)"`
	UpdatedAt int64          `gorm:"column:updated_at;type:BIGINT;autoUpdateTime"`
	DeletedBy string         `gorm:"column:created_by;type:VARCHAR(100)"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
