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
	Lat          string `json:"lat" form:"lat" gorm:"not null; type:varchar(255)"`
	Long         string `json:"long" form:"long" gorm:"not null; type:varchar(255)"`
	Status       string `json:"status" form:"status" gorm:"not null; type:varchar(100); default:unverification"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
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
		Lat:          data.Lat,
		Long:         data.Long,
		Status:       data.Status,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func toCoreList(data []Restaurant) []restaurants.Core {
	result := []restaurants.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
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
		Lat:          core.Lat,
		Long:         core.Long,
		Status:       core.Status,
		UserID:       uint(core.User.ID),
	}
}
