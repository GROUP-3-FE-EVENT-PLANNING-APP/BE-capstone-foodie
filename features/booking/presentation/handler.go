package presentation

import (
	"capstone/group3/features/booking"
	_requestBooking "capstone/group3/features/booking/presentation/request"
	_responseBooking "capstone/group3/features/booking/presentation/response"
	_middlewares "capstone/group3/features/middlewares"
	_helper "capstone/group3/helper"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	BookingBusiness booking.Business
}

func NewBookingHandler(business booking.Business) *BookingHandler {
	return &BookingHandler{
		BookingBusiness: business,
	}
}

func (h *BookingHandler) BookingResto(c echo.Context) error {
	// inisialiasi variabel dengan type struct dari request
	var dataBooking _requestBooking.Booking

	// binding data resto
	errBind := c.Bind(&dataBooking)

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	// get params
	id := c.Param("id")

	idResto, _ := strconv.Atoi(id)

	// ekstrak token
	data, errToken := _middlewares.ExtractToken(c)
	idToken := data["userId"].(float64)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	dataBooking_ := _requestBooking.ToCore(dataBooking)
	dataBooking_.UserID = uint(idToken)
	dataBooking_.RestaurantID = uint(idResto)

	response, Token, RedirectURL, err := h.BookingBusiness.BookingRestoBusiness(dataBooking_)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}

	if response == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}

	if response == -2 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("the restaurant available table, not enough to serve your request"))
	}

	if response == -3 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("booking failed"))
	}

	var dataResponse _responseBooking.Booking

	dataResponse.Token = Token
	dataResponse.RedirectURL = RedirectURL

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", dataResponse))

}

func (h *BookingHandler) AcceptPayment(c echo.Context) error {

	var data _requestBooking.PaymentWebhook

	// binding data
	errBind := c.Bind(&data)

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	dataPayment := _requestBooking.ToCorePaymentWebhook(data)
	response, err := h.BookingBusiness.PaymentBusiness(dataPayment)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to accept your payment"))
	}

	if response == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("transaction status not match"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("payment successful"))

}
