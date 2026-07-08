package service

import (
	"room-service-msc/infrastructure/message"
	"room-service-msc/repository"

	"gorm.io/gorm"
)

type HotelService interface {
	IsExistsHotel(*gorm.DB, string) error
}

type hotelService struct {
	repository repository.HotelRepository
}

func NewHotelService(repository repository.HotelRepository) HotelService {
	return &hotelService{
		repository: repository,
	}
}

func (u *hotelService) IsExistsHotel(db *gorm.DB, hotelID string) error {

	hotel, err := u.repository.FindOne(db, &repository.HotelQuery{ID: hotelID})
	if err != nil {
		return message.ErrValidateHotel
	}

	if hotel == nil {
		return message.ErrHotelNotFound
	}

	return nil
}
