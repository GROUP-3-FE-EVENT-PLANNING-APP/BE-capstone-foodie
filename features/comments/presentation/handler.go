package presentation

import (
	"capstone/group3/features/comments"
	_requestComment "capstone/group3/features/comments/presentation/request"
	_responseComment "capstone/group3/features/comments/presentation/response"
	_middleware "capstone/group3/features/middlewares"
	_helper "capstone/group3/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentBusiness comments.Business
}

func NewCommentHandler(commentBusiness comments.Business) *CommentHandler {
	return &CommentHandler{
		commentBusiness: commentBusiness,
	}
}

func (h *CommentHandler) PostComment(c echo.Context) error {
	userReq := _requestComment.Comment{}
	fromToken, _ := _middleware.ExtractToken(c)
	idFromToken := fromToken["userId"].(float64)
	id := c.Param("id")
	idResto, _ := strconv.Atoi(id)
	err := c.Bind(&userReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	dataComment := _requestComment.ToCore(userReq)
	dataComment.User.ID = int(idFromToken)
	dataComment.Restaurant.ID = idResto
	row, errCreate := h.commentBusiness.CreateComment(dataComment)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}

	if row == -2 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("rating cant be lower than 1"))
	}

	if row == -3 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("rating cant be greater than 5"))
	}

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to post comment and rating"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *CommentHandler) GetComment(c echo.Context) error {
	id := c.Param("id")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	idResto, _ := strconv.Atoi(id)
	result, errGet := h.commentBusiness.GetCommentByIdResto(idResto, limitint, offsetint)
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data comment"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseComment.FromCoreList(result)))
}

func (h *CommentHandler) GetRating(c echo.Context) error {
	id := c.Param("id")
	idResto, _ := strconv.Atoi(id)
	result, errGet := h.commentBusiness.GetRatingByIdResto(idResto)
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data rating"))
	}
	data := _responseComment.RatingResto{}
	data.Rating = result
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", data))
}
