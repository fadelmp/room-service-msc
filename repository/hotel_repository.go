package repository

import (
	"errors"
	"room-service-msc/domain"
	"room-service-msc/infrastructure/logging"
	"time"

	"gorm.io/gorm"
)

// Interface
type HotelRepositoryInterface interface {
	FindAll(*gorm.DB, *HotelQuery) ([]*domain.Hotel, error)
	FindOne(*gorm.DB, *HotelQuery) (*domain.Hotel, error)

	Create(*gorm.DB, *domain.Hotel) error
	Update(*gorm.DB, *domain.Hotel) error
	Delete(*gorm.DB, *string) error
}

// Class
type HotelRepository struct {
	log *logging.Context
}

type HotelQuery struct {
	ID       string
	Name     string
	Code     string
	CodeLike string
	Status   string
}

// Constructor
func NewHotelRepository(log *logging.Context) *HotelRepository {
	return &HotelRepository{
		log: log,
	}
}

// Implementation
func (r *HotelRepository) FindAll(db *gorm.DB, query *HotelQuery) ([]*domain.Hotel, error) {

	var datas []*domain.Hotel
	if err := db.Scopes(r.queryChain(query)).Find(&datas).Error; err != nil {
		logging.Failed(r.log, err)
		return nil, err
	}

	logging.SuccessList(r.log, len(datas))
	return datas, nil
}

func (r *HotelRepository) FindOne(db *gorm.DB, query *HotelQuery) (*domain.Hotel, error) {

	// Get from DB
	var data domain.Hotel
	if err := db.Scopes(r.queryChain(query)).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // atau sentinel error domain.ErrHotelNotFound
		}
		logging.Failed(r.log, err)
		return nil, err
	}

	logging.Success(r.log, data.ID)
	return &data, nil
}

func (r *HotelRepository) Create(db *gorm.DB, data *domain.Hotel) error {

	data.CreatedAt = time.Now().Unix()
	data.UpdatedAt = time.Now().Unix()

	if err := db.Create(data).Error; err != nil {
		logging.Failed(r.log, err)
		return err
	}

	logging.Success(r.log, data.ID)
	return nil
}

func (r *HotelRepository) Update(db *gorm.DB, data *domain.Hotel) error {

	data.UpdatedAt = time.Now().Unix()

	if err := db.Updates(data).Error; err != nil {
		logging.Failed(r.log, err)
		return err
	}

	logging.Success(r.log, data.ID)
	return nil
}

func (r *HotelRepository) Delete(db *gorm.DB, id *string) error {

	// Manual soft delete
	if err := db.Where("id = ?", id).Delete(&domain.Hotel{}).Error; err != nil {
		logging.Failed(r.log, err)
		return err
	}

	logging.Success(r.log, *id)
	return nil
}

func (r *HotelRepository) queryChain(query *HotelQuery) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		db = r.idEqual(db, query.ID)
		db = r.nameEqual(db, query.Name)
		db = r.codeEqual(db, query.Code)
		db = r.codeLike(db, query.CodeLike)

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

func (r *HotelRepository) codeLike(db *gorm.DB, code string) *gorm.DB {

	if code == "" {
		return db
	}

	return db.Where("code LIKE %" + code + "%")
}
