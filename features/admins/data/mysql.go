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
