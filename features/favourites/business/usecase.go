package business

import (
	"capstone/group3/features/favourites"
)

type favouriteUseCase struct {
	favouriteData favourites.Data
}

func NewFavouriteBusiness(dataFav favourites.Data) favourites.Business {
	return &favouriteUseCase{
		favouriteData: dataFav,
	}
}

func (uc *favouriteUseCase) AddFav(idResto, idFromToken int) (row int, err error) {
	row, err = uc.favouriteData.AddFavDB(idResto, idFromToken)
	return row, err
}

func (uc *favouriteUseCase) DeleteFavourite(idResto, idFromToken int) (row int, err error) {
	row, err = uc.favouriteData.DeleteFavDB(idResto, idFromToken)
	return row, err
}

func (uc *favouriteUseCase) GetMyFav(limitint, offsetint, idFromToken int) (response []favourites.RestoCore, err error) {
	response, err = uc.favouriteData.AllRestoData(limitint, offsetint, idFromToken)

	if err == nil {
		for i := 0; i < len(response); i++ {

			// get rating
			rating, _ := uc.favouriteData.RatingData(response[i].ID)
			response[i].Rating = rating

			// get resto image url
			restoImg, _ := uc.favouriteData.RestoImageData(response[i].ID)

			response[i].RestoImages = restoImg
		}
	}

	return response, err
}
