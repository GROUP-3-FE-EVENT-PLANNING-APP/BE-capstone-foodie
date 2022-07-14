package response

import (
	"capstone/group3/features/users"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarUrl string    `json:"avatar_url"`
	Handphone string    `json:"handphone"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func FromCore(data users.Core) User {
	return User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		AvatarUrl: data.AvatarUrl,
		Handphone: data.Handphone,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for k, _ := range data {
		result = append(result, FromCore(data[k]))
	}
	return result
}
