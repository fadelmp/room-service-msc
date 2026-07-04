package usecase

import (
	"errors"
	"room-service-msc/dto"
	"room-service-msc/infrastructure/db"
	"room-service-msc/infrastructure/message"
	"room-service-msc/mapper"
	"room-service-msc/repository"
)

type HotelUsecaseInterface interface {
	CreateDefaultHotel(*dto.HotelDto) (*dto.HotelDto, error)
}

type HotelUsecase struct {
	mapper mapper.HotelMapper
	repo   repository.HotelRepository
}

func NewHotelUsecase(mapper mapper.HotelMapper, repo repository.HotelRepository) HotelUsecase {
	return HotelUsecase{
		mapper: mapper,
		repo:   repo,
	}
}

func (u *HotelUsecase) CreateDefaultHotel(hotelDto *dto.HotelDto) (*dto.HotelDto, error) {

	hotel := u.mapper.ToHotel(hotelDto)

	tx, err := db.BeginTransaction()
	if err != nil {
		return nil, err
	}

	if err := u.repo.Create(tx, hotel); err != nil {
		tx.Rollback()
		return nil, errors.New(message.ErrCreateHotelFailed)
	}

	tx.Commit()
	return u.mapper.ToHotelDto(hotel), nil
}
