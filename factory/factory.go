package factory

import (
	_userBusiness "capstone/group3/features/users/business"
	_userData "capstone/group3/features/users/data"
	_userPresentation "capstone/group3/features/users/presentation"

	_restaurantBusiness "capstone/group3/features/restaurants/business"
	_restaurantData "capstone/group3/features/restaurants/data"
	_restaurantPresentation "capstone/group3/features/restaurants/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter       *_userPresentation.UserHandler
	RestaurantPresenter *_restaurantPresentation.RestaurantHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// dbConn := config.InitDB()
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	restaurantData := _restaurantData.NewRestaurantRepository(dbConn)
	restaurantBusiness := _restaurantBusiness.NewRestaurantBusiness(restaurantData)
	restaurantPresentation := _restaurantPresentation.NewRestaurantHandler(restaurantBusiness)

	return Presenter{
		UserPresenter:       userPresentation,
		RestaurantPresenter: restaurantPresentation,
	}
}
