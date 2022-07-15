package data

import (
	"capstone/group3/features/comments"
	"time"

	"gorm.io/gorm"
)

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
	Name             string
	AvatarUrl        string
	Comments_Ratings []Comments_Ratings `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type Restaurant struct {
	gorm.Model
	Name             string
	Comments_Ratings []Comments_Ratings `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE"`
}

// DTO

func (data *Comments_Ratings) toCore() comments.Core {
	return comments.Core{
		ID: int(data.ID),
		User: comments.User{
			ID:        int(data.UserID),
			Name:      data.User.Name,
			AvatarUrl: data.User.AvatarUrl,
		},
		Restaurant: comments.Restaurant{
			ID:   int(data.RestaurantID),
			Name: data.Restaurant.Name,
		},
		Comment:   data.Comment,
		Rating:    data.Rating,
		CreatedAt: data.CreatedAt,
	}
}

func toCoreList(data []Comments_Ratings) []comments.Core {
	result := []comments.Core{}
	for k := range data {
		result = append(result, data[k].toCore())
	}
	return result
}

func FromCore(core comments.Core) Comments_Ratings {
	return Comments_Ratings{
		UserID: uint(core.User.ID),
		User: User{
			Name:      core.User.Name,
			AvatarUrl: core.User.AvatarUrl,
		},
		RestaurantID: uint(core.Restaurant.ID),
		Restaurant: Restaurant{
			Name: core.Restaurant.Name,
		},
		Comment:   core.Comment,
		Rating:    core.Rating,
		CreatedAt: core.CreatedAt,
	}
}
