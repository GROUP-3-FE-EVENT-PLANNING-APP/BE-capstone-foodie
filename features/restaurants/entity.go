package restaurants

import (
	"time"
)

type Core struct {
	ID           int
	User         User
	RestoName    string
	Location     string
	MenuImageUrl string
	Category     string
	TableQuota   uint
	BookingFee   uint64
	Latitude     string
	Longitude    string
	Status       string
	FileImageUrl string
	Facility     string
	RestoImages  []RestoImage
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CoreList struct {
	ID            int
	RestoName     string
	Location      string
	Category      string
	TableQuota    uint
	Rating        float64
	RestoImageUrl string
	RestoImages   []RestoImage
}

type RestoImage struct {
	ID            int
	UserID        int
	User          User
	RestaurantID  int
	Restaurant    Core
	RestoImageUrl string
}

type User struct {
	ID        int
	Name      string
	AvatarUrl string
}

type Business interface {
	CreateRestoBusiness(data Core) (response int, err error)
	UpdateRestoBusiness(data Core, idUser int) (response int, err error)
	DetailImageRestoBusiness(id int) (imageMenu, imageFile string, err error)
	DeleteRestoBusiness(idUser int) (response int, err error)
	UploadImageRestoBusiness(data RestoImage) (response int, err error)
	AllRestoBusiness(limit, offset int) (result []CoreList, err error)
}

type Data interface {
	InsertRestoData(data Core) (response int, err error)
	UpdateRestoData(data Core, idUser int) (response int, err error)
	DetailImageRestoData(id int) (imageMenu, imageFile string, err error)
	DeleteRestoData(idUser int) (response int, err error)
	UploadImageRestoData(data RestoImage) (response int, err error)
	AllRestoData(limit, offset int) (result []CoreList, err error)
	RatingData(idResto int) (result float64, err error)
	RestoImageData(idResto int) (result string, err error)
}
