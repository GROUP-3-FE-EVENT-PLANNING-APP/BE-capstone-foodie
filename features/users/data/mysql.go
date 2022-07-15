package data

import (
	"capstone/group3/features/middlewares"
	"capstone/group3/features/users"
	"errors"
	"fmt"

	_bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.Data {
	return &mysqlUserRepository{
		DB: db,
	}
}

func (repo *mysqlUserRepository) InsertData(input users.Core) (row int, err error) {
	user := FromCore(input)
	passwordHashed, errorHash := _bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorHash != nil {
		fmt.Println("Error hash", errorHash.Error())
	}
	user.Password = string(passwordHashed)
	resultcreate := repo.DB.Create(&user)
	if resultcreate.Error != nil {
		return 0, resultcreate.Error
	}
	if resultcreate.RowsAffected != 1 {
		return 0, errors.New("failed to insert data, your email is already registered")
	}
	return int(resultcreate.RowsAffected), nil
}

func (repo *mysqlUserRepository) LoginUserDB(authData users.AuthRequestData) (data map[string]interface{}, err error) {
	userData := User{}
	result := repo.DB.Where("email = ?", authData.Email).First(&userData)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected != 1 {
		return nil, errors.New("failed to login")
	}

	errCrypt := _bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(authData.Password))
	if errCrypt != nil {
		return nil, errors.New("password incorrect")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID), userData.AvatarUrl, userData.Role, userData.Handphone, userData.Email)
	if errToken != nil {
		return nil, errToken
	}
	var dataToken = map[string]interface{}{}
	dataToken["id"] = int(userData.ID)
	dataToken["token"] = token
	dataToken["name"] = userData.Name
	dataToken["avatarUrl"] = userData.AvatarUrl
	dataToken["role"] = userData.Role
	dataToken["handphone"] = userData.Handphone
	dataToken["email"] = userData.Email
	return dataToken, nil
}

func (repo *mysqlUserRepository) UpdateDataDB(data map[string]interface{}, idFromToken int) (row int, err error) {
	result := repo.DB.Model(&User{}).Where("id = ?", idFromToken).Updates(data)
	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update data")
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) SelectDataByMe(idFromToken int) (data users.Core, err error) {
	dataUser := User{}
	result := repo.DB.Find(&dataUser, idFromToken)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	return dataUser.toCore(), nil
}

func (repo *mysqlUserRepository) DeleteDataByIdDB(idFromToken int) (row int, err error) {
	result := repo.DB.Delete(&User{}, idFromToken)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to delete data")
	}
	return int(result.RowsAffected), nil

}
