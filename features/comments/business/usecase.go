package business

import (
	"capstone/group3/features/comments"
	"errors"
)

type commentUseCase struct {
	commentData comments.Data
}

func NewCommentBusiness(dataComment comments.Data) comments.Business {
	return &commentUseCase{
		commentData: dataComment,
	}
}

func (uc *commentUseCase) CreateComment(input comments.Core) (row int, err error) {

	if input.Comment == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	} else if input.Rating < 1 {
		return -2, errors.New("rating cant be lower than 1")
	} else if input.Rating > 5 {
		return -3, errors.New("rating cant be greater than 5")
	}
	row, err = uc.commentData.InsertComment(input)
	return row, err
}

func (uc *commentUseCase) GetCommentByIdResto(idResto, limitint, offsetint int) (data []comments.Core, err error) {
	data, err = uc.commentData.SelectCommentByIdResto(idResto, limitint, offsetint)
	return data, err
}

func (uc *commentUseCase) GetRatingByIdResto(idResto int) (rating float64, err error) {
	rating, err = uc.commentData.SelectRatingByIdResto(idResto)
	return rating, err
}
