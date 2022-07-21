package data

import (
	"capstone/group3/features/booking"
	"time"

	"gorm.io/gorm"
)

type Restaurant struct {
	ID         uint
	TableQuota uint
	BookingFee uint64
	Booking    []Booking `gorm:"foreignKey:RestaurantID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	ID        uint
	Name      string
	Email     string    `gorm:"unique"`
	Handphone string    `gorm:"unique"`
	Booking   []Booking `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type Booking struct {
	// gorm.Model
	ID            uint `gorm:"primaryKey;autoIncrement"`
	UserID        uint `json:"user_id" form:"user_id"`
	User          User
	RestaurantID  uint `json:"restaurant_id" form:"restaurant_id"`
	Restaurant    Restaurant
	TransactionID string `json:"transaction_id" form:"transaction_id" gorm:"type:varchar(255)"`
	OrderID       string `json:"order_id" form:"order_id" gorm:"type:varchar(255)"`
	TableQuota    uint   `json:"table_quota" form:"table_quota" gorm:"not null;type:integer"`
	BookingFee    uint64 `json:"booking_fee" form:"booking_fee" gorm:"not null;type:bigint(20)"`
	PaymentStatus string `gorm:"type:varchar(100)"`
	BookingStatus string `gorm:"type:varchar(100)"`
	Date          string `json:"date" form:"date" gorm:"not null;type:date"`
	Time          string `json:"time" form:"time" gorm:"not null;type:varchar(10)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func fromCore(core booking.Core) Booking {
	return Booking{
		UserID:        core.UserID,
		RestaurantID:  core.RestaurantID,
		TransactionID: core.TransactionID,
		OrderID:       core.OrderID,
		TableQuota:    core.TableQuota,
		BookingFee:    core.BookingFee,
		PaymentStatus: core.PaymentStatus,
		BookingStatus: core.BookingStatus,
		Date:          core.Date,
		Time:          core.Time,
	}
}

func toCore(data Booking) booking.Core {
	return booking.Core{
		UserID:        data.UserID,
		RestaurantID:  data.RestaurantID,
		TransactionID: data.TransactionID,
		OrderID:       data.OrderID,
		TableQuota:    data.TableQuota,
		BookingFee:    data.BookingFee,
		PaymentStatus: data.PaymentStatus,
		BookingStatus: data.BookingStatus,
		Date:          data.Date,
		Time:          data.Time,
		UpdatedAt:     data.UpdatedAt,
	}
}
