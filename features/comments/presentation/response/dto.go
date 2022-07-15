package response

import (
	"capstone/group3/features/comments"
	"time"
)

type Comment struct {
	Name      string    `json:"name"`
	AvatarUrl string    `json:"avatar_url"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

func FromCore(data comments.Core) Comment {
	return Comment{
		Name:      data.User.Name,
		AvatarUrl: data.User.AvatarUrl,
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []comments.Core) []Comment {
	result := []Comment{}
	for k, _ := range data {
		result = append(result, FromCore(data[k]))
	}
	return result
}

type RatingResto struct {
	Rating float64 `json:"rating"`
}
