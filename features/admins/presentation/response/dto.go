package response

import (
	"capstone/group3/features/admins"
	"time"
)

type User struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	AvatarUrl string    `json:"avatar_url" form:"avatar_url"`
	Handphone string    `json:"handphone" form:"handphone"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type Restaurant struct {
	ID            int     `json:"id" form:"id"`
	RestoName     string  `json:"resto_name" form:"resto_name"`
	RestoImageUrl string  `json:"resto_image_url" form:"resto_image_url"`
	Location      string  `json:"location" form:"location"`
	Status        string  `json:"status" form:"status"`
	Category      string  `json:"category" form:"category"`
	Rating        float64 `json:"rating" form:"rating"`
}

type DetailRestaurant struct {
	ID           int          `json:"id" form:"id"`
	RestoName    string       `json:"resto_name" form:"resto_name"`
	Category     string       `json:"category" form:"category"`
	Location     string       `json:"location" form:"location"`
	FileImageUrl string       `json:"file_image_url" form:"file_image_url"`
	Facilities   []Facility   `json:"facilities"`
	RestoImages  []RestoImage `json:"resto_images"`
	TableQuota   uint         `json:"table_quota" form:"table_quota"`
	BookingFee   uint64       `json:"booking_fee" form:"booking_fee"`
	Rating       float64      `json:"rating" form:"rating"`
	OwnerName    string       `json:"owner_name" form:"owner_name"`
	Email        string       `json:"email" form:"email"`
	Handphone    string       `json:"handphone" form:"handphone"`
	Status       string       `json:"status" form:"status"`
	Latitude     string       `json:"latitude" form:"latitude"`
	Longitude    string       `json:"longitude" form:"longitude"`
}

type Facility struct {
	Facility string `json:"facility"`
}

type RestoImage struct {
	RestoImageUrl string `json:"resto_image_url"`
}

func FormCore(data admins.Core) User {
	return User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		AvatarUrl: data.AvatarUrl,
		Handphone: data.Handphone,
		CreatedAt: data.CreatedAt,
	}
}

func FormCoreRestaurant(data admins.CoreRestaurant) Restaurant {
	return Restaurant{
		ID:            data.ID,
		RestoName:     data.RestoName,
		RestoImageUrl: data.RestoImageUrl,
		Location:      data.Location,
		Status:        data.Status,
		Category:      data.Category,
		Rating:        data.Rating,
	}
}

func FromCoreDetail(data admins.CoreDetailRestaurant) DetailRestaurant {
	return DetailRestaurant{
		ID:           data.ID,
		RestoName:    data.RestoName,
		Location:     data.Location,
		FileImageUrl: data.FileImageUrl,
		Category:     data.Category,
		TableQuota:   data.TableQuota,
		BookingFee:   data.BookingFee,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		Rating:       data.Rating,
		Status:       data.Status,
		OwnerName:    data.OwnerName,
		Email:        data.Email,
		Handphone:    data.Handphone,
		RestoImages:  FromRestoImageCoreList(data.RestoImages),
		Facilities:   FromRestoFacilityCoreList(data.Facilities),
	}
}

func FromCoreList(data []admins.Core) []User {
	result := []User{}

	for k, _ := range data {
		result = append(result, FormCore(data[k]))
	}

	return result
}

func FromCoreRestaurantList(data []admins.CoreRestaurant) []Restaurant {
	result := []Restaurant{}

	for k, _ := range data {
		result = append(result, FormCoreRestaurant(data[k]))
	}

	return result
}

func FromRestoFacilityCore(data admins.Facility) Facility {
	return Facility{
		Facility: data.Facility,
	}
}

func FromRestoFacilityCoreList(data []admins.Facility) []Facility {
	result := []Facility{}
	for key := range data {
		result = append(result, FromRestoFacilityCore(data[key]))
	}
	return result
}

func FromRestoImageCoreList(data []admins.RestoImage) []RestoImage {
	result := []RestoImage{}
	for key := range data {
		result = append(result, FromRestoImageCore(data[key]))
	}
	return result
}

func FromRestoImageCore(data admins.RestoImage) RestoImage {
	return RestoImage{
		RestoImageUrl: data.RestoImageUrl,
	}
}
