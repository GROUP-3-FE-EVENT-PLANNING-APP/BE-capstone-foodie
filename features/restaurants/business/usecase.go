package business

import (
	"capstone/group3/features/restaurants"
)

type restaurantUseCase struct {
	restaurantData restaurants.Data
}

func NewRestaurantBusiness(rstData restaurants.Data) restaurants.Business {
	return &restaurantUseCase{
		restaurantData: rstData,
	}
}

func (uc *restaurantUseCase) CreateRestoBusiness(newData restaurants.Core) (response int, err error) {
	response, err = uc.restaurantData.InsertRestoData(newData)

	return response, err
}
