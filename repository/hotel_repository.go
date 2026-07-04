package repository

import (
	"errors"
	"room-service-msc/domain"
	"room-service-msc/infrastructure/logging"

	"gorm.io/gorm"
)

// Interface
type HotelRepositoryInterface interface {
	FindAll(*domain.HotelQueryModel) ([]*domain.Hotel, error)
	FindOne(*domain.HotelQueryModel) (*domain.Hotel, error)

	Create(*gorm.DB, *domain.Hotel) error
	Update(*gorm.DB, *domain.Hotel) error
	Delete(*gorm.DB, *string) error
}

// Class
type HotelRepository struct {
	db  *gorm.DB
	log *logging.Context
}

// Constructor
func NewHotelRepository(db *gorm.DB, log *logging.Context) *HotelRepository {
	return &HotelRepository{
		db:  db,
		log: log,
	}
}

// Implementation
func (r *HotelRepository) FindAll(query *domain.HotelQueryModel) ([]*domain.Hotel, error) {

	var datas []*domain.Hotel
	if err := r.db.Scopes(r.queryChain(query)).Find(&datas).Error; err != nil {
		logging.Failed(r.log, err)
		return nil, err
	}

	logging.SuccessList(r.log, len(datas))
	return datas, nil
}

func (r *HotelRepository) FindOne(query *domain.HotelQueryModel) (*domain.Hotel, error) {

	// Get from DB
	var data domain.Hotel
	if err := r.db.Scopes(r.queryChain(query)).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // atau sentinel error domain.ErrHotelNotFound
		}
		logging.Failed(r.log, err)
		return nil, err
	}

	logging.Success(r.log, data.ID)
	return &data, nil
}

func (r *HotelRepository) Create(tx *gorm.DB, data *domain.Hotel) error {

	if err := tx.Create(data).Error; err != nil {
		logging.Failed(r.log, err)
		return err
	}

	logging.Success(r.log, data.ID)
	return nil
}

func (r *HotelRepository) Update(tx *gorm.DB, data *domain.Hotel) error {

	if err := tx.Updates(data).Error; err != nil {
		logging.Failed(r.log, err)
		return err
	}

	logging.Success(r.log, data.ID)
	return nil
}

func (r *HotelRepository) Delete(tx *gorm.DB, id *string) error {

	// Manual soft delete
	if err := tx.Model(&domain.Hotel{}).Delete("id = ?", id).Error; err != nil {
		logging.Failed(r.log, err)
		return err
	}

	logging.Success(r.log, *id)
	return nil
}

func (r *HotelRepository) queryChain(query *domain.HotelQueryModel) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		db = r.idEqual(db, query.ID)
		db = r.nameEqual(db, query.Name)
		db = r.codeEqual(db, query.Code)

		return db
	}
}

func (r *HotelRepository) idEqual(db *gorm.DB, id string) *gorm.DB {

	if id != "" {
		return db.Where("id = ?", id)
	}
	return db
}

func (r *HotelRepository) nameEqual(db *gorm.DB, name string) *gorm.DB {

	if name != "" {
		return db.Where("name = ?", name)
	}
	return db
}

func (r *HotelRepository) codeEqual(db *gorm.DB, code string) *gorm.DB {

	if code != "" {
		return db.Where("code = ?", code)
	}
	return db
}
