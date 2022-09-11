package controllers

import (
	"log"
	"time"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/AlexVonEinzbern/go-rest-api/utils"
)

//TODO: Implement CreatePayment
func CreatePayment(createpayment models.PaymentCreate) models.Payment {

	payment := models.Payment{
		ID: utils.IdGenerator(),
		PaymentBase: models.PaymentBase{
			OrderID:      createpayment.OrderID,
			CreditCardID: createpayment.CreditCardID}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with login creation")

	result := conn.Create(&payment)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", payment)

	return payment
}

//TODO: Implement SearchPayment
func SearchPayment(id string) (models.Payment, error) {
	var payment models.Payment
	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with payment searching by id")

	result := conn.First(&payment, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return payment, result.Error
	}

	log.Println("Found record: ", payment)
	return payment, result.Error
}

//TODO: Implement SearchPaymentDate
func SearchPaymentDate(date string) ([]models.Payment, error) {
	var payments []models.Payment
	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	start_date, _ := time.Parse("2006-01-02", date)
	end_date := start_date.AddDate(0, 0, 1)

	log.Println("Starting with payments searching by date")

	result := conn.Where("created_at between ? and ?", start_date, end_date).Find(&payments)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return payments, result.Error
	}

	log.Println("Found record: ", payments)
	return payments, result.Error
}