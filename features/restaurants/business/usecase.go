package business

import (
	"capstone/group3/features/restaurants"
	"errors"
	"fmt"
)

type restaurantUseCase struct {
	restaurantData restaurants.Data
}

func NewRestaurantBusiness(rstData restaurants.Data) restaurants.Business {
	return &restaurantUseCase{
		restaurantData: rstData,
	}
}

func (uc *restaurantUseCase) CreateRestoBusiness(newData restaurants.Core) (response int, err error) {
	response, err = uc.restaurantData.InsertRestoData(newData)

	return response, err
}

func (uc *restaurantUseCase) UpdateRestoBusiness(editData restaurants.Core, idUser int) (response int, err error) {
	if editData.RestoName == "" || editData.Location == "" || editData.MenuImageUrl == "" || editData.Category == "" || editData.TableQuota == 0 || editData.BookingFee == 0 || editData.Latitude == "" || editData.Longitude == "" || editData.Facility == "" {
		return 0, errors.New("all input data must be filled")
	}

	response, err = uc.restaurantData.UpdateRestoData(editData, idUser)

	return response, err
}

func (uc *restaurantUseCase) DetailImageRestoBusiness(id int) (imageMenu, imageFile string, err error) {
	imageMenu, imageFile, err = uc.restaurantData.DetailImageRestoData(id)

	return imageMenu, imageFile, err
}

func (uc *restaurantUseCase) DeleteRestoBusiness(idUser int) (row int, err error) {
	row, err = uc.restaurantData.DeleteRestoData(idUser)

	return row, err
}

func (uc *restaurantUseCase) UploadImageRestoBusiness(newData restaurants.RestoImage) (response int, err error) {
	response, err = uc.restaurantData.UploadImageRestoData(newData)

	return response, err
}

func (uc *restaurantUseCase) AllRestoBusiness(limit, offset int) (response []restaurants.CoreList, err error) {
	response, err = uc.restaurantData.AllRestoData(limit, offset)

	if err == nil {
		for i := 0; i < len(response); i++ {

			// get rating
			rating, _ := uc.restaurantData.RatingData(response[i].ID)
			response[i].Rating = rating

			// get resto image url
			restoImg, _ := uc.restaurantData.RestoImageData(response[i].ID)

			response[i].RestoImages = append(response[i].RestoImages, restaurants.RestoImage{RestoImageUrl: restoImg})
		}
	}

	return response, err
}

func (uc *restaurantUseCase) DetailRestoBusiness(id int) (response restaurants.CoreDetail, err error) {
	response, err = uc.restaurantData.DetailRestoData(id)

	// check table quota
	tableQouta, errCheck := uc.restaurantData.CheckTableQuotaData(id)

	if errCheck != nil {
		return restaurants.CoreDetail{}, errCheck
	}

	// check table qouta after
	// tableQoutaAfter, errCheck := uc.restaurantData.CheckTableQuotaData(id)

	response.TableQuota = response.TableQuota - uint(tableQouta)

	if err == nil {

		// get rating
		rating, _ := uc.restaurantData.RatingData(response.ID)
		response.Rating = rating

		// get resto image
		restoImgs, _ := uc.restaurantData.RestoImagesData(response.ID)
		for i := 0; i < len(restoImgs); i++ {
			response.RestoImages = append(response.RestoImages, restaurants.RestoImage{RestoImageUrl: restoImgs[i]})
		}

		// get facility
		facilities, _ := uc.restaurantData.FacilitiesData(response.ID)
		for i := 0; i < len(facilities); i++ {
			response.Facilities = append(response.Facilities, restaurants.Facility{Facility: facilities[i]})
		}

		// get comment
		comments, _ := uc.restaurantData.CommentsData(response.ID)
		for i := 0; i < len(comments); i++ {
			response.Comments = append(response.Comments, restaurants.Comment{UserID: comments[i].UserID, Comment: comments[i].Comment})
		}

	}

	return response, err
}

func (uc *restaurantUseCase) MyRestoBusiness(idUser int) (response restaurants.CoreMyDetail, err error) {
	response, err = uc.restaurantData.MyRestoData(idUser)

	if err.Error() == "restaurant not found" {
		return restaurants.CoreMyDetail{}, fmt.Errorf("restaurant not found")
	}

	if err == nil {

		// get rating
		rating, _ := uc.restaurantData.RatingData(response.ID)
		response.Rating = rating

		// get resto image url
		restoImg, _ := uc.restaurantData.RestoImageData(response.ID)
		response.RestoImageUrl = restoImg

	}

	return response, err
}
