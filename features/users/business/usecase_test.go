package business

import (
	"capstone/group3/features/users"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// success

type mockUserData struct{}

func (mock mockUserData) InsertData(input users.Core) (row int, err error) {
	return 1, nil
}

func (mock mockUserData) LoginUserDB(authData users.AuthRequestData) (data map[string]interface{}, err error) {
	dataToken := make(map[string]interface{})
	dataToken["token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJhdmF0YXJVcmwiOiJodHRwczovL2d1ZGFuZ2dyb3VwMy5zMy5hbWF6b25hd3MuY29tLzIwMjItMDctMDYlMjAwNSUzQTU5JTNBMTAuanBnIiwiZXhwIjoxNjU3NzU0NjM4LCJ1c2VySWQiOjF9.bYlOfoHZS2JupZoEJgmVBelnzgPDX_bWHtV46IVNIAE"
	dataToken["name"] = "dwi2"
	dataToken["avatar_url"] = "https://gudanggroup3.s3.amazonaws.com/2022-07-06%2005%3A59%3A10.jpg"
	dataToken["email"] = "dwi2@mail.com"
	dataToken["handphone"] = "085608560856"
	return dataToken, nil
}

func (mockUserData) UpdateDataDB(data map[string]interface{}, idUser int) (row int, err error) {
	return 1, nil
}

func (mock mockUserData) SelectDataByMe(idFromToken int) (data users.Core, err error) {
	return users.Core{
		ID:        1,
		Name:      "dwi2",
		Email:     "dwi@mail.com",
		AvatarUrl: "https://gudanggroup3.s3.amazonaws.com/2022-07-06%2005%3A59%3A10.jpg",
		Handphone: "085608560856",
	}, nil
}

func (mock mockUserData) DeleteDataByIdDB(idFromToken int) (row int, err error) {
	return 1, nil
}

// failed
type mockUserDataFailed struct{}

func (mock mockUserDataFailed) InsertData(input users.Core) (row int, err error) {
	return 0, fmt.Errorf("failed to insert data ")
}

func (mock mockUserDataFailed) LoginUserDB(authData users.AuthRequestData) (data map[string]interface{}, err error) {
	return map[string]interface{}{}, fmt.Errorf("failed to login")
}

func (mock mockUserDataFailed) UpdateDataDB(data map[string]interface{}, idUser int) (row int, err error) {
	return 0, fmt.Errorf("failed to update data ")
}

func (mock mockUserDataFailed) SelectDataByMe(idFromToken int) (data users.Core, err error) {
	return users.Core{}, fmt.Errorf("failed to get user data ")
}

func (mock mockUserDataFailed) DeleteDataByIdDB(idFromToken int) (row int, err error) {
	return 0, fmt.Errorf("failed to delete data ")
}

func TestCreateData(t *testing.T) {
	t.Run("Test Create Data Success", func(t *testing.T) {
		input := users.Core{
			Name:      "dwi",
			Email:     "dwi@mail.com",
			Password:  "qwerty",
			Handphone: "085608560856",
		}
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.CreateData(input)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})

	t.Run("Test Create Data Failed", func(t *testing.T) {
		input := users.Core{
			Name:      "dwi",
			Email:     "dwi@mail.com",
			Password:  "qwerty",
			Handphone: "085608560856",
		}
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.CreateData(input)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestLoginUserDB(t *testing.T) {
	t.Run("Test Login User DB Success", func(t *testing.T) {
		authData := users.AuthRequestData{
			Email:    "dwi@mail.com",
			Password: "qwerty",
		}
		userBusiness := NewUserBusiness(mockUserData{})
		dataToken, err := userBusiness.LoginUser(authData)
		assert.Nil(t, err)
		assert.Equal(t, "dwi2", dataToken["name"])
	})

	t.Run("Test Login User DB Failed", func(t *testing.T) {
		authData := users.AuthRequestData{
			Email:    "dwi@mail.com",
			Password: "qwerty",
		}
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		data, err := userBusiness.LoginUser(authData)
		assert.NotNil(t, err)
		assert.Equal(t, "dwi2", data["name"])
	})
}
