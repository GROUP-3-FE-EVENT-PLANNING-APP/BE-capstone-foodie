package data

import (
	"capstone/group3/features/admins"
	"fmt"

	"gorm.io/gorm"
)

type mysqlAdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(conn *gorm.DB) admins.Data {
	return &mysqlAdminRepository{
		db: conn,
	}
}

func (repo *mysqlAdminRepository) AllUserData(limit, offset, idUser int) (response []admins.Core, err error) {
	var dataUsers []User
	var detaiUser User
	// cek id login(role == admin)
	resultCheck := repo.db.Table("users").Where("id = ?", idUser).First(&detaiUser)

	if resultCheck.Error != nil {
		return []admins.Core{}, resultCheck.Error
	}

	if detaiUser.Role != "admin" {
		return []admins.Core{}, fmt.Errorf("not admin")
	}

	result := repo.db.Table("users").Where("role = ?", "user").Order("id desc").Limit(limit).Offset(offset).Find(&dataUsers)

	if result.Error != nil {
		return []admins.Core{}, result.Error
	}

	return toCoreList(dataUsers), nil
}

func (repo *mysqlAdminRepository) AllRestoData(limit, offset, idUser int) (response []admins.CoreRestaurant, err error) {
	var dataResto []Restaurant
	var detaiUser User
	// cek id login(role == admin)
	resultCheck := repo.db.Table("users").Where("id = ?", idUser).First(&detaiUser)

	if resultCheck.Error != nil {
		return []admins.CoreRestaurant{}, resultCheck.Error
	}

	if detaiUser.Role != "admin" {
		return []admins.CoreRestaurant{}, fmt.Errorf("not admin")
	}

	result := repo.db.Table("restaurants").Order("status asc").Limit(limit).Offset(offset).Find(&dataResto)

	if result.Error != nil {
		return []admins.CoreRestaurant{}, result.Error
	}

	return toCoreRestaurantList(dataResto), nil
}

func (repo *mysqlAdminRepository) RestoImageData(idResto int) (response string, err error) {
	data := RestoImage{}

	result := repo.db.Where("restaurant_id = ?", idResto).First(&data)

	if result.Error != nil {
		return "", result.Error
	}

	return data.RestoImageUrl, nil

}

func (repo *mysqlAdminRepository) RatingData(idResto int) (response float64, err error) {
	dataComment := Comments_Ratings{}

	result := repo.db.Select("ROUND(AVG(rating), 2) as rating").Where("restaurant_id = ?", idResto).First(&dataComment)

	if result.Error != nil {
		return 0.0, result.Error
	}

	return dataComment.Rating, nil

}

func (repo *mysqlAdminRepository) RestoImagesData(idResto int) (response []string, err error) {
	data := []RestoImage{}

	result := repo.db.Where("restaurant_id = ?", idResto).Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	var imgs []string

	for i := 0; i < len(data); i++ {
		imgs = append(imgs, data[i].RestoImageUrl)
	}

	return imgs, nil

}

func (repo *mysqlAdminRepository) FacilitiesData(idResto int) (response []string, err error) {
	data := []Facility{}

	result := repo.db.Where("restaurant_id = ?", idResto).Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	var facilities []string

	for i := 0; i < len(data); i++ {
		facilities = append(facilities, data[i].Facility)
	}

	return facilities, nil

}

func (repo *mysqlAdminRepository) DetailRestoData(id, idUser int) (response admins.CoreDetailRestaurant, err error) {
	var dataResto RestaurantDetail

	var detaiUser User
	// cek id login(role == admin)
	resultCheck := repo.db.Table("users").Where("id = ?", idUser).First(&detaiUser)

	if resultCheck.Error != nil {
		return admins.CoreDetailRestaurant{}, resultCheck.Error
	}

	if detaiUser.Role != "admin" {
		return admins.CoreDetailRestaurant{}, fmt.Errorf("not admin")
	}

	result := repo.db.Table("restaurants").Preload("User").First(&dataResto, "id = ?", id)

	if result.RowsAffected != 1 {
		return admins.CoreDetailRestaurant{}, fmt.Errorf("resto not found")
	}

	if result.Error != nil {
		return admins.CoreDetailRestaurant{}, result.Error
	}

	dataResto.OwnerName = dataResto.User.Name
	dataResto.Email = dataResto.User.Email
	dataResto.Handphone = dataResto.User.Handphone

	return dataResto.toCoreDetail(), nil
}

func (repo *mysqlAdminRepository) VerifRestoData(id, idUser int) (response int, err error) {
	var detaiUser User
	// cek id login(role == admin)
	resultCheck := repo.db.Table("users").Where("id = ?", idUser).First(&detaiUser)

	if resultCheck.Error != nil {
		return 0, resultCheck.Error
	}

	if detaiUser.Role != "admin" {
		return 0, fmt.Errorf("not admin")
	}

	result := repo.db.Table("restaurants").Where("id = ?", id).Updates(Restaurant{Status: "verified"})

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("resto not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}
