package comments

import "time"

type Core struct {
	ID         int
	User       User
	Restaurant Restaurant
	Comment    string
	Rating     float64
	CreatedAt  time.Time
	DeletedAt  time.Time
}

type User struct {
	ID        int
	Name      string
	AvatarUrl string
}

type Restaurant struct {
	ID   int
	Name string
}

type Business interface {
	CreateComment(input Core) (row int, err error)
	GetCommentByIdResto(idResto, limitint, offsetint int) (data []Core, err error)
	GetRatingByIdResto(idResto int) (rating float64, err error)
}

type Data interface {
	InsertComment(input Core) (row int, err error)
	SelectCommentByIdResto(idResto, limitint, offsetint int) (data []Core, err error)
	SelectRatingByIdResto(idResto int) (rating float64, err error)
}
