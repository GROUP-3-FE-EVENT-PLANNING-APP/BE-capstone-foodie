package data

import (
	"capstone/group3/config"
	"capstone/group3/features/favourites"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFavDB(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{}, &RestoImage{}, &Comments_Ratings{}, &Restaurant{}, &User{})

	db.AutoMigrate(&User{}, &Restaurant{}, &Favourite{}, &RestoImage{}, &Comments_Ratings{})

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

}

func TestAddFavDBFailed(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{}, &RestoImage{}, &Comments_Ratings{}, &Restaurant{}, &User{})

	db.AutoMigrate(&User{}, &Restaurant{}, &Favourite{}, &RestoImage{}, &Comments_Ratings{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})
	db.Create(&RestoImage{RestoImageUrl: "foto"})
	db.Create(&Comments_Ratings{UserID: 1,
		RestaurantID: 1,
		Comment:      "bagus",
		Rating:       5,
	})
	db.Create(&Favourite{UserID: 1, RestaurantID: 1})

	repo := NewFavouriteRepository(db)

	t.Run("Test Add Fav Failed -1", func(t *testing.T) {

		row, err := repo.AddFavDB(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, -1, row)
	})

	t.Run("Test Add Fav Failed 0", func(t *testing.T) {

		row, err := repo.AddFavDB(2, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestDeleteFavDB(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{}, &RestoImage{}, &Comments_Ratings{}, &Restaurant{}, &User{})

	db.AutoMigrate(&User{}, &Restaurant{}, &Favourite{}, &RestoImage{}, &Comments_Ratings{})

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

	db.Migrator().DropTable(&Favourite{})

	t.Run("Test Delete Fav Failed", func(t *testing.T) {

		row, err := repo.DeleteFavDB(2, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})

}

func TestRatingData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{}, &RestoImage{}, &Comments_Ratings{}, &Restaurant{}, &User{})

	db.AutoMigrate(&User{}, &Restaurant{}, &RestoImage{}, &Comments_Ratings{}, &Favourite{})

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

	t.Run("Test Get Rating Data Failed", func(t *testing.T) {

		result, err := repo.RatingData(3)
		assert.Nil(t, err)
		assert.Equal(t, 0.0, result)
	})

}

func TestRatingDataFailed(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{}, &RestoImage{}, &Comments_Ratings{}, &Restaurant{}, &User{})

	db.AutoMigrate(&User{}, &Restaurant{}, &RestoImage{}, &Favourite{})

	db.Create(&User{Name: "dwi"})
	db.Create(&Restaurant{RestoName: "depot puas"})
	db.Create(&RestoImage{RestoImageUrl: "foto"})
	db.Create(&Favourite{UserID: 1, RestaurantID: 1})

	repo := NewFavouriteRepository(db)

	t.Run("Test Get Rating Data Failed", func(t *testing.T) {

		result, err := repo.RatingData(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0.0, result)
	})

}

func TestRestoImageData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{}, &RestoImage{}, &Comments_Ratings{}, &Restaurant{}, &User{})

	db.AutoMigrate(&User{}, &Restaurant{}, &Favourite{}, &RestoImage{}, &Comments_Ratings{})

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

	db.Migrator().DropTable(&RestoImage{})

	t.Run("Test Restaurant Image Data Failed", func(t *testing.T) {

		result, err := repo.RestoImageData(1)
		assert.NotNil(t, err)
		assert.Equal(t, "", result)
	})

}

func TestAllRestoData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Favourite{}, &RestoImage{}, &Comments_Ratings{}, &Restaurant{}, &User{})

	db.AutoMigrate(&User{}, &Restaurant{}, &Favourite{}, &RestoImage{}, &Comments_Ratings{})

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

	db.Migrator().DropTable(&Favourite{})

	t.Run("Test Get Rating Data Failed", func(t *testing.T) {

		result, err := repo.AllRestoData(1, 0, 1)
		assert.NotNil(t, err)
		assert.Equal(t, []favourites.RestoCore{}, result)
	})

}
