package data

import (
	"capstone/group3/features/restaurants"
	"fmt"

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

	return int(result.RowsAffected), err
}
