package business

import (
	"capstone/group3/features/restaurants"
	"errors"
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

func (uc *restaurantUseCase) UpdateRestoBusiness(editData restaurants.Core, idUser int) (response int, err error) {
	if editData.RestoName == "" || editData.Location == "" || editData.MenuImageUrl == "" || editData.Category == "" || editData.TableQuota == 0 || editData.BookingFee == 0 || editData.Latitude == "" || editData.Longitude == "" || editData.Facility == "" {
		return 0, errors.New("all input data must be filled")
	}

	response, err = uc.restaurantData.UpdateRestoData(editData, idUser)

	return response, err
}

func (uc *restaurantUseCase) DetailImageRestoBusiness(id int) (imageMenu, imageFile string, err error) {
	imageMenu, imageFile, err = uc.restaurantData.DetailImageRestoData(id)

	return imageMenu, imageFile, err
}
