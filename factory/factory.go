package factory

import (
	_userBusiness "capstone/group3/features/users/business"
	_userData "capstone/group3/features/users/data"
	_userPresentation "capstone/group3/features/users/presentation"

	_restaurantBusiness "capstone/group3/features/restaurants/business"
	_restaurantData "capstone/group3/features/restaurants/data"
	_restaurantPresentation "capstone/group3/features/restaurants/presentation"

	_commentBusiness "capstone/group3/features/comments/business"
	_commentData "capstone/group3/features/comments/data"
	_commentPresentation "capstone/group3/features/comments/presentation"

<<<<<<< HEAD
	_favouriteBusiness "capstone/group3/features/favourites/business"
	_favouriteData "capstone/group3/features/favourites/data"
	_favouritePresentation "capstone/group3/features/favourites/presentation"
=======
	_adminBusiness "capstone/group3/features/admins/business"
	_adminData "capstone/group3/features/admins/data"
	_adminPresentation "capstone/group3/features/admins/presentation"
>>>>>>> 4882855d35667ca1be0b3d787e87aea188bd13f0

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter       *_userPresentation.UserHandler
	RestaurantPresenter *_restaurantPresentation.RestaurantHandler
	CommentPresenter    *_commentPresentation.CommentHandler
	FavouritePresenter  *_favouritePresentation.FavouriteHandler
	AdminPresenter      *_adminPresentation.AdminHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// dbConn := config.InitDB()
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	restaurantData := _restaurantData.NewRestaurantRepository(dbConn)
	restaurantBusiness := _restaurantBusiness.NewRestaurantBusiness(restaurantData)
	restaurantPresentation := _restaurantPresentation.NewRestaurantHandler(restaurantBusiness)

	commentData := _commentData.NewCommentRepository(dbConn)
	commentBusiness := _commentBusiness.NewCommentBusiness(commentData)
	commentPresentation := _commentPresentation.NewCommentHandler(commentBusiness)

<<<<<<< HEAD
	favouriteData := _favouriteData.NewFavouriteRepository(dbConn)
	favouriteBusiness := _favouriteBusiness.NewFavouriteBusiness(favouriteData)
	favouritePresentation := _favouritePresentation.NewFavouriteHandler(favouriteBusiness)
=======
	adminData := _adminData.NewAdminRepository(dbConn)
	adminBusiness := _adminBusiness.NewAdminBusiness(adminData)
	adminPresentation := _adminPresentation.NewAdminHandler(adminBusiness)
>>>>>>> 4882855d35667ca1be0b3d787e87aea188bd13f0

	return Presenter{
		UserPresenter:       userPresentation,
		RestaurantPresenter: restaurantPresentation,
		CommentPresenter:    commentPresentation,
<<<<<<< HEAD
		FavouritePresenter:  favouritePresentation,
=======
		AdminPresenter:      adminPresentation,
>>>>>>> 4882855d35667ca1be0b3d787e87aea188bd13f0
	}
}
