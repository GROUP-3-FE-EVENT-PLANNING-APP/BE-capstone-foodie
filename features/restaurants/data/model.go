package data

import (
	"capstone/group3/features/restaurants"
	"time"

	"gorm.io/gorm"
)

type Restaurant struct {
	// gorm.Model
	ID           uint `gorm:"primaryKey;autoIncrement"`
	UserID       uint `json:"user_id" form:"user_id" gorm:"unique"`
	User         User
	RestoName    string `json:"resto_name" form:"resto_name" gorm:"not null; type:varchar(255); unique"`
	Location     string `json:"location" form:"location" gorm:"not null; type:text"`
	MenuImageUrl string `json:"menu_image_url" form:"menu_image_url" gorm:"not null; type:varchar(255)"`
	Category     string `json:"category" form:"category" gorm:"not null; type:varchar(100)"`
	TableQuota   uint   `json:"table_quota" form:"table_quota" gorm:"not null; type:integer"`
	BookingFee   uint64 `json:"booking_fee" form:"booking_fee" gorm:"not null; type:bigint(20)"`
	Latitude     string `json:"latitude" form:"latitude" gorm:"not null; type:varchar(255)"`
	Longitude    string `json:"longitude" form:"longitude" gorm:"not null; type:varchar(255)"`
	Status       string `json:"status" form:"status" gorm:"not null; type:varchar(100); default:unverification"`
	FileImageUrl string `json:"file_image_url" form:"file_image_url" gorm:"not null; type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Facilities   []Facility     `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE"`
	RestoImages  []RestoImage   `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE"`
}

type RestaurantData struct {
	ID            uint
	RestoName     string `json:"resto_name" form:"resto_name"`
	Location      string `json:"location" form:"location"`
	Category      string `json:"category" form:"category"`
	TableQuota    uint   `json:"table_quota" form:"table_quota"`
	RestoImageUrl string `json:"resto_image_url" form:"resto_image_url"`
	RestoImages   []RestoImage
}

type RestaurantDetail struct {
	ID            uint
	Rating        float64 `json:"rating" form:"rating"`
	RestoName     string  `json:"resto_name" form:"resto_name"`
	Location      string  `json:"location" form:"location"`
	Category      string  `json:"category" form:"category"`
	Status        string  `json:"status" form:"status"`
	RestoImageUrl string  `json:"resto_image_url" form:"resto_image_url"`
}

type RestoImage struct {
	// gorm.Model
	ID            uint `gorm:"primaryKey;autoIncrement"`
	RestaurantID  uint `json:"restaurant_id" form:"restaurant_id"`
	Restaurant    Restaurant
	RestoImageUrl string `json:"resto_image_url" form:"resto_image_url" gorm:"not null; type:varchar(255);"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Facility struct {
	// gorm.Model
	ID           uint `gorm:"primaryKey;autoIncrement"`
	RestaurantID uint `json:"restaurant_id" form:"restaurant_id"`
	Restaurant   Restaurant
	Facility     string `json:"facility" form:"facility" gorm:"not null; type:varchar(255);"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
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
	Name        string       `json:"name" form:"name"`
	AvatarUrl   string       `json:"avatar_url" form:"avatar_url"`
	Restaurants []Restaurant `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (data *Restaurant) toCore() restaurants.Core {
	return restaurants.Core{
		ID: int(data.ID),
		User: restaurants.User{
			ID:        int(data.UserID),
			Name:      data.User.Name,
			AvatarUrl: data.User.AvatarUrl,
		},
		RestoName:    data.RestoName,
		Location:     data.Location,
		MenuImageUrl: data.MenuImageUrl,
		Category:     data.Category,
		TableQuota:   data.TableQuota,
		BookingFee:   data.BookingFee,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		Status:       data.Status,
		FileImageUrl: data.FileImageUrl,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func (data *Restaurant) toCoreDetail() restaurants.CoreDetail {
	return restaurants.CoreDetail{
		ID:           int(data.ID),
		RestoName:    data.RestoName,
		Location:     data.Location,
		MenuImageUrl: data.MenuImageUrl,
		Category:     data.Category,
		TableQuota:   uint(data.TableQuota),
		BookingFee:   data.BookingFee,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		FileImageUrl: data.FileImageUrl,
	}
}

func (data *RestaurantDetail) toCoreMyResto() restaurants.CoreMyDetail {
	return restaurants.CoreMyDetail{
		ID:            int(data.ID),
		RestoName:     data.RestoName,
		Location:      data.Location,
		RestoImageUrl: data.RestoImageUrl,
		Rating:        data.Rating,
		Category:      data.Category,
		Status:        data.Status,
	}
}

func (data *Comments_Ratings) toCoreComment() restaurants.Comment {
	return restaurants.Comment{
		ID:      int(data.ID),
		UserID:  int(data.UserID),
		Comment: data.Comment,
	}
}

func (data *Restaurant) toCore_() restaurants.CoreList {
	return restaurants.CoreList{
		ID:         int(data.ID),
		RestoName:  data.RestoName,
		Location:   data.Location,
		Category:   data.Category,
		TableQuota: data.TableQuota,
	}
}

func (data *RestoImage) toCoreRestoImage() restaurants.RestoImage {
	return restaurants.RestoImage{
		ID:            int(data.ID),
		RestoImageUrl: data.RestoImageUrl,
	}
}

func toCoreListImage(data []RestoImage) []restaurants.RestoImage {
	result := []restaurants.RestoImage{}
	for key := range data {
		result = append(result, data[key].toCoreRestoImage())
	}
	return result
}

func toCoreList(data []Restaurant) []restaurants.CoreList {
	result := []restaurants.CoreList{}
	for key := range data {
		result = append(result, data[key].toCore_())
	}
	return result
}

func toCoreCommentList(data []Comments_Ratings) []restaurants.Comment {
	result := []restaurants.Comment{}
	for key := range data {
		result = append(result, data[key].toCoreComment())
	}
	return result
}

func fromCore(core restaurants.Core) Restaurant {
	return Restaurant{
		RestoName:    core.RestoName,
		Location:     core.Location,
		MenuImageUrl: core.MenuImageUrl,
		Category:     core.Category,
		TableQuota:   core.TableQuota,
		BookingFee:   core.BookingFee,
		Latitude:     core.Latitude,
		Longitude:    core.Longitude,
		Status:       core.Status,
		FileImageUrl: core.FileImageUrl,
		UserID:       uint(core.User.ID),
	}
}

func fromCoreRestoImage(core restaurants.RestoImage) RestoImage {
	return RestoImage{
		RestaurantID:  uint(core.RestaurantID),
		RestoImageUrl: core.RestoImageUrl,
	}
}
