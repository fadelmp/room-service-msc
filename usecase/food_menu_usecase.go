package usecase

import (
	"errors"
	"room-service-msc/dto"
	"room-service-msc/infrastructure/db"
	"room-service-msc/infrastructure/message"
	"room-service-msc/mapper"
	"room-service-msc/repository"
	"room-service-msc/utils"

	"gorm.io/gorm"
)

type FoodMenuUsecase interface {
	Create(*dto.FoodMenuDto) (*dto.FoodMenuDto, error)
}

type foodMenuUsecase struct {
	mapper       mapper.FoodMenuMapper
	hotelUsecase HotelUsecase
	repo         repository.FoodMenuRepository
}

func NewFoodMenuUsecase(
	mapper mapper.FoodMenuMapper,
	hotelUsecase HotelUsecase,
	repo repository.FoodMenuRepository,
) FoodMenuUsecase {
	return &foodMenuUsecase{
		mapper:       mapper,
		hotelUsecase: hotelUsecase,
		repo:         repo,
	}
}

func (u *foodMenuUsecase) Create(foodMenuDto *dto.FoodMenuDto) (*dto.FoodMenuDto, error) {

	inst, err := db.DBInstance()
	if err != nil {
		return nil, errors.New(message.ErrInitializeDB)
	}

	isFoodMenuNotExists, err := u.validateName(inst, foodMenuDto.Name)
	if err != nil {
		return nil, errors.New(message.ErrFoodMenuExists)
	}

	if !isFoodMenuNotExists {
		return nil, errors.New(message.ErrFoodMenuExists)
	}

	isHotelExists, err := u.hotelUsecase.ValidateHotel(inst, foodMenuDto.HotelID)
	if err != nil {
		return nil, errors.New(message.ErrHotelNotFound)
	}

	if !isHotelExists {
		return nil, errors.New(message.ErrHotelNotFound)
	}

	tx := inst.Begin()

	foodMenu := u.mapper.ToFoodMenu(foodMenuDto)

	foodMenu.ID, err = utils.GenerateUUIDv7()
	if err != nil {
		return nil, errors.New(message.Failed)
	}

	if err := u.repo.Create(tx, foodMenu); err != nil {
		tx.Rollback()
		return nil, errors.New(message.ErrCreateHotelFailed)
	}

	tx.Commit()

	return u.mapper.ToFoodMenuDto(foodMenu), nil
}

func (u *foodMenuUsecase) validateName(db *gorm.DB, name string) (bool, error) {

	foodMenu, err := u.repo.FindOne(db, &repository.FoodMenuQuery{Name: name})
	if err != nil {
		return false, errors.New(message.ErrGetDataFromDB)
	}

	if foodMenu == nil {
		return false, nil
	}

	return true, nil
}
