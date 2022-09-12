package controllers

import (
	"log"
	"time"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/AlexVonEinzbern/go-rest-api/utils"
)

//TODO: Implement CreateLogin
func CreateLogin(createlogin models.LoginCreate) models.Login {

	login := models.Login{
		ID: utils.IdGenerator(),
		LoginBase: models.LoginBase{
			CustomerID: createlogin.CustomerID}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with login creation")

	result := conn.Create(&login)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", login)

	return login
}

//TODO: Implement SearchLogin
func SearchLogin(id string) (models.Login, error) {
	var login models.Login
	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with login searching by id")

	result := conn.First(&login, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return login, result.Error
	}

	log.Println("Found record: ", login)
	return login, result.Error
}

//TODO: Implement SearchLoginDate
func SearchLoginDate(date string) ([]models.Login, error) {
	var logins []models.Login
	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()
	start_date, _ := time.Parse("2006-01-02", date)
	end_date := start_date.AddDate(0, 0, 1)

	start := start_date.Format("2006-01-02")
	end := end_date.Format("2006-01-02")

	log.Println("Starting with login searching by date")
	result := conn.Where("created_at BETWEEN ? AND ?", start, end).Find(&logins)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return logins, result.Error
	}

	log.Println("Found record: ", logins)
	return logins, result.Error
}
