package data

import (
	"capstone/group3/config"
	"capstone/group3/features/users"

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
}
