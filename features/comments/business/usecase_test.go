package business

import (
	"capstone/group3/features/comments"
	"fmt"

	"capstone/group3/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertComment(t *testing.T) {
	repo := new(mocks.Data)
	insertData := comments.Core{ID: 1,
		User:       comments.User{ID: 1},
		Restaurant: comments.Restaurant{ID: 1},
		Comment:    "alta",
		Rating:     5,
	}
	// returnData := users.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertComment", mock.Anything).Return(1, nil).Once()
		srv := NewCommentBusiness(repo)

		res, err := srv.CreateComment(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		repo.On("InsertComment", mock.Anything).Return(0, fmt.Errorf("failed to insert comment")).Once()
		srv := NewCommentBusiness(repo)

		res, err := srv.CreateComment(insertData)
		assert.EqualError(t, err, "failed to insert comment")
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		insertDataMissingValue := comments.Core{ID: 1,
			User:       comments.User{ID: 1},
			Restaurant: comments.Restaurant{ID: 1},
			Rating:     5,
		}
		// repo.On("InsertComment", mock.Anything).Return(-1, errors.New("please make sure all fields are filled in correctly")).Once()
		srv := NewCommentBusiness(repo)

		res, err := srv.CreateComment(insertDataMissingValue)
		assert.EqualError(t, err, "please make sure all fields are filled in correctly")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		insertDataLowerValue := comments.Core{ID: 1,
			User:       comments.User{ID: 1},
			Restaurant: comments.Restaurant{ID: 1},
			Comment:    "bagus",
			Rating:     0,
		}

		srv := NewCommentBusiness(repo)

		res, err := srv.CreateComment(insertDataLowerValue)
		assert.EqualError(t, err, "rating cant be lower than 1")
		assert.Equal(t, -2, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		insertDataHigherValue := comments.Core{ID: 1,
			User:       comments.User{ID: 1},
			Restaurant: comments.Restaurant{ID: 1},
			Comment:    "bagus",
			Rating:     6,
		}

		srv := NewCommentBusiness(repo)

		res, err := srv.CreateComment(insertDataHigherValue)
		assert.EqualError(t, err, "rating cant be greater than 5")
		assert.Equal(t, -3, res)
		repo.AssertExpectations(t)
	})

	// t.Run("Error insert when incomplete data", func(t *testing.T) {
	// 	/*
	// 	   dont need to write repo.On because this test case dont need to call data layer. just handle on business layer.
	// 	*/
	// 	// repo.On("InsertData", mock.Anything).Return(-1, errors.New("all input data must be filled")).Once()
	// 	srv := NewUserBusiness(repo)

	// 	_, err := srv.InsertData(users.Core{})
	// 	assert.EqualError(t, err, "all input data must be filled")
	// 	repo.AssertExpectations(t)
	// })
}

func TestSelectCommentByIdResto(t *testing.T) {
	repo := new(mocks.Data)
	result := []comments.Core{
		{
			ID:      1,
			Comment: "bagus",
			Rating:  5,
		},
	}

	t.Run("Success Select Comment", func(t *testing.T) {
		repo.On("SelectCommentByIdResto", mock.Anything, mock.Anything, mock.Anything).Return(result, nil).Once()
		srv := NewCommentBusiness(repo)

		res, err := srv.GetCommentByIdResto(1, 1, 0)
		assert.NoError(t, err)
		assert.Equal(t, result[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})
}

func TestSelectRatingByIdResto(t *testing.T) {
	repo := new(mocks.Data)
	result := 4.6

	t.Run("Success Select Rating", func(t *testing.T) {
		repo.On("SelectRatingByIdResto", mock.Anything).Return(result, nil).Once()
		srv := NewCommentBusiness(repo)

		res, err := srv.GetRatingByIdResto(1)
		assert.NoError(t, err)
		assert.Equal(t, result, res)
		repo.AssertExpectations(t)
	})
}

// success

// type mockCommentData struct{}

// func (mock mockCommentData) InsertComment(input comments.Core) (row int, err error) {
// 	return 1, nil
// }

// func (mock mockCommentData) SelectCommentByIdResto(idResto, limitint, offsetint int) (data []comments.Core, err error) {
// 	return []comments.Core{
// 		{
// 			ID:      1,
// 			Comment: "bagus",
// 			Rating:  5,
// 		},
// 	}, nil
// }

// func (mock mockCommentData) SelectRatingByIdResto(idResto int) (rating float64, err error) {
// 	return 4.5, nil
// }

// // failed
// type mockCommentDataFailed struct{}

// func (mock mockCommentDataFailed) InsertComment(input comments.Core) (row int, err error) {
// 	return 0, fmt.Errorf("failed to insert comment")
// }

// func (mock mockCommentDataFailed) SelectCommentByIdResto(idResto, limitint, offsetint int) (data []comments.Core, err error) {
// 	return []comments.Core{}, fmt.Errorf("failed to get comment data ")
// }

// func (mock mockCommentDataFailed) SelectRatingByIdResto(idResto int) (rating float64, err error) {
// 	return 9, fmt.Errorf("failed to get rating")
// }

// func TestInsertComment(t *testing.T) {
// 	t.Run("Test InsertComment Success", func(t *testing.T) {
// 		input := comments.Core{
// 			Comment: "bagus",
// 			Rating:  5,
// 		}
// 		commentBusiness := NewCommentBusiness(mockCommentData{})
// 		result, err := commentBusiness.CreateComment(input)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, result)
// 	})

// 	t.Run("Test InsertComment Failed", func(t *testing.T) {
// 		input := comments.Core{
// 			Comment: "bagus",
// 			Rating:  5,
// 		}
// 		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
// 		result, err := commentBusiness.CreateComment(input)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, 0, result)
// 	})

// 	t.Run("Test InsertComment Failed", func(t *testing.T) {
// 		input := comments.Core{
// 			Rating: 5,
// 		}
// 		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
// 		result, err := commentBusiness.CreateComment(input)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, -1, result)
// 	})

// 	t.Run("Test InsertComment Failed", func(t *testing.T) {
// 		input := comments.Core{
// 			Comment: "bagus",
// 			Rating:  0,
// 		}
// 		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
// 		result, err := commentBusiness.CreateComment(input)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, -2, result)
// 	})

// 	t.Run("Test InsertComment Failed", func(t *testing.T) {
// 		input := comments.Core{
// 			Comment: "bagus",
// 			Rating:  6,
// 		}
// 		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
// 		result, err := commentBusiness.CreateComment(input)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, -3, result)
// 	})
// }

// func TestSelectCommentByIdResto(t *testing.T) {
// 	t.Run("Test SelectCommentById Success", func(t *testing.T) {
// 		commentBusiness := NewCommentBusiness(mockCommentData{})
// 		result, err := commentBusiness.GetCommentByIdResto(1, 1, 0)
// 		assert.Nil(t, err)
// 		assert.Equal(t, result[0].Comment, result[0].Comment)
// 	})

// 	t.Run("Test SelectCommentById Failed", func(t *testing.T) {
// 		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
// 		result, err := commentBusiness.GetCommentByIdResto(1, 1, 0)
// 		assert.Nil(t, err)
// 		assert.Equal(t, []comments.Core{}, result)
// 	})
// }
