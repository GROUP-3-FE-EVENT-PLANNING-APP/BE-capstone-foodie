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
