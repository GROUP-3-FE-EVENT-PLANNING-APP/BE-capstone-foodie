package data

import (
	"capstone/group3/features/admins"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	AvatarUrl string
	Handphone string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (data *User) toCore() admins.Core {
	return admins.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		AvatarUrl: data.AvatarUrl,
		Handphone: data.Handphone,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []User) []admins.Core {
	result := []admins.Core{}

	for key := range data {
		result = append(result, data[key].toCore())
	}

	return result
}

func formCore(core admins.Core) User {
	return User{
		ID:        uint(core.ID),
		Name:      core.Name,
		Email:     core.Email,
		AvatarUrl: core.AvatarUrl,
		Handphone: core.Handphone,
	}
}
