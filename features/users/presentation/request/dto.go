package request

import "capstone/group3/features/users"

type User struct {
	Name      string `json:"name" form:"name" validate:"required,min=2"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required"`
	AvatarUrl string `form:"avatar_url"`
	Handphone string `json:"handphone" form:"handphone" validate:"required,min=10"`
}

func ToCore(req User) users.Core {
	return users.Core{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		AvatarUrl: req.AvatarUrl,
		Handphone: req.Handphone,
	}
}
