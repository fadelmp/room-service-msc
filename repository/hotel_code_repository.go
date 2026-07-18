package repository

import (
	"errors"
	"room-service-msc/domain"
	"room-service-msc/infrastructure/logging"
	"time"

	"gorm.io/gorm"
)

// Interface
type HotelCodeRepository interface {
	FindAll(*gorm.DB, *HotelCodeQuery) ([]*domain.HotelCode, error)
	Create(*gorm.DB, *domain.Hotel) error
	Count(*gorm.DB, *HotelCodeQuery) (int64, error)
}

// Class
type hotelCodeRepository struct{}

type HotelCodeQuery struct {
	Code string
}

// Constructor
func NewHotelCodeRepository() HotelCodeRepository {
	return &hotelCodeRepository{}
}

// Implementation
func (r *hotelCodeRepository) FindAll(db *gorm.DB, query *HotelCodeQuery) ([]*domain.HotelCode, error) {

	// Get from DB
	var datas []*domain.HotelCode
	if err := db.Scopes(r.queryChain(query)).Find(&datas).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // atau sentinel error domain.ErrHotelNotFound
		}
		logging.Failed(err)
		return nil, err
	}

	logging.SuccessList(len(datas))
	return datas, nil
}

func (r *hotelCodeRepository) Count(db *gorm.DB, query *HotelCodeQuery) (int64, error) {

	var count int64
	if err := db.Model(&domain.Hotel{}).Scopes(r.queryChain(query)).Count(&count).Error; err != nil {
		logging.Failed(err)
		return 0, err
	}

	logging.Success(string(count))
	return count, nil
}

func (r *hotelCodeRepository) Create(db *gorm.DB, data *domain.Hotel) error {

	data.CreatedAt = time.Now().Unix()
	data.UpdatedAt = time.Now().Unix()

	if err := db.Create(data).Error; err != nil {
		logging.Failed(err)
		return err
	}

	logging.Success(data.ID)
	return nil
}

func (r *hotelCodeRepository) queryChain(query *HotelCodeQuery) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		db = WhereEqual(db, "code", query.Code)

		return db
	}
}
