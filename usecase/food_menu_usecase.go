package usecase

import (
	"errors"
	"room-service-msc/dto"
	"room-service-msc/infrastructure/db"
	"room-service-msc/infrastructure/message"
	"room-service-msc/mapper"
	"room-service-msc/repository"
	"room-service-msc/service"
	"room-service-msc/utils"

	"gorm.io/gorm"
)

type FoodMenuUsecase interface {
	Create(*dto.FoodMenuDto) (*dto.FoodMenuDto, error)
}

type foodMenuUsecase struct {
	mapper       mapper.FoodMenuMapper
	hotelService service.HotelService
	repo         repository.FoodMenuRepository
}

func NewFoodMenuUsecase(
	mapper mapper.FoodMenuMapper,
	hotelService service.HotelService,
	repo repository.FoodMenuRepository,
) FoodMenuUsecase {
	return &foodMenuUsecase{
		mapper:       mapper,
		hotelService: hotelService,
		repo:         repo,
	}
}

func (u *foodMenuUsecase) Create(foodMenuDto *dto.FoodMenuDto) (*dto.FoodMenuDto, error) {

	inst, err := db.DBInstance()
	if err != nil {
		return nil, errors.New(message.ErrInitializeDB)
	}

	err = u.validateName(inst, foodMenuDto.Name, foodMenuDto.HotelID)
	if err != nil {
		return nil, errors.New(message.ErrFoodMenuExists)
	}

	err = u.hotelService.IsExistsHotel(inst, foodMenuDto.HotelID)
	if err != nil {
		return nil, err
	}

	foodMenu := u.mapper.ToFoodMenu(foodMenuDto)

	foodMenu.ID, err = utils.GenerateUUIDv7()
	if err != nil {
		return nil, errors.New(message.Failed)
	}

	tx := inst.Begin()

	if err := u.repo.Create(tx, foodMenu); err != nil {
		tx.Rollback()
		return nil, errors.New(message.ErrCreateHotelFailed)
	}

	tx.Commit()

	return u.mapper.ToFoodMenuDto(foodMenu), nil
}

func (u *foodMenuUsecase) validateName(db *gorm.DB, name, hotelID string) error {

	foodMenu, err := u.repo.FindOne(db, &repository.FoodMenuQuery{Name: name, HotelID: hotelID})
	if err != nil {
		return errors.New(message.ErrGetDataFromDB)
	}

	if foodMenu == nil {
		return message.ErrHotelNotFound
	}

	return nil
}
