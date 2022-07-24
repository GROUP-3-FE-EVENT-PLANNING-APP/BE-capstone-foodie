package data

import (
	"capstone/group3/config"
	"capstone/group3/features/comments"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertComment(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Comments_Ratings{}, &Restaurant{}, &User{})
	db.AutoMigrate(&User{}, &Restaurant{}, &Comments_Ratings{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})

	repo := NewCommentRepository(db)

	t.Run("Test Post Comment", func(t *testing.T) {
		mockUser := comments.Core{User: comments.User{ID: 1},
			Restaurant: comments.Restaurant{ID: 1},
			Comment:    "bagus",
			Rating:     5,
		}
		row, err := repo.InsertComment(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})

	t.Run("Test Post Comment Failed Resto ID not found", func(t *testing.T) {
		mockUser := comments.Core{User: comments.User{ID: 1},
			Restaurant: comments.Restaurant{ID: 2},
			Comment:    "bagus",
			Rating:     5,
		}
		row, err := repo.InsertComment(mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})

}

func TestSelectCommentByIdResto(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Comments_Ratings{}, &Restaurant{}, &User{})
	db.AutoMigrate(&User{}, &Restaurant{}, &Comments_Ratings{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 1,
		Comment:      "bagus",
		Rating:       5,
	})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 1,
		Comment:      "keren",
		Rating:       4,
	})

	repo := NewCommentRepository(db)

	t.Run("Test Get Comments", func(t *testing.T) {

		result, err := repo.SelectCommentByIdResto(1, 2, 0)
		assert.Nil(t, err)
		assert.Equal(t, "keren", result[0].Comment)
	})

	db.Migrator().DropTable(&Comments_Ratings{})

	t.Run("Test Create User", func(t *testing.T) {

		resultfailed, errfailed := repo.SelectCommentByIdResto(2, 1, 0)
		assert.NotNil(t, errfailed)
		assert.Equal(t, []comments.Core{}, resultfailed)
	})

}

func TestSelectRatingByIdResto(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Comments_Ratings{}, &Restaurant{}, &User{})
	db.AutoMigrate(&User{}, &Restaurant{}, &Comments_Ratings{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 1,
		Comment:      "bagus",
		Rating:       5,
	})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 1,
		Comment:      "keren",
		Rating:       4,
	})

	repo := NewCommentRepository(db)

	t.Run("Test Get Rating", func(t *testing.T) {

		result, err := repo.SelectRatingByIdResto(1)
		assert.Nil(t, err)
		assert.Equal(t, 4.5, result)
	})

	db.Migrator().DropTable(&Comments_Ratings{})

	t.Run("Test Get Rating Failed", func(t *testing.T) {

		result, err := repo.SelectRatingByIdResto(2)
		assert.NotNil(t, err)
		assert.Equal(t, 0.0, result)
	})

}
