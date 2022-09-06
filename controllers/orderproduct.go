package controllers

import (
	"log"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
)

//TODO: Implement CreateOrderProduct
func CreateOrderProduct(createorderproduct models.OrderProductCreate) models.OrderProduct {

	orderproduct := models.OrderProduct{
		OrderProductBase: models.OrderProductBase{
			UnitPrice: createorderproduct.UnitPrice,
			Quantity:  createorderproduct.Quantity,
			Discount:  createorderproduct.Discount,
			ProductID: createorderproduct.ProductID,
			OrderID:   createorderproduct.OrderID}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with orderproduct creation")

	result := conn.Create(&orderproduct)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", orderproduct)

	return orderproduct
}

//TODO: Implement SearchOrderProducts
func SearchOrderProducts() []models.OrderProduct {
	var orderproducts []models.OrderProduct

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with orderproducts searching...")

	result := conn.Find(&orderproducts)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", orderproducts)

	return orderproducts
}

//TODO: Implement SearchOrderProduct
func SearchOrderProduct(id string) (models.OrderProduct, error) {
	var orderproduct models.OrderProduct
	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with orderproduct searching by id")

	result := conn.First(&orderproduct, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return orderproduct, result.Error
	}

	log.Println("Found record: ", orderproduct)
	return orderproduct, result.Error
}
