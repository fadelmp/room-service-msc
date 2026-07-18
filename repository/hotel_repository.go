package repository

import (
	"errors"
	"room-service-msc/domain"
	"room-service-msc/infrastructure/logging"
	"time"

	"gorm.io/gorm"
)

// Interface
type HotelRepository interface {
	FindAll(*gorm.DB, *HotelQuery) ([]*domain.Hotel, error)
	FindOne(*gorm.DB, *HotelQuery) (*domain.Hotel, error)

	Create(*gorm.DB, *domain.Hotel) error
	Update(*gorm.DB, *domain.Hotel) error
	Delete(*gorm.DB, string) error
}

// Class
type hotelRepository struct{}

type HotelQuery struct {
	ID       string
	Name     string
	Code     string
	CodeLike string
	Status   string
}

// Constructor
func NewHotelRepository() HotelRepository {
	return &hotelRepository{}
}

// Implementation
func (r *hotelRepository) FindAll(db *gorm.DB, query *HotelQuery) ([]*domain.Hotel, error) {

	var datas []*domain.Hotel
	if err := db.Scopes(r.queryChain(query)).Find(&datas).Error; err != nil {
		logging.Failed(err)
		return nil, err
	}

	logging.SuccessList(len(datas))
	return datas, nil
}

func (r *hotelRepository) FindOne(db *gorm.DB, query *HotelQuery) (*domain.Hotel, error) {

	// Get from DB
	var data domain.Hotel
	if err := db.Scopes(r.queryChain(query)).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // atau sentinel error domain.ErrHotelNotFound
		}
		logging.Failed(err)
		return nil, err
	}

	logging.Success(data.ID)
	return &data, nil
}

func (r *hotelRepository) Create(db *gorm.DB, data *domain.Hotel) error {

	data.CreatedAt = time.Now().Unix()
	data.UpdatedAt = time.Now().Unix()

	if err := db.Create(data).Error; err != nil {
		logging.Failed(err)
		return err
	}

	logging.Success(data.ID)
	return nil
}

func (r *hotelRepository) Update(db *gorm.DB, data *domain.Hotel) error {

	data.UpdatedAt = time.Now().Unix()

	if err := db.Updates(data).Error; err != nil {
		logging.Failed(err)
		return err
	}

	logging.Success(data.ID)
	return nil
}

func (r *hotelRepository) Delete(db *gorm.DB, id string) error {

	// Manual soft delete
	if err := db.Where("id = ?", id).Delete(&domain.Hotel{}).Error; err != nil {
		logging.Failed(err)
		return err
	}

	logging.Success(id)
	return nil
}

func (r *hotelRepository) queryChain(query *HotelQuery) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		db = WhereEqual(db, "id", query.ID)
		db = WhereEqual(db, "name", query.Name)
		db = WhereEqual(db, "code", query.Code)
		db = WhereLike(db, "code", query.CodeLike)

		return db
	}
}
