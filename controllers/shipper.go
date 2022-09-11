package controllers

import (
	"log"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/AlexVonEinzbern/go-rest-api/utils"
)

//TODO: Implement CreateShipper
func CreateShipper(createshipper models.ShipperCreate) models.Shipper {

	shipper := models.Shipper{
		ID: utils.IdGenerator(),
		ShipperBase: models.ShipperBase{
			CompanyName: createshipper.CompanyName,
			Phone:       createshipper.Phone}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with shipper creation")

	result := conn.Create(&shipper)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", shipper)

	return shipper
}

//TODO: Implement SearchShippers
func SearchShippers() []models.Shipper {
	var shippers []models.Shipper

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with shippers searching...")

	result := conn.Find(&shippers)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", shippers)

	return shippers
}

//TODO: Implement UpdateShipper
func UpdateShipper(id string, shipperUpdate map[string]interface{}) (models.Shipper, error) {
	var shipper models.Shipper

	conn := db.DBConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with shipper update")

	if err := conn.First(&shipper, "id = ?", id).Error; err != nil {
		log.Println("DB error ", err)
		return shipper, err

	}

	if CompanyName := shipperUpdate["CompanyName"].(string); CompanyName == "" {
		shipperUpdate["CompanyName"] = shipper.CompanyName
	}

	if Phone := shipperUpdate["Phone"].(string); Phone == "" {
		shipperUpdate["Phone"] = shipper.Phone
	}

	result := conn.Model(&shipper).Updates(shipperUpdate)
	//.Updates(characterUpdate)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return shipper, result.Error
	}

	return shipper, result.Error
}

//TODO: Implement DeleteShipper
func DeleteShipper(id string) (models.Shipper, error) {
	var shipper models.Shipper

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with shipper delete")

	result := conn.Delete(&shipper, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return shipper, result.Error
	}

	log.Println("Shipper deleted: ", shipper)

	return shipper, result.Error
}