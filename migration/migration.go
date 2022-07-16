package migration

import (
	_mComments "capstone/group3/features/comments/data"
	_mRestaurants "capstone/group3/features/restaurants/data"
	_mUsers "capstone/group3/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mRestaurants.Restaurant{})
	db.AutoMigrate(&_mRestaurants.Facility{})
	db.AutoMigrate(&_mRestaurants.RestoImage{})
	db.AutoMigrate(&_mComments.Comments_Ratings{})
}
