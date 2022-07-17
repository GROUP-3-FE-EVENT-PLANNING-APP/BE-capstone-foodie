package presentation

import (
	"capstone/group3/features/favourites"
	_responseFavourite "capstone/group3/features/favourites/presentation/response"
	_middleware "capstone/group3/features/middlewares"
	_helper "capstone/group3/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FavouriteHandler struct {
	favouriteBusiness favourites.Business
}

func NewFavouriteHandler(favBusiness favourites.Business) *FavouriteHandler {
	return &FavouriteHandler{
		favouriteBusiness: favBusiness,
	}
}

func (h *FavouriteHandler) PostFav(c echo.Context) error {
	id := c.Param("id")
	idResto, _ := strconv.Atoi(id)
	fromToken, _ := _middleware.ExtractToken(c)
	idFromToken := fromToken["userId"].(float64)
	row, errCreate := h.favouriteBusiness.AddFav(idResto, int(idFromToken))

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to add resto into your favourites list"))
	}

	if row == -1 {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to add, resto is already in your favourites list"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to add resto into your favourites list"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *FavouriteHandler) DeleteFav(c echo.Context) error {
	id := c.Param("id")
	idResto, _ := strconv.Atoi(id)
	fromToken, _ := _middleware.ExtractToken(c)
	idFromToken := fromToken["userId"].(float64)
	row, errCreate := h.favouriteBusiness.DeleteFavourite(idResto, int(idFromToken))

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete resto from your favourites list"))
	}

	if row != 1 {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete resto from your favourites list"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *FavouriteHandler) MyFav(c echo.Context) error {
	fromToken, _ := _middleware.ExtractToken(c)
	idFromToken, _ := fromToken["userId"].(float64)
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	data, errGet := h.favouriteBusiness.GetMyFav(limitint, offsetint, int(idFromToken))

	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get my favourites resto data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseFavourite.FromCoreList(data)))
}
