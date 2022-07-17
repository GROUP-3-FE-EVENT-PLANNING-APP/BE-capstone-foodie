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
