package booking

import (
	_helper "capstone/group3/helper"
)

type Core struct {
	UserID              uint
	User                User
	RestaurantID        uint
	TransactionID       string
	OrderID             string
	TableQuota          uint
	BookingFee          uint64
	PaymentStatus       string
	BookingStatus       string
	MidtransToken       string
	MidtransRedirectURL string
	Date                string
	Time                string
}

type User struct {
	Name      string
	Email     string
	Handphone string
}

type Restaurant struct {
	ID         uint
	TableQuota uint
	BookingFee uint64
}

type PaymentWebhook struct {
	TransactionStatus string
	OrderID           string
}

type Business interface {
	BookingRestoBusiness(data Core) (row int, token, redirectURL string, err error)
	PaymentBusiness(data PaymentWebhook) (row int, err error)
}

type Data interface {
	BookingRestoData(data Core) (response int, err error)
	CheckTableReservedData(idResto int) (response int, err error)
	CheckTableQuotaData(idResto int) (response Restaurant, err error)
	GetUserData(idUser int) (response _helper.DetailPayment, err error)
	PaymentData(data PaymentWebhook) (row int, err error)
}
