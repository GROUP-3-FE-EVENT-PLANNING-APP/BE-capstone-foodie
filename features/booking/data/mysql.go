package data

import (
	"capstone/group3/features/booking"
	_helper "capstone/group3/helper"
	"fmt"

	"gorm.io/gorm"
)

type mysqlBookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(conn *gorm.DB) booking.Data {
	return &mysqlBookingRepository{
		db: conn,
	}
}

func (repo *mysqlBookingRepository) BookingRestoData(data booking.Core) (response int, err error) {
	dataBooking := fromCore(data)
	result := repo.db.Create(&dataBooking)

	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to booking data")
	}

	return int(result.RowsAffected), err
}

func (repo *mysqlBookingRepository) CheckTableReservedData(idResto int) (response int, err error) {
	var datas []Booking

	result := repo.db.Table("bookings").Where("restaurant_id = ? AND booking_status = ?", idResto, "active").Find(&datas)

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, nil
	}

	tableQouta := 0
	for _, v := range datas {
		tableQouta += int(v.TableQuota)
	}

	return tableQouta, err
}

func (repo *mysqlBookingRepository) CheckTableQuotaData(idResto int) (response booking.Restaurant, err error) {
	var data Restaurant

	result := repo.db.Where("id = ?", idResto).First(&data)

	if result.Error != nil {
		return booking.Restaurant{}, result.Error
	}

	if result.RowsAffected != 1 {
		return booking.Restaurant{}, fmt.Errorf("failed to check table quota data")
	}

	response.TableQuota = data.TableQuota
	response.BookingFee = data.BookingFee

	return response, err
}

func (repo *mysqlBookingRepository) GetUserData(idUser int) (response _helper.DetailPayment, err error) {
	var data User

	result := repo.db.Where("id = ?", idUser).First(&data)

	if result.Error != nil {
		return _helper.DetailPayment{}, result.Error
	}

	if result.RowsAffected != 1 {
		return _helper.DetailPayment{}, fmt.Errorf("failed to check user data")
	}

	response.Name = data.Name
	response.Email = data.Email
	response.Handphone = data.Handphone

	return response, err
}

func (repo *mysqlBookingRepository) PaymentData(data booking.PaymentWebhook) (row int, response booking.Core, email, name string, err error) {
	var dataBooking Booking

	if data.TransactionStatus != "settlement" {
		return -1, booking.Core{}, "", "", fmt.Errorf("status not match")
	}

	// update booking
	updateBooking := repo.db.Table("bookings").Where("order_id = ?", data.OrderID).Update("payment_status", "accepted")

	if updateBooking.Error != nil {
		return 0, booking.Core{}, "", "", updateBooking.Error
	}

	if updateBooking.RowsAffected != 1 {
		return 0, booking.Core{}, "", "", fmt.Errorf("order_id not found")
	}

	// get booking
	getBooking := repo.db.Preload("User").Where("order_id = ?", data.OrderID).First(&dataBooking)

	if getBooking.Error != nil {
		return 0, booking.Core{}, "", "", getBooking.Error
	}

	if getBooking.RowsAffected != 1 {
		return 0, booking.Core{}, "", "", fmt.Errorf("order_id not found")
	}
	fmt.Println("namemysql: ", dataBooking.User.Name)
	fmt.Println("emailmysql: ", dataBooking.User.Email)
	return int(updateBooking.RowsAffected), toCore(dataBooking), dataBooking.User.Email, dataBooking.User.Name, err
}
