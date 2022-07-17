package business

import (
	"capstone/group3/features/admins"
)

type adminUseCase struct {
	adminData admins.Data
}

func NewAdminBusiness(admData admins.Data) admins.Business {
	return &adminUseCase{
		adminData: admData,
	}
}

func (uc *adminUseCase) AllUserBusiness(limit, offset, idUser int) (response []admins.Core, err error) {
	response, err = uc.adminData.AllUserData(limit, offset, idUser)

	return response, err
}

func (uc *adminUseCase) AllRestoBusiness(limit, offset, idUser int) (response []admins.CoreRestaurant, err error) {
	response, err = uc.adminData.AllRestoData(limit, offset, idUser)

	if err == nil {
		for i := 0; i < len(response); i++ {

			// get rating
			rating, _ := uc.adminData.RatingData(response[i].ID)
			response[i].Rating = rating

			// get resto image url
			restoImg, _ := uc.adminData.RestoImageData(response[i].ID)

			response[i].RestoImageUrl = restoImg
		}
	}

	return response, err
}

func (uc *adminUseCase) DetailRestoBusiness(id, idUser int) (response admins.CoreDetailRestaurant, err error) {
	response, err = uc.adminData.DetailRestoData(id, idUser)

	if err == nil {

		// get rating
		rating, _ := uc.adminData.RatingData(response.ID)
		response.Rating = rating

		// get resto image
		restoImgs, _ := uc.adminData.RestoImagesData(response.ID)
		for i := 0; i < len(restoImgs); i++ {
			response.RestoImages = append(response.RestoImages, admins.RestoImage{RestoImageUrl: restoImgs[i]})
		}

		// get facility
		facilities, _ := uc.adminData.FacilitiesData(response.ID)
		for i := 0; i < len(restoImgs); i++ {
			response.Facilities = append(response.Facilities, admins.Facility{Facility: facilities[i]})
		}

	}

	return response, err
}

func (uc *adminUseCase) VerifRestoBusiness(id, idUser int) (response int, err error) {
	response, err = uc.adminData.VerifRestoData(id, idUser)

	return response, err
}
