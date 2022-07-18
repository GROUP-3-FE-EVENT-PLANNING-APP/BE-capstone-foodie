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

type Restaurant_ struct {
	ID          int          `json:"id" form:"id"`
	RestoName   string       `json:"resto_name" form:"resto_name"`
	Location    string       `json:"location" form:"location"`
	Category    string       `json:"category" form:"category"`
	TableQuota  uint         `json:"table_quota" form:"table_quota"`
	Rating      float64      `json:"rating" form:"rating"`
	RestoImages []RestoImage `json:"resto_images"`
}

type RestaurantDetail struct {
	ID           int          `json:"id" form:"id"`
	RestoName    string       `json:"resto_name" form:"resto_name"`
	Category     string       `json:"category" form:"category"`
	Location     string       `json:"location" form:"location"`
	MenuImageUrl string       `json:"menu_image_url" form:"menu_image_url"`
	FileImageUrl string       `json:"file_image_url" form:"file_image_url"`
	Facilities   []Facility   `json:"facilities"`
	RestoImages  []RestoImage `json:"resto_images"`
	TableQuota   uint         `json:"table_quota" form:"table_quota"`
	BookingFee   uint64       `json:"booking_fee" form:"booking_fee"`
	Rating       float64      `json:"rating" form:"rating"`
	Comments     []Comment    `json:"comments"`
	Latitude     string       `json:"latitude" form:"latitude"`
	Longitude    string       `json:"longitude" form:"langitude"`
}

type RestaurantMyDetail struct {
	ID            int     `json:"id" form:"id"`
	Rating        float64 `json:"rating" form:"rating"`
	RestoName     string  `json:"resto_name" form:"resto_name"`
	Location      string  `json:"location" form:"location"`
	Category      string  `json:"category" form:"category"`
	Status        string  `json:"status" form:"status"`
	RestoImageUrl string  `json:"resto_image_url" form:"resto_image_url"`
}

type Comment struct {
	UserID  int    `json:"user_id"`
	Comment string `json:"comment"`
}

type Facility struct {
	Facility string `json:"facility"`
}

type RestoImage struct {
	RestoImageUrl string `json:"resto_image_url"`
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

func FromCoreDetail(data restaurants.CoreDetail) RestaurantDetail {
	return RestaurantDetail{
		ID:           data.ID,
		RestoName:    data.RestoName,
		Location:     data.Location,
		MenuImageUrl: data.MenuImageUrl,
		FileImageUrl: data.FileImageUrl,
		Category:     data.Category,
		TableQuota:   data.TableQuota,
		BookingFee:   data.BookingFee,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		Rating:       data.Rating,
		RestoImages:  FromRestoImageCoreList(data.RestoImages),
		Facilities:   FromRestoFacilityCoreList(data.Facilities),
		Comments:     FromRestoCommentCoreList(data.Comments),
	}
}

func FromCoreDetailMyResto(data restaurants.CoreMyDetail) RestaurantMyDetail {
	return RestaurantMyDetail{
		ID:            data.ID,
		RestoName:     data.RestoName,
		Location:      data.Location,
		Rating:        data.Rating,
		Status:        data.Status,
		Category:      data.Category,
		RestoImageUrl: data.RestoImageUrl,
	}
}

func FromCoreAll(data restaurants.CoreList) Restaurant_ {
	return Restaurant_{
		ID:          data.ID,
		RestoName:   data.RestoName,
		Location:    data.Location,
		Category:    data.Category,
		TableQuota:  data.TableQuota,
		Rating:      data.Rating,
		RestoImages: FromRestoImageCoreList(data.RestoImages)}
}

func FromRestoImageCore(data restaurants.RestoImage) RestoImage {
	return RestoImage{
		RestoImageUrl: data.RestoImageUrl,
	}
}

func FromRestoFacilityCore(data restaurants.Facility) Facility {
	return Facility{
		Facility: data.Facility,
	}
}

func FromRestoFacilityCoreList(data []restaurants.Facility) []Facility {
	result := []Facility{}
	for key := range data {
		result = append(result, FromRestoFacilityCore(data[key]))
	}
	return result
}

func FromRestoImageCoreList(data []restaurants.RestoImage) []RestoImage {
	result := []RestoImage{}
	for key := range data {
		result = append(result, FromRestoImageCore(data[key]))
	}
	return result
}

func FromCoreList(data []restaurants.CoreList) []Restaurant_ {
	result := []Restaurant_{}
	for key := range data {
		result = append(result, FromCoreAll(data[key]))
	}
	return result
}

func FromRestoCommentCore(data restaurants.Comment) Comment {
	return Comment{
		UserID:  data.UserID,
		Comment: data.Comment,
	}
}

func FromRestoCommentCoreList(data []restaurants.Comment) []Comment {
	result := []Comment{}
	for key := range data {
		result = append(result, FromRestoCommentCore(data[key]))
	}
	return result
}
