package presentation

import (
	"capstone/group3/features/restaurants"
	_requestRestaurant "capstone/group3/features/restaurants/presentation/request"
	"fmt"
	"time"

	// _responseRestaurant "capstone/group3/features/restaurants/presentation/response"
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

	// formfile data image menu
	fileData, fileInfo, fileErr := c.Request().FormFile("menu_image_url")

	// return err jika missing file
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
	}

	// cek ekstension file upload
	extension, err_check_extension := _helper.CheckFileExtension(fileInfo.Filename)
	if err_check_extension != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file extension error"))
	}

	// check file size
	err_check_size := _helper.CheckFileSize(fileInfo.Size)
	if err_check_size != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("file size error"))
	}

	// memberikan nama file
	fileName := time.Now().Format("2006-01-02 15:04:05") + "." + extension

	url, errUploadImg := _helper.UploadImageToS3(fileName, fileData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to upload file"))
	}

	// ekstrak token
	idToken, _, _, _, _, errToken := _middlewares.ExtractToken(c)

	// return jika errorToken
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	// inissialisasi newResto.UserId = idToken(userid)
	newResto.UserId = idToken
	//
	newResto.MenuImageUrl = url

	dataResto := _requestRestaurant.ToCore(newResto)
	_, err := h.RestaurantBusiness.CreateRestoBusiness(dataResto)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))

	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))

}
