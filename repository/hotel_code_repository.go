package repository

import (
	"errors"
	"room-service-msc/domain"
	"room-service-msc/infrastructure/logging"
	"time"

	"gorm.io/gorm"
)

// Interface
type HotelCodeRepositoryInterface interface {
	FindAll(*gorm.DB, *HotelCodeQuery) ([]*domain.HotelCode, error)
	Create(*gorm.DB, *domain.HotelCode) error
	Count(*gorm.DB, *domain.HotelCode) (int64, error)
}

// Class
type HotelCodeRepository struct {
	log *logging.Context
}

type HotelCodeQuery struct {
	ID   string
	Code string
}

// Constructor
func NewHotelCodeRepository(log *logging.Context) *HotelCodeRepository {
	return &HotelCodeRepository{
		log: log,
	}
}

// Implementation
func (r *HotelCodeRepository) FindAll(db *gorm.DB, query *HotelCodeQuery) ([]*domain.Hotel, error) {

	// Get from DB
	var datas []*domain.Hotel
	if err := db.Scopes(r.queryChain(query)).Find(&datas).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // atau sentinel error domain.ErrHotelNotFound
		}
		logging.Failed(r.log, err)
		return nil, err
	}

	logging.SuccessList(r.log, len(datas))
	return datas, nil
}

func (r *HotelCodeRepository) Count(db *gorm.DB, query *HotelCodeQuery) (int64, error) {

	var count int64
	if err := db.Model(&domain.Hotel{}).Scopes(r.queryChain(query)).Count(&count).Error; err != nil {
		logging.Failed(r.log, err)
		return 0, err
	}

	logging.Success(r.log, string(count))
	return count, nil
}

func (r *HotelCodeRepository) Create(db *gorm.DB, data *domain.Hotel) error {

	data.CreatedAt = time.Now().Unix()
	data.UpdatedAt = time.Now().Unix()

	if err := db.Create(data).Error; err != nil {
		logging.Failed(r.log, err)
		return err
	}

	logging.Success(r.log, data.ID)
	return nil
}

func (r *HotelCodeRepository) queryChain(query *HotelCodeQuery) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		db = r.codeEqual(db, query.Code)

		return db
	}
}

func (r *HotelCodeRepository) codeEqual(db *gorm.DB, code string) *gorm.DB {

	if code == "" {
		return db
	}

	return db.Where("code = ?", code)
}
