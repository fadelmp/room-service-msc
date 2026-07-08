package mapper

import (
	"room-service-msc/domain"
	"room-service-msc/dto"
)

type FoodMenuMapper interface {
	ToFoodMenuDtoList([]*domain.FoodMenu) []*dto.FoodMenuDto
	ToFoodMenuDto(*domain.FoodMenu) *dto.FoodMenuDto
	ToFoodMenu(*dto.FoodMenuDto) *domain.FoodMenu
}

// Class
type foodMenuMapper struct{}

// Constructor
func NewFoodMenuMapper() FoodMenuMapper {
	return &foodMenuMapper{}
}

func (m *foodMenuMapper) ToFoodMenuDtoList(foodMenus []*domain.FoodMenu) []*dto.FoodMenuDto {

	foodMenuDtos := make([]*dto.FoodMenuDto, 0, len(foodMenus))
	for _, value := range foodMenus {

		if value == nil {
			continue
		}

		foodMenuDtos = append(foodMenuDtos, m.ToFoodMenuDto(value))
	}

	return foodMenuDtos
}

func (m *foodMenuMapper) ToFoodMenuDto(foodMenu *domain.FoodMenu) *dto.FoodMenuDto {

	if foodMenu == nil {
		return nil
	}

	foodMenuDto := &dto.FoodMenuDto{}
	foodMenuDto.ID = foodMenu.ID
	foodMenuDto.HotelID = foodMenu.HotelID
	foodMenuDto.Name = foodMenu.Name
	foodMenuDto.Description = foodMenu.Description
	foodMenuDto.Price = foodMenu.Price
	foodMenuDto.Category = foodMenu.Category
	foodMenuDto.IsAvailable = foodMenu.IsAvailable

	return foodMenuDto

}

func (m *foodMenuMapper) ToFoodMenu(foodMenuDto *dto.FoodMenuDto) *domain.FoodMenu {

	if foodMenuDto == nil {
		return nil
	}

	foodMenu := &domain.FoodMenu{}

	foodMenu.ID = foodMenuDto.ID
	foodMenu.HotelID = foodMenuDto.HotelID
	foodMenu.Name = foodMenuDto.Name
	foodMenu.Description = foodMenuDto.Description
	foodMenu.Price = foodMenuDto.Price
	foodMenu.Category = foodMenuDto.Category
	foodMenu.IsAvailable = foodMenuDto.IsAvailable

	return foodMenu
}
