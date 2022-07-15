package presentation

import (
	"fmt"
	"net/http"
	"time"

	_middleware "capstone/group3/features/middlewares"
	"capstone/group3/features/users"
	_requestUser "capstone/group3/features/users/presentation/request"
	_responseUser "capstone/group3/features/users/presentation/response"
	_helper "capstone/group3/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(userBusiness users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: userBusiness,
	}
}

func (h *UserHandler) PostUser(c echo.Context) error {
	userReq := _requestUser.User{}
	err := c.Bind(&userReq)

	validate := validator.New()
	if errValidate := validate.Struct(userReq); errValidate != nil {
		return errValidate
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	dataUser := _requestUser.ToCore(userReq)
	row, errCreate := h.userBusiness.CreateData(dataUser)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your email or handphone number is already registered"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *UserHandler) LoginAuth(c echo.Context) error {
	authData := users.AuthRequestData{}
	c.Bind(&authData)
	fromToken, e := h.userBusiness.LoginUser(authData)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
	}

	data := map[string]interface{}{
		"id":        fromToken["id"],
		"token":     fromToken["token"],
		"name":      fromToken["name"],
		"avatarUrl": fromToken["avatarUrl"],
		"role":      fromToken["role"],
		"handphone": fromToken["handphone"],
		"email":     fromToken["email"],
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("login success", data))
}

func (h *UserHandler) PutUser(c echo.Context) error {
	fromToken, _ := _middleware.ExtractToken(c)
	idFromToken := fromToken["userId"].(float64)
	userReq := _requestUser.User{}
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("avatar_url")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get file"))
		}

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

		userReq.AvatarUrl = url
	}

	dataUser := _requestUser.ToCore(userReq)
	row, errUpd := h.userBusiness.UpdateData(dataUser, int(idFromToken))
	if errUpd != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to update data users"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data users"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *UserHandler) GetByMe(c echo.Context) error {
	fromToken, _ := _middleware.ExtractToken(c)
	idFromToken := fromToken["userId"].(float64)
	result, errGet := h.userBusiness.GetUserByMe(int(idFromToken))
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data user"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseUser.FromCore(result)))
}

func (h *UserHandler) DeleteByID(c echo.Context) error {
	fromromToken, _ := _middleware.ExtractToken(c)
	idFromToken := fromromToken["userId"].(float64)

	row, errDel := h.userBusiness.DeleteDataById(int(idFromToken))
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data user"))
	}
	if row != 1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data user"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}
