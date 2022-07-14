package users

import "time"

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	AvatarUrl string
	Handphone string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Business interface {
	CreateData(input Core) (row int, err error)
	LoginUser(authData AuthRequestData) (id int, token, name, avatarUrl, role, handphone, email string, err error)
	UpdateData(input Core, idFromToken int) (row int, err error)
	GetUserByMe(idFromToken int) (data Core, err error)
	DeleteDataById(idFromToken int) (row int, err error)
}

type Data interface {
	InsertData(input Core) (row int, err error)
	LoginUserDB(authData AuthRequestData) (id int, token, name, avatarUrl, role, handphone, email string, err error)
	UpdateDataDB(data map[string]interface{}, idFromToken int) (row int, err error)
	SelectDataByMe(idFromToken int) (data Core, err error)
	DeleteDataByIdDB(idFromToken int) (row int, err error)
}
