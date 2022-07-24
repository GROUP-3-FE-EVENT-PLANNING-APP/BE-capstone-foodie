package data

import (
	"capstone/group3/config"
	"capstone/group3/features/users"
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	repo := NewUserRepository(db)

	t.Run("Test Create User", func(t *testing.T) {
		mockUser := users.Core{Name: "dwi",
			Email:    "dwi@gmail.com",
			Password: "qwerty",
		}
		row, err := repo.InsertData(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})

	t.Run("Test Create User", func(t *testing.T) {
		mockUser := users.Core{Name: "dwi",
			Email:     "dwi@gmail.com",
			Password:  "qwerty",
			Handphone: "9987",
		}
		row, err := repo.InsertData(mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})

}

func TestLoginUserDB(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	repo := NewUserRepository(db)
	data := users.Core{
		Name:     "dwi",
		Email:    "dwiatmokop@gmail.com",
		Password: "qwerty",
	}
	_, err := repo.InsertData(data)
	if err != nil {
		fmt.Errorf("error inserting")
	}

	t.Run("Test Login User", func(t *testing.T) {
		mockUser := users.AuthRequestData{
			Email:    "dwiatmokop@gmail.com",
			Password: "qwerty",
		}
		row, err := repo.LoginUserDB(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, row["id"])
	})

	t.Run("Test Login User Failed", func(t *testing.T) {
		mockUser := users.AuthRequestData{
			Email:    "dwiatmoko@gmail.com",
			Password: "qwerty",
		}
		result, err := repo.LoginUserDB(mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, map[string]interface{}(nil), result)
	})

}

func TestUpdateDataDB(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	repo := NewUserRepository(db)
	data := users.Core{
		Name:     "dwi",
		Email:    "dwiatmokop@gmail.com",
		Password: "qwerty",
	}
	_, err := repo.InsertData(data)
	if err != nil {
		fmt.Errorf("error inserting")
	}

	t.Run("Test Update User", func(t *testing.T) {
		mockUser := map[string]interface{}{}
		mockUser["name"] = "dwiatmoko"
		row, err := repo.UpdateDataDB(mockUser, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})

	t.Run("Test Update User Failed", func(t *testing.T) {
		mockUser := map[string]interface{}{}
		mockUser["name"] = "dwiatmoko"
		row, err := repo.UpdateDataDB(mockUser, 2)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})

}

func TestSelectDataByMe(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	repo := NewUserRepository(db)
	data := users.Core{
		Name:     "dwi",
		Email:    "dwiatmokop@gmail.com",
		Password: "qwerty",
	}
	_, err := repo.InsertData(data)
	if err != nil {
		fmt.Errorf("error inserting")
	}

	t.Run("Test Get My Profile", func(t *testing.T) {
		result, err := repo.SelectDataByMe(1)
		assert.Nil(t, err)
		assert.Equal(t, "dwi", result.Name)
	})

	// t.Run("Test Get My Profile", func(t *testing.T) {
	// 	row, err := repo.SelectDataByMe(3)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, users.Core{}, row)
	// })

}

func TestDeleteDataByIdDB(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	repo := NewUserRepository(db)
	data := users.Core{
		Name:     "dwi",
		Email:    "dwiatmokop@gmail.com",
		Password: "qwerty",
	}
	_, err := repo.InsertData(data)
	if err != nil {
		fmt.Errorf("error inserting")
	}

	t.Run("Test Delete Account", func(t *testing.T) {
		row, err := repo.DeleteDataByIdDB(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})

	t.Run("Test Delete Account Failed", func(t *testing.T) {
		row, err := repo.DeleteDataByIdDB(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})

}
