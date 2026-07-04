package mapper

import (
	"room-service-msc/domain"
	"room-service-msc/dto"
)

type HotelMapperInterface interface {
	ToHotelDtoList([]*domain.Hotel) []*dto.HotelDto
	ToHotelDto(*domain.Hotel) *dto.HotelDto
	ToHotel(*dto.HotelDto) *domain.Hotel
}

// Class
type HotelMapper struct{}

// Constructor
func NewHotelMapper() *HotelMapper {
	return &HotelMapper{}
}

func (m *HotelMapper) ToHotelDtoList(hotels []*domain.Hotel) []*dto.HotelDto {

	hotelDtos := make([]*dto.HotelDto, 0, len(hotels))
	for _, value := range hotels {

		if value == nil {
			continue
		}

		hotelDtos = append(hotelDtos, m.ToHotelDto(value))
	}

	return hotelDtos
}

func (m *HotelMapper) ToHotelDto(hotel *domain.Hotel) *dto.HotelDto {

	if hotel == nil {
		return nil
	}

	hotelDto := &dto.HotelDto{}
	hotelDto.ID = hotel.ID
	hotelDto.Name = hotel.Name
	hotelDto.Code = hotel.Code
	hotelDto.Timezone = hotel.Timezone
	hotelDto.ReceptionistWhatsapp = hotel.ReceptionistWhatsapp
	hotelDto.LogoURL = hotel.LogoURL
	hotelDto.WifiSSID = hotel.WifiSSID
	hotelDto.WifiUsername = hotel.WifiUsername
	hotelDto.WifiPassword = hotel.WifiPassword
	hotelDto.BreakfastTime = hotel.BreakfastTime
	hotelDto.BreakfastLocation = hotel.BreakfastLocation
	hotelDto.CheckoutTime = hotel.CheckoutTime
	hotelDto.LateCheckoutPolicy = hotel.LateCheckoutPolicy

	return hotelDto

}

func (m *HotelMapper) ToHotel(hotelDto *dto.HotelDto) *domain.Hotel {

	if hotelDto == nil {
		return nil
	}

	hotel := &domain.Hotel{}

	hotel.ID = hotelDto.ID
	hotel.Name = hotelDto.Name
	hotel.Code = hotelDto.Code
	hotel.Timezone = hotelDto.Timezone
	hotel.ReceptionistWhatsapp = hotelDto.ReceptionistWhatsapp
	hotel.LogoURL = hotelDto.LogoURL
	hotel.WifiSSID = hotelDto.WifiSSID
	hotel.WifiUsername = hotelDto.WifiUsername
	hotel.WifiPassword = hotelDto.WifiPassword
	hotel.BreakfastTime = hotelDto.BreakfastTime
	hotel.BreakfastLocation = hotelDto.BreakfastLocation
	hotel.CheckoutTime = hotelDto.CheckoutTime
	hotel.LateCheckoutPolicy = hotelDto.LateCheckoutPolicy

	return hotel
}
