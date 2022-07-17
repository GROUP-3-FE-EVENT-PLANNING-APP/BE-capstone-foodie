package presentation

import (
	"capstone/group3/features/restaurants"
	_requestRestaurant "capstone/group3/features/restaurants/presentation/request"
	_responseRestaurant "capstone/group3/features/restaurants/presentation/response"

	"strconv"

	_middlewares "capstone/group3/features/middlewares"
	_helper "capstone/group3/helper"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RestaurantHandler struct {
	RestaurantBusiness restaurants.Business
}

func NewRestaurantHandler(business restaurants.Business) *RestaurantHandler {
	return &RestaurantHandler{
		RestaurantBusiness: business,
	}
}

func (h *RestaurantHandler) CreateResto(c echo.Context) error {
	// inisialiasi variabel dengan type struct dari request
	var newResto _requestRestaurant.Restaurant

	// binding data resto
	errBind := c.Bind(&newResto)

	validate := validator.New()
	if errValidate := validate.Struct(newResto); errValidate != nil {
		return errValidate
	}

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	// menu image url
	fileData, fileInfo, fileErr := c.Request().FormFile("menu_image_url")

	// return err jika missing file
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
	}

	uploadImage, errUploadImage := _helper.UploadImage(fileData, fileInfo)

	if errUploadImage != nil {
		if errUploadImage.Error() == "failed to get file" {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
		}

		if errUploadImage.Error() == "file extension error" {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
		}

		if errUploadImage.Error() == "file size error" {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
		}

		if errUploadImage.Error() == "failed to upload file" {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file"))
		}
	}

	// file image url
	// menu image url
	fileData_, fileInfo_, fileErr_ := c.Request().FormFile("file_image_url")

	// return err jika missing file
	if fileErr_ == http.ErrMissingFile || fileErr_ != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
	}

	uploadFileUrl, errUploadFileUrl := _helper.UploadImage(fileData_, fileInfo_)

	if errUploadFileUrl != nil {
		if errUploadFileUrl.Error() == "failed to get file" {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
		}

		if errUploadFileUrl.Error() == "file extension error" {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
		}

		if errUploadFileUrl.Error() == "file size error" {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
		}

		if errUploadFileUrl.Error() == "failed to upload file" {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file"))
		}
	}

	// ekstrak token
	data, errToken := _middlewares.ExtractToken(c)
	idToken := data["userId"].(float64)

	// return jika errorToken
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	// inissialisasi newResto.UserId = idToken(userid)
	newResto.UserId = int(idToken)
	//
	newResto.MenuImageUrl = uploadImage
	newResto.FileImageUrl = uploadFileUrl

	dataResto := _requestRestaurant.ToCore(newResto)
	_, err := h.RestaurantBusiness.CreateRestoBusiness(dataResto)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))

	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))

}

func (h *RestaurantHandler) UpdateResto(c echo.Context) error {
	var editResto _requestRestaurant.Restaurant

	errBind := c.Bind(&editResto)

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	// ekstrak token
	data, errToken := _middlewares.ExtractToken(c)
	idToken := data["userId"].(float64)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	// formfile data image menu
	fileData, fileInfo, fileErr := c.Request().FormFile("menu_image_url")

	if fileErr == http.ErrMissingFile || fileErr != nil {
		// tidak ingin update image
		// get image lama
		resImg, _, err := h.RestaurantBusiness.DetailImageRestoBusiness(int(idToken))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to update data"))
		}

		editResto.MenuImageUrl = resImg
	} else {
		// ingin update image menu

		uploadImageUrl, errUploadImageUrl := _helper.UploadImage(fileData, fileInfo)

		if errUploadImageUrl != nil {
			if errUploadImageUrl.Error() == "failed to get file" {
				return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
			}

			if errUploadImageUrl.Error() == "file extension error" {
				return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
			}

			if errUploadImageUrl.Error() == "file size error" {
				return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
			}

			if errUploadImageUrl.Error() == "failed to upload file" {
				return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file"))
			}
		}

		//
		editResto.MenuImageUrl = uploadImageUrl
	}

	// file image url
	fileData_, fileInfo_, fileErr_ := c.Request().FormFile("file_image_url")

	if fileErr_ == http.ErrMissingFile || fileErr_ != nil {
		// tidak ingin update image
		// get image lama
		_, fileImage, err := h.RestaurantBusiness.DetailImageRestoBusiness(int(idToken))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to update data"))
		}

		editResto.FileImageUrl = fileImage
	} else {
		// ingin update image menu

		uploadFileUrl, errUploadFileUrl := _helper.UploadImage(fileData_, fileInfo_)

		if errUploadFileUrl != nil {
			if errUploadFileUrl.Error() == "failed to get file" {
				return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
			}

			if errUploadFileUrl.Error() == "file extension error" {
				return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
			}

			if errUploadFileUrl.Error() == "file size error" {
				return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
			}

			if errUploadFileUrl.Error() == "failed to upload file" {
				return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file"))
			}
		}

		//
		editResto.FileImageUrl = uploadFileUrl
	}

	dtResto := _requestRestaurant.ToCore(editResto)

	_, err := h.RestaurantBusiness.UpdateRestoBusiness(dtResto, int(idToken))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to update data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *RestaurantHandler) DestroyResto(c echo.Context) error {
	// ekstrak token
	data, errToken := _middlewares.ExtractToken(c)
	idToken := data["userId"].(float64)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	_, err := h.RestaurantBusiness.DeleteRestoBusiness(int(idToken))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *RestaurantHandler) UploadImageResto(c echo.Context) error {
	var newImage _requestRestaurant.RestoImage

	// menu image url
	fileData, fileInfo, fileErr := c.Request().FormFile("resto_image_url")

	// return err jika missing file
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
	}

	uploadImage, errUploadImage := _helper.UploadImage(fileData, fileInfo)

	if errUploadImage != nil {
		if errUploadImage.Error() == "failed to get file" {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
		}

		if errUploadImage.Error() == "file extension error" {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
		}

		if errUploadImage.Error() == "file size error" {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
		}

		if errUploadImage.Error() == "failed to upload file" {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file"))
		}
	}

	// ekstrak token
	data, errToken := _middlewares.ExtractToken(c)
	idToken := data["userId"].(float64)

	// return jika errorToken
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	newImage.UserId = int(idToken)
	newImage.RestoImageUrl = uploadImage

	dataImage := _requestRestaurant.ToCoreRestoImage(newImage)
	_, err := h.RestaurantBusiness.UploadImageRestoBusiness(dataImage)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload data"))

	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))

}

func (h *RestaurantHandler) AllResto(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)

	result, err := h.RestaurantBusiness.AllRestoBusiness(limitint, offsetint)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseRestaurant.FromCoreList(result)))
}

func (h *RestaurantHandler) DetailResto(c echo.Context) error {
	id := c.Param("id")

	idResto, _ := strconv.Atoi(id)

	result, err := h.RestaurantBusiness.DetailRestoBusiness(idResto)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to detail data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseRestaurant.FromCoreDetail(result)))

}

func (h *RestaurantHandler) MyResto(c echo.Context) error {
	// ekstrak token
	data, errToken := _middlewares.ExtractToken(c)
	idToken := data["userId"].(float64)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	result, err := h.RestaurantBusiness.MyRestoBusiness(int(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseRestaurant.FromCoreDetailMyResto(result)))

}
