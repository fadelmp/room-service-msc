package usecase

import (
	"errors"
	"fmt"
	"room-service-msc/domain"
	"room-service-msc/dto"
	"room-service-msc/infrastructure/message"
	"room-service-msc/mapper"
	"room-service-msc/repository"
	"room-service-msc/utils"

	"gorm.io/gorm"
)

type HotelUsecase interface {
	CreateDefaultHotel(*dto.CreateHotelDto) (*dto.HotelDto, error)
}

type hotelUsecase struct {
	mapper        mapper.HotelMapper
	hotelRepo     repository.HotelRepository
	hotelCodeRepo repository.HotelCodeRepository
}

func NewHotelUsecase(
	mapper mapper.HotelMapper,
	hotelRepo repository.HotelRepository,
	hotelCodeRepo repository.HotelCodeRepository,
) HotelUsecase {
	return &hotelUsecase{
		mapper:        mapper,
		hotelRepo:     hotelRepo,
		hotelCodeRepo: hotelCodeRepo,
	}
}

func (u *hotelUsecase) CreateDefaultHotel(hotelDto *dto.CreateHotelDto) (*dto.HotelDto, error) {

	// inst, err := db.DBInstance()
	// if err != nil {
	// 	return nil, errors.New(message.ErrInitializeDB)
	// }

	// hotel, err := u.checkHotelName(inst, hotelDto)
	// if err != nil {
	// 	return nil, err
	// }

	// tx := inst.Begin()

	// hotelCode, err := u.generateHotelCode(tx, hotelDto)
	// if err != nil {
	// 	return nil, errors.New(message.ErrGenerateHotelCode)
	// }

	// hotel := u.mapper.ToHotel(hotelDto)

	// tx, err := db.BeginTransaction()
	// if err != nil {
	// 	return nil, err
	// }

	// if err := u.hotelRepo.Create(tx, hotel); err != nil {
	// 	tx.Rollback()
	// 	return nil, errors.New(message.ErrCreateHotelFailed)
	// }

	// tx.Commit()
	// return u.mapper.ToHotelDto(hotel), nil
	return nil, nil
}

func (u *hotelUsecase) checkHotelName(inst *gorm.DB, hotelDto *dto.CreateHotelDto) (*domain.Hotel, error) {

	hotel, err := u.hotelRepo.FindOne(inst, &repository.HotelQuery{Name: hotelDto.Name})
	if err != nil {
		return nil, errors.New(message.ErrHotelExists)
	}

	return hotel, nil
}

func (u *hotelUsecase) generateHotelCode(tx *gorm.DB, hotelDto *dto.CreateHotelDto) (*string, error) {

	prefix := utils.GetCode(hotelDto.Name)

	hotelCodeCount, err := u.hotelCodeRepo.Count(tx, &repository.HotelCodeQuery{
		Code: prefix,
	})
	if err != nil {
		return nil, errors.New(message.ErrBeginTrx)
	}

	newSequence := hotelCodeCount + 1
	hotelCode := fmt.Sprintf("%s-%03d", prefix, newSequence)

	return &hotelCode, nil
}
