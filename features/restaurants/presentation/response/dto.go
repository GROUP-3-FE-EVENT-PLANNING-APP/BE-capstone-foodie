package response

import (
	"capstone/group3/features/restaurants"
	"time"
)

type Restaurant struct {
	ID           int       `json:"id" form:"id"`
	RestoName    string    `json:"resto_name" form:"resto_name"`
	Location     string    `json:"location" form:"location"`
	MenuImageUrl string    `json:"menu_image_url" form:"menu_image_url"`
	Category     string    `json:"category" form:"category"`
	TableQuota   uint      `json:"table_quota" form:"table_quota"`
	BookingFee   uint64    `json:"booking_fee" form:"booking_fee"`
	Latitude     string    `json:"latitude" form:"latitude"`
	Longitude    string    `json:"longitude" form:"longitude"`
	Status       string    `json:"status" form:"status"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	User         User      `json:"user" form:"user"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

func FromCore(data restaurants.Core) Restaurant {
	return Restaurant{
		ID:           data.ID,
		RestoName:    data.RestoName,
		Location:     data.Location,
		MenuImageUrl: data.MenuImageUrl,
		Category:     data.Category,
		TableQuota:   data.TableQuota,
		BookingFee:   data.BookingFee,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		Status:       data.Status,
		CreatedAt:    data.CreatedAt,
		User: User{
			ID:        data.User.ID,
			Name:      data.User.Name,
			AvatarUrl: data.User.AvatarUrl,
		},
	}
}

func FromCoreList(data []restaurants.Core) []Restaurant {
	result := []Restaurant{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
