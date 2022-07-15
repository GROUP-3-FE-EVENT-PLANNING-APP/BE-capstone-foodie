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
	Facility     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type User struct {
	ID        int
	Name      string
	AvatarUrl string
}

type Business interface {
	CreateRestoBusiness(data Core) (response int, err error)
}

type Data interface {
	InsertRestoData(data Core) (response int, err error)
}
