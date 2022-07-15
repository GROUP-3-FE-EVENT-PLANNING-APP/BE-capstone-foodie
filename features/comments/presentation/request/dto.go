package request

import (
	"capstone/group3/features/comments"
)

type Comment struct {
	Comment string  `json:"comment" form:"comment"`
	Rating  float64 `json:"rating" form:"rating"`
}

func ToCore(req Comment) comments.Core {
	return comments.Core{
		Comment: req.Comment,
		Rating:  req.Rating,
	}
}
