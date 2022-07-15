package request

import (
	"capstone/group3/features/restaurants"
)

type Restaurant struct {
	RestoName    string `json:"resto_name" form:"resto_name" validate:"required,min=3"`
	Location     string `json:"location" form:"location" validate:"required"`
	MenuImageUrl string `json:"menu_image_url" form:"menu_image_url"`
	Category     string `json:"category" form:"category" validate:"required"`
	TableQuota   uint   `json:"table_quota" form:"table_quota" validate:"required,number"`
	BookingFee   uint64 `json:"booking_fee" form:"booking_fee" validate:"required,number"`
	Latitude     string `json:"latitude" form:"latitude" validate:"required"`
	Longitude    string `json:"longitude" form:"longitude" validate:"required"`
	Status       string `json:"status" form:"status" `
	UserId       int    `json:"user_id" form:"user_id"`
	Facility     string `json:"facility" form:"facility" validate:"required"`
}

func ToCore(req Restaurant) restaurants.Core {
	return restaurants.Core{
		RestoName:    req.RestoName,
		Location:     req.Location,
		MenuImageUrl: req.MenuImageUrl,
		Category:     req.Category,
		TableQuota:   req.TableQuota,
		BookingFee:   req.BookingFee,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		Status:       req.Status,
		Facility:     req.Facility,
		User: restaurants.User{
			ID: req.UserId,
		},
	}
}
