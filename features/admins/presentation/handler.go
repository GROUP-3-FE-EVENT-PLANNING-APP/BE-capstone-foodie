package presentation

import (
	"capstone/group3/features/admins"
	"strconv"

	_responseAdmin "capstone/group3/features/admins/presentation/response"
	_middlewares "capstone/group3/features/middlewares"
	_helper "capstone/group3/helper"

	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminBusiness admins.Business
}

func NewAdminHandler(business admins.Business) *AdminHandler {
	return &AdminHandler{
		AdminBusiness: business,
	}
}

func (h *AdminHandler) AllUser(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)

	// ekstrak token
	data, errToken := _middlewares.ExtractToken(c)
	idToken := data["userId"].(float64)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	result, err := h.AdminBusiness.AllUserBusiness(limitint, offsetint, int(idToken))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseAdmin.FromCoreList(result)))

}

func (h *AdminHandler) AllResto(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)

	// ekstrak token
	data, errToken := _middlewares.ExtractToken(c)
	idToken := data["userId"].(float64)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	result, err := h.AdminBusiness.AllRestoBusiness(limitint, offsetint, int(idToken))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseAdmin.FromCoreRestaurantList(result)))

}

func (h *AdminHandler) DetailResto(c echo.Context) error {
	// ekstrak token
	data, errToken := _middlewares.ExtractToken(c)
	idToken := data["userId"].(float64)

	id := c.Param("id")

	idResto, _ := strconv.Atoi(id)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	result, err := h.AdminBusiness.DetailRestoBusiness(idResto, int(idToken))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to detail data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseAdmin.FromCoreDetail(result)))

}
