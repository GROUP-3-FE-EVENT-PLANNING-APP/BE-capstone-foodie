package admins

import (
	"time"
)

type Core struct {
	ID        int
	Name      string
	Email     string
	AvatarUrl string
	Handphone string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CoreRestaurant struct {
	ID            int
	RestoName     string
	RestoImageUrl string
	Location      string
	Status        string
	Category      string
	Rating        float64
}

type CoreDetailRestaurant struct {
	ID           int
	RestoName    string
	Category     string
	Location     string
	FileImageUrl string
	Facilities   []Facility
	RestoImages  []RestoImage
	TableQuota   uint
	BookingFee   uint64
	Rating       float64
	OwnerName    string
	Email        string
	Handphone    string
	Status       string
	Latitude     string
	Longitude    string
}

type Comment struct {
	Comment string
}

type Facility struct {
	Facility string
}

type RestoImage struct {
	RestoImageUrl string
}

type Business interface {
	AllUserBusiness(limit, offset, idUser int) (result []Core, err error)
	AllRestoBusiness(limit, offset, idUser int) (result []CoreRestaurant, err error)
	DetailRestoBusiness(id int, idUser int) (result CoreDetailRestaurant, err error)
}

type Data interface {
	AllUserData(limit, offset, idUser int) (result []Core, err error)
	AllRestoData(limit, offset, idUser int) (result []CoreRestaurant, err error)
	RestoImageData(id int) (img string, err error)
	RatingData(idResto int) (result float64, err error)
	DetailRestoData(id int, idUser int) (result CoreDetailRestaurant, err error)
	RestoImagesData(idResto int) (result []string, err error)
	FacilitiesData(idResto int) (result []string, err error)
}
