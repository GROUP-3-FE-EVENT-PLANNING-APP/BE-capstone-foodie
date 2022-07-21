package business

import (
	"capstone/group3/features/booking"
	_helper "capstone/group3/helper"
	"fmt"
	"time"
)

type bookingUseCase struct {
	bookingData booking.Data
}

func NewBookingBusiness(bkData booking.Data) booking.Business {
	return &bookingUseCase{
		bookingData: bkData,
	}
}

func (uc *bookingUseCase) BookingRestoBusiness(data booking.Core) (row int, token, redirectURL string, err error) {
	if data.TableQuota == 0 || data.Date == "" || data.Time == "" {
		return -1, "", "", fmt.Errorf("please make sure all fields are filled in correctly")
	}

	// check table quota
	tableReserved, errCheckReserved := uc.bookingData.CheckTableReservedData(int(data.RestaurantID))

	if errCheckReserved != nil {
		return 0, "", "", errCheckReserved
	}

	// get table quota resto
	dataResto, errCheck := uc.bookingData.CheckTableQuotaData(int(data.RestaurantID))

	if errCheck != nil {
		return 0, "", "", errCheck
	}

	if (int(dataResto.TableQuota) - tableReserved) < int(data.TableQuota) {
		return -2, "", "", fmt.Errorf("the restaurant available table, not enough to serve your request")
	}

	totalPayment := data.TableQuota * uint(dataResto.BookingFee)
	data.BookingFee = uint64(totalPayment)

	// get user
	getUser, _ := uc.bookingData.GetUserData(int(data.UserID))

	getUser.TotalPayment = data.BookingFee

	// payment gateway
	getOrderID, getPayment := _helper.Payment(getUser)

	data.OrderID = getOrderID
	data.MidtransToken = getPayment.Token
	data.MidtransRedirectURL = getPayment.RedirectURL
	data.PaymentStatus = "pending"
	data.BookingStatus = "active"

	// if true payment gateway send email

	response, err := uc.bookingData.BookingRestoData(data)

	if err != nil {
		return -3, "", "", fmt.Errorf("booking failed")
	}

	return response, getPayment.Token, getPayment.RedirectURL, err
}

func (uc *bookingUseCase) PaymentBusiness(data booking.PaymentWebhook) (row int, err error) {
	row, response, err := uc.bookingData.PaymentData(data)

	//init the loc
	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	convertGMT7 := response.UpdatedAt.In(loc)

	convertString := convertGMT7.Format("02-01-2006 15:04:05")

	// format rupiah
	formatRupiah := _helper.FormatRupiah(response.BookingFee)

	if row == 1 {
		_helper.SendEmail(_helper.Recipient{
			Name:         response.User.Name,
			Email:        response.User.Email,
			OrderID:      response.OrderID,
			TotalPayment: formatRupiah,
			PaymentTime:  convertString,
		})
	}

	return row, err
}
