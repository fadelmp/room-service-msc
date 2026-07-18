package repository

import (
	"errors"
	"room-service-msc/domain"
	"room-service-msc/infrastructure/logging"
	"room-service-msc/infrastructure/message"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Interface
type FoodMenuRepository interface {
	FindAll(*gorm.DB, *FoodMenuQuery) ([]*domain.FoodMenu, error)
	FindOne(*gorm.DB, *FoodMenuQuery) (*domain.FoodMenu, error)

	Create(*gorm.DB, *domain.FoodMenu) error
	Update(*gorm.DB, *domain.FoodMenu) error
	Delete(*gorm.DB, string) error
}

// Class
type foodMenuRepository struct{}

type FoodMenuQuery struct {
	ID              string
	HotelID         string
	Name            string
	NameLike        string
	DescriptionLike string
	Category        string
	Categories      []string
	Availability    *bool
}

// Constructor
func NewFoodMenuRepository() FoodMenuRepository {
	return &foodMenuRepository{}
}

// Implementation
func (r *foodMenuRepository) FindAll(db *gorm.DB, query *FoodMenuQuery) ([]*domain.FoodMenu, error) {

	var datas []*domain.FoodMenu
	if err := db.Scopes(r.queryChain(query)).Find(&datas).Error; err != nil {
		logging.Failed(err)
		return nil, err
	}

	logging.SuccessList(len(datas))
	return datas, nil
}

func (r *foodMenuRepository) FindOne(db *gorm.DB, query *FoodMenuQuery) (*domain.FoodMenu, error) {

	// Get from DB
	var data domain.FoodMenu
	if err := db.Scopes(r.queryChain(query)).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(message.ErrFoodMenuNotFound)
		}
		logging.Failed(err, zap.String("id", data.ID))
		return nil, err
	}

	logging.Success(data.ID)
	return &data, nil
}

func (r *foodMenuRepository) Create(db *gorm.DB, data *domain.FoodMenu) error {

	if err := db.Create(data).Error; err != nil {
		logging.Failed(err, zap.String("id", data.ID))
		return err
	}

	logging.Success(data.ID)
	return nil
}

func (r *foodMenuRepository) Update(db *gorm.DB, data *domain.FoodMenu) error {

	if err := db.Model(&domain.FoodMenu{}).Where("id = ?", data.ID).Updates(data).Error; err != nil {
		logging.Failed(err, zap.String("id", data.ID))
		return err
	}

	logging.Success(data.ID)
	return nil
}

func (r *foodMenuRepository) Delete(db *gorm.DB, id string) error {

	// Soft delete
	if err := db.Where("id = ?", id).Delete(&domain.FoodMenu{}).Error; err != nil {
		logging.Failed(err, zap.String("id", id))
		return err
	}

	logging.Success(id)
	return nil
}

func (r *foodMenuRepository) queryChain(query *FoodMenuQuery) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		db = WhereEqual(db, "id", query.ID)
		db = WhereEqual(db, "hotel_id", query.HotelID)
		db = WhereEqual(db, "name", query.Name)
		db = WhereLike(db, "name", query.NameLike)
		db = WhereLike(db, "description", query.DescriptionLike)
		db = WhereEqual(db, "category", query.Category)
		db = WhereIn(db, "category", query.Categories)
		db = WherePtr(db, "category", query.Availability)

		return db
	}
}
