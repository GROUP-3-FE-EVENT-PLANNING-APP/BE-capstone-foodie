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

func (data *Favourite) toCore_() favourites.RestoCore {
	return favourites.RestoCore{
		ID:        int(data.RestaurantID),
		RestoName: data.Restaurant.RestoName,
		Rating:    data.Restaurant.Rating,
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
