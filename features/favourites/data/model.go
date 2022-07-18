package data

import (
	"capstone/group3/features/favourites"
	"time"

	"gorm.io/gorm"
)

type Favourite struct {
	gorm.Model
	UserID       int
	User         User
	RestaurantID int
	Restaurant   Restaurant
}

type Comments_Ratings struct {
	gorm.Model
	UserID       uint
	User         User
	RestaurantID uint
	Restaurant   Restaurant
	Comment      string
	Rating       float64
	CreatedAt    time.Time
}

type User struct {
	gorm.Model
	Name      string
	Favourite []Favourite `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type Restaurant struct {
	gorm.Model
	RestoName     string       `gorm:"not null; type:varchar(255); unique"`
	RestoImages   []RestoImage `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE"`
	Rating        float64
	Category      string
	Location      string
	Favourite     []Favourite `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE"`
	RestoImageUrl string
}

type RestoImage struct {
	// gorm.Model
	gorm.Model
	RestaurantID  uint
	Restaurant    Restaurant
	RestoImageUrl string
}

// DTO

// func (data *Favourite) toCore() favourites.Core {
// 	return favourites.Core{
// 		UserID: data.UserID,
// 		RestaurantID: data.RestaurantID,
// 	}
// }

// func toCoreList(data []Comments_Ratings) []comments.Core {
// 	result := []comments.Core{}
// 	for k := range data {
// 		result = append(result, data[k].toCore())
// 	}
// 	return result
// }

// func FromCore(core comments.Core) Comments_Ratings {
// 	return Comments_Ratings{
// 		UserID: uint(core.User.ID),
// 		User: User{
// 			Name:      core.User.Name,
// 			AvatarUrl: core.User.AvatarUrl,
// 		},
// 		RestaurantID: uint(core.Restaurant.ID),
// 		Restaurant: Restaurant{
// 			Name: core.Restaurant.Name,
// 		},
// 		Comment:   core.Comment,
// 		Rating:    core.Rating,
// 		CreatedAt: core.CreatedAt,
// 	}
// }

func (data *Favourite) toCore_() favourites.RestoCore {
	return favourites.RestoCore{
		ID:        int(data.ID),
		RestoName: data.Restaurant.RestoName,
		Location:  data.Restaurant.Location,
		Category:  data.Restaurant.Category,
	}
}

func toCoreList(data []Favourite) []favourites.RestoCore {
	result := []favourites.RestoCore{}
	for key := range data {
		result = append(result, data[key].toCore_())
	}
	return result
}
