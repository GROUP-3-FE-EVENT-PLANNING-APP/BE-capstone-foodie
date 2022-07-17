package data

import (
	"capstone/group3/features/admins"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	AvatarUrl string
	Handphone string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Restaurant struct {
	ID            int
	RestoName     string
	RestoImageUrl string
	Location      string
	Status        string
	Category      string
	Rating        float64
}

type RestaurantDetail struct {
	ID           uint
	UserID       uint
	User         User
	RestoName    string
	Location     string
	MenuImageUrl string
	Category     string
	TableQuota   uint
	BookingFee   uint64
	Latitude     string
	Longitude    string
	FileImageUrl string
	Status       string
	OwnerName    string
	Email        string
	Handphone    string
}

type RestoImage struct {
	RestoImageUrl string
}

type Facility struct {
	Facility string
}

type Comments_Ratings struct {
	Rating float64
}

func (data *User) toCore() admins.Core {
	return admins.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		AvatarUrl: data.AvatarUrl,
		Handphone: data.Handphone,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func (data *RestaurantDetail) toCoreDetail() admins.CoreDetailRestaurant {
	return admins.CoreDetailRestaurant{
		ID:           int(data.ID),
		RestoName:    data.RestoName,
		Location:     data.Location,
		Category:     data.Category,
		TableQuota:   uint(data.TableQuota),
		BookingFee:   data.BookingFee,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		FileImageUrl: data.FileImageUrl,
		Status:       data.Status,
		OwnerName:    data.OwnerName,
		Email:        data.Email,
		Handphone:    data.Handphone,
	}
}

func (data *Restaurant) toCoreRestaurant() admins.CoreRestaurant {
	return admins.CoreRestaurant{
		ID:            int(data.ID),
		RestoName:     data.RestoName,
		RestoImageUrl: data.RestoImageUrl,
		Location:      data.Location,
		Status:        data.Status,
		Category:      data.Category,
		Rating:        data.Rating,
	}
}

func toCoreList(data []User) []admins.Core {
	result := []admins.Core{}

	for key := range data {
		result = append(result, data[key].toCore())
	}

	return result
}

func toCoreRestaurantList(data []Restaurant) []admins.CoreRestaurant {
	result := []admins.CoreRestaurant{}

	for key := range data {
		result = append(result, data[key].toCoreRestaurant())
	}

	return result
}

func formCore(core admins.Core) User {
	return User{
		ID:        uint(core.ID),
		Name:      core.Name,
		Email:     core.Email,
		AvatarUrl: core.AvatarUrl,
		Handphone: core.Handphone,
	}
}
