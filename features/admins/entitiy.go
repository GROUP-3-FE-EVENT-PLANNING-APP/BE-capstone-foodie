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

type Business interface {
	AllUserBusiness(limit, offset, idUser int) (result []Core, err error)
}

type Data interface {
	AllUserData(limit, offset, idUser int) (result []Core, err error)
}
