package data

import (
	"capstone/group3/config"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFavDB(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{})
	db.Migrator().DropTable(&RestoImage{})
	db.Migrator().DropTable(&Comments_Ratings{})
	db.Migrator().DropTable(&Restaurant{})
	db.Migrator().DropTable(&User{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Favourite{})
	db.AutoMigrate(&RestoImage{})
	db.AutoMigrate(&Comments_Ratings{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})
	db.Create(&RestoImage{RestoImageUrl: "foto"})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 2,
		Comment:      "bagus",
		Rating:       5,
	})

	repo := NewFavouriteRepository(db)

	t.Run("Test Add Fav", func(t *testing.T) {

		row, err := repo.AddFavDB(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})

	t.Run("Test Add Fav Failed", func(t *testing.T) {

		row, err := repo.AddFavDB(2, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})

}

func TestDeleteFavDB(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{})
	db.Migrator().DropTable(&RestoImage{})
	db.Migrator().DropTable(&Comments_Ratings{})
	db.Migrator().DropTable(&Restaurant{})
	db.Migrator().DropTable(&User{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Favourite{})
	db.AutoMigrate(&RestoImage{})
	db.AutoMigrate(&Comments_Ratings{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})
	db.Create(&RestoImage{RestoImageUrl: "foto"})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 2,
		Comment:      "bagus",
		Rating:       5,
	})
	db.Create(&Favourite{UserID: 1, RestaurantID: 1})

	repo := NewFavouriteRepository(db)

	t.Run("Test Delete Fav", func(t *testing.T) {

		row, err := repo.DeleteFavDB(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})

	// t.Run("Test Delete Fav Failed", func(t *testing.T) {

	// 	row, err := repo.DeleteFavDB(1, 1)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, 0, row)
	// })

}

func TestRatingData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{})
	db.Migrator().DropTable(&RestoImage{})
	db.Migrator().DropTable(&Comments_Ratings{})
	db.Migrator().DropTable(&Restaurant{})
	db.Migrator().DropTable(&User{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Favourite{})
	db.AutoMigrate(&RestoImage{})
	db.AutoMigrate(&Comments_Ratings{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})
	db.Create(&RestoImage{RestoImageUrl: "foto"})
	db.Create(&Favourite{UserID: 1, RestaurantID: 1})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 1,
		Comment:      "bagus",
		Rating:       5,
	})

	repo := NewFavouriteRepository(db)

	t.Run("Test Get Rating Data", func(t *testing.T) {

		result, err := repo.RatingData(1)
		assert.Nil(t, err)
		assert.Equal(t, 5.0, result)
	})

	// t.Run("Test Get Rating Data Failed", func(t *testing.T) {

	// 	result, err := repofailed.RatingData(2)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, 0.0, result)
	// })

}

func TestRestoImageData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{})
	db.Migrator().DropTable(&RestoImage{})
	db.Migrator().DropTable(&Comments_Ratings{})
	db.Migrator().DropTable(&Restaurant{})
	db.Migrator().DropTable(&User{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Favourite{})
	db.AutoMigrate(&RestoImage{})
	db.AutoMigrate(&Comments_Ratings{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})
	db.Create(&RestoImage{RestaurantID: 1, RestoImageUrl: "foto"})
	db.Create(&Favourite{UserID: 1, RestaurantID: 1})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 1,
		Comment:      "bagus",
		Rating:       5,
	})

	repo := NewFavouriteRepository(db)

	t.Run("Test Restaurant Image Data", func(t *testing.T) {

		result, err := repo.RestoImageData(1)
		assert.Nil(t, err)
		assert.Equal(t, "foto", result)
	})

	// t.Run("Test Get Rating Data Failed", func(t *testing.T) {

	// 	result, err := repofailed.RatingData(2)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, 0.0, result)
	// })

}

func TestAllRestoData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{})
	db.Migrator().DropTable(&RestoImage{})
	db.Migrator().DropTable(&Comments_Ratings{})
	db.Migrator().DropTable(&Restaurant{})
	db.Migrator().DropTable(&User{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Favourite{})
	db.AutoMigrate(&RestoImage{})
	db.AutoMigrate(&Comments_Ratings{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})
	db.Create(&Restaurant{RestoName: "depot nikmat"})
	db.Create(&RestoImage{RestaurantID: 1, RestoImageUrl: "foto"})
	db.Create(&RestoImage{RestaurantID: 2, RestoImageUrl: "foto2"})
	db.Create(&Favourite{UserID: 1, RestaurantID: 1})
	db.Create(&Favourite{UserID: 1, RestaurantID: 2})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 1,
		Comment:      "bagus",
		Rating:       5,
	})

	repo := NewFavouriteRepository(db)

	t.Run("Test Data My Fav List", func(t *testing.T) {

		result, err := repo.AllRestoData(1, 0, 1)
		assert.Nil(t, err)
		assert.Equal(t, "depot nikmat", result[0].RestoName)
	})

	// t.Run("Test Get Rating Data Failed", func(t *testing.T) {

	// 	result, err := repofailed.RatingData(2)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, 0.0, result)
	// })

}
