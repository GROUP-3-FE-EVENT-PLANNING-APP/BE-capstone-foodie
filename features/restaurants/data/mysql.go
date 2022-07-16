package data

import (
	"capstone/group3/features/restaurants"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type mysqlRestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(conn *gorm.DB) restaurants.Data {
	return &mysqlRestaurantRepository{
		db: conn,
	}
}

func (repo *mysqlRestaurantRepository) InsertRestoData(input restaurants.Core) (response int, err error) {
	restaurant := fromCore(input)
	result := repo.db.Create(&restaurant)

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}

	facilitiesArray := strings.Split(input.Facility, ",")

	for _, v := range facilitiesArray {
		var facility Facility

		facility.RestaurantID = restaurant.ID
		facility.Facility = strings.TrimSpace(v)

		result_ := repo.db.Create(&facility)

		if result_.Error != nil {
			return 0, result.Error
		}

	}

	return int(result.RowsAffected), err
}

func (repo *mysqlRestaurantRepository) UpdateRestoData(editData restaurants.Core, idUser int) (response int, err error) {

	resto := fromCore(editData)

	result := repo.db.Model(Restaurant{}).Where("user_id = ?", idUser).Updates(&resto).First(&resto)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("restaurant not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	var facility Facility

	// delete facility lama
	_ = repo.db.Unscoped().Where("restaurant_id = ?", resto.ID).Delete(&facility)
	// resultDelFacility := repo.db.Unscoped().Delete(&facility)

	// split data facility
	facilitiesArray := strings.Split(editData.Facility, ",")

	for _, v := range facilitiesArray {
		var facility Facility

		facility.RestaurantID = resto.ID
		facility.Facility = strings.TrimSpace(v)

		result_ := repo.db.Create(&facility)

		if result_.Error != nil {
			return 0, result.Error
		}

	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlRestaurantRepository) DetailImageRestoData(id int) (imageMenu, imageFile string, err error) {
	var dataResto Restaurant

	result := repo.db.Preload("User").First(&dataResto, "user_id = ?", id)

	if result.RowsAffected != 1 {
		return "", "", fmt.Errorf("restaurant not found")
	}

	if result.Error != nil {
		return "", "", result.Error
	}

	return dataResto.MenuImageUrl, dataResto.FileImageUrl, nil
}

func (repo *mysqlRestaurantRepository) DeleteRestoData(idUser int) (row int, err error) {
	var dataResto Restaurant

	result := repo.db.Unscoped().Where("user_id = ?", idUser).Delete(&dataResto)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("restaurant not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlRestaurantRepository) UploadImageRestoData(input restaurants.RestoImage) (response int, err error) {
	dataImage := fromCoreRestoImage(input)

	var dataResto Restaurant

	searchResto := repo.db.Table("restaurants").Where("user_id = ?", input.UserID).First(&dataResto)

	if searchResto.Error != nil || searchResto.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to upload data")
	}

	dataImage.RestaurantID = dataResto.ID

	result := repo.db.Create(&dataImage)

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to upload data")
	}

	return int(result.RowsAffected), err
}

func (repo *mysqlRestaurantRepository) AllRestoData(limit, offset int) (response []restaurants.CoreList, err error) {
	var dataResto []Restaurant

	result := repo.db.Preload("RestoImages").Model(&Restaurant{}).Select("id, category, resto_name, location, table_quota").Where("status = ?", "verification").Order("id desc").Limit(limit).Offset(offset).Find(&dataResto)

	if result.Error != nil {
		return []restaurants.CoreList{}, result.Error
	}

	// fmt.Println(dataResto[0].RestoImages[0].RestoImageUrl)

	return toCoreList(dataResto), nil
}

func (repo *mysqlRestaurantRepository) RatingData(idResto int) (response float64, err error) {
	dataComment := Comments_Ratings{}

	result := repo.db.Select("ROUND(AVG(rating), 2) as rating").Where("restaurant_id = ?", idResto).First(&dataComment)

	if result.Error != nil {
		return 0.0, result.Error
	}

	return dataComment.Rating, nil

}

func (repo *mysqlRestaurantRepository) RestoImageData(idResto int) (response string, err error) {
	data := RestoImage{}

	result := repo.db.Where("restaurant_id = ?", idResto).First(&data)

	if result.Error != nil {
		return "", result.Error
	}

	return data.RestoImageUrl, nil

}
