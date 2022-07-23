package business

import (
	"capstone/group3/features/favourites"
	"capstone/group3/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddFavDB(t *testing.T) {
	repo := new(mocks.FavouriteData)

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("AddFavDB", mock.Anything, mock.Anything).Return(1, nil).Once()
		srv := NewFavouriteBusiness(repo)

		res, err := srv.AddFav(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

}

func TestDeleteFavDB(t *testing.T) {
	repo := new(mocks.FavouriteData)

	t.Run("Success Delete", func(t *testing.T) {
		repo.On("DeleteFavDB", mock.Anything, mock.Anything).Return(1, nil).Once()
		srv := NewFavouriteBusiness(repo)

		res, err := srv.DeleteFavourite(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

}

func TestAllRestoData(t *testing.T) {
	repo := new(mocks.FavouriteData)

	dataRestoCore := []favourites.RestoCore{
		{
			ID:          1,
			Rating:      4.5,
			RestoImages: "foto",
		},
	}

	t.Run("Success Get My Fav", func(t *testing.T) {
		repo.On("RatingData", mock.Anything, mock.Anything, mock.Anything).Return(4.5, nil).Once()

		repo.On("RestoImageData", mock.Anything, mock.Anything, mock.Anything).Return("foto", nil).Once()

		repo.On("AllRestoData", mock.Anything, mock.Anything, mock.Anything).Return(dataRestoCore, nil).Once()
		srv := NewFavouriteBusiness(repo)

		res, err := srv.GetMyFav(1, 0, 1)
		assert.NoError(t, err)
		assert.Equal(t, 4.5, res[0].Rating)
		repo.AssertExpectations(t)
	})
}
