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
