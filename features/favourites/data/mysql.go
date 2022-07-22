package data

import (
	"capstone/group3/features/favourites"

	"gorm.io/gorm"
)

type mysqlFavouriteRepository struct {
	DB *gorm.DB
}

func NewFavouriteRepository(db *gorm.DB) favourites.Data {
	return &mysqlFavouriteRepository{
		DB: db,
	}
}

func (repo *mysqlFavouriteRepository) AddFavDB(idResto, idFromToken int) (row int, err error) {
	dataFav := Favourite{}
	dataFav.UserID = idFromToken
	dataFav.RestaurantID = idResto
	checkDB := Favourite{}
	resultCheck := repo.DB.Where("restaurant_id=? and user_id=?", idResto, idFromToken).First(&checkDB)
	if resultCheck.Error == nil {
		return -1, resultCheck.Error
	} else {
		resultAdd := repo.DB.Create(&dataFav)
		if resultAdd.Error != nil {
			return 0, resultAdd.Error
		}
	}

	return 1, nil
}

func (repo *mysqlFavouriteRepository) RatingData(idResto int) (response float64, err error) {
	dataComment := Comments_Ratings{}

	result := repo.DB.Select("ROUND(AVG(rating), 2) as rating").Where("restaurant_id = ?", idResto).First(&dataComment)

	if result.Error != nil {
		return 0.0, result.Error
	}

	return dataComment.Rating, nil

}

func (repo *mysqlFavouriteRepository) RestoImageData(idResto int) (response string, err error) {
	data := RestoImage{}

	result := repo.DB.Where("restaurant_id = ?", idResto).First(&data)

	if result.Error != nil {
		return "", result.Error
	}

	return data.RestoImageUrl, nil

}

func (repo *mysqlFavouriteRepository) DeleteFavDB(idResto, idFromToken int) (row int, err error) {
	dataFav := Favourite{}
	dataFav.UserID = idFromToken
	dataFav.RestaurantID = idResto
	resultDel := repo.DB.Unscoped().Where("restaurant_id=? and user_id=?", idResto, idFromToken).Delete(&dataFav)
	if resultDel.Error != nil {
		return 0, resultDel.Error
	}
	return int(resultDel.RowsAffected), nil
}

func (repo *mysqlFavouriteRepository) AllRestoData(limitint, offsetint, idFromToken int) (response []favourites.RestoCore, err error) {
	var dataFav []Favourite

	result := repo.DB.Preload("Restaurant").Where("favourites.user_id = ?", idFromToken).Limit(limitint).Offset(offsetint).Order("created_at DESC").Find(&dataFav)

	if result.Error != nil {
		return []favourites.RestoCore{}, result.Error
	}

	return toCoreList(dataFav), nil
}
