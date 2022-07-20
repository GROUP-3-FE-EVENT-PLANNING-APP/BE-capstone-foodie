package request

import (
	"capstone/group3/features/booking"
)

type Booking struct {
	TableQuota uint   `json:"table_quota" form:"table_quota" validate:"required,number"`
	BookingFee uint64 `json:"booking_fee" form:"booking_fee"`
	Date       string `json:"date" form:"date" validate:"required"`
	Time       string `json:"time" form:"time" validate:"required"`
}

func ToCore(req Booking) booking.Core {
	return booking.Core{
		TableQuota: req.TableQuota,
		Date:       req.Date,
		Time:       req.Time,
	}
}
