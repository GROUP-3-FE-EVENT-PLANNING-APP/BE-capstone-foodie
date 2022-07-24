package data

import (
	"capstone/group3/features/comments"

	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) comments.Data {
	return &mysqlCommentRepository{
		DB: db,
	}
}

func (repo *mysqlCommentRepository) InsertComment(input comments.Core) (row int, err error) {
	comment := FromCore(input)
	resultInsert := repo.DB.Create(&comment)
	if resultInsert.Error != nil {
		return 0, resultInsert.Error
	}
	return int(resultInsert.RowsAffected), nil
}

func (repo *mysqlCommentRepository) SelectCommentByIdResto(idResto, limitint, offsetint int) (data []comments.Core, err error) {
	dataComment := []Comments_Ratings{}
	result := repo.DB.Limit(limitint).Offset(offsetint).Preload("User").Order("created_at DESC").Where("restaurant_id = ?", idResto).Find(&dataComment)
	if result.Error != nil {
		return []comments.Core{}, result.Error
	}
	return toCoreList(dataComment), nil
}

func (repo *mysqlCommentRepository) SelectRatingByIdResto(idResto int) (rating float64, err error) {
	dataComment := Comments_Ratings{}
	result := repo.DB.Select("ROUND(AVG(rating), 2) as rating").Where("restaurant_id = ?", idResto).First(&dataComment)
	if result.Error != nil {
		return 0.0, result.Error
	}
	return dataComment.Rating, nil
}
