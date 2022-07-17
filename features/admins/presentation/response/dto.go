package response

import (
	"capstone/group3/features/admins"
	"time"
)

type User struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	AvatarUrl string    `json:"avatar_url" form:"avatar_url"`
	Handphone string    `json:"handphone" form:"handphone"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func FormCore(data admins.Core) User {
	return User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		AvatarUrl: data.AvatarUrl,
		Handphone: data.Handphone,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []admins.Core) []User {
	result := []User{}

	for k, _ := range data {
		result = append(result, FormCore(data[k]))
	}

	return result
}
