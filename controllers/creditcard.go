package controllers

import (
	"log"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/AlexVonEinzbern/go-rest-api/utils"
)

//TODO: Implement CreateCreditCard
func CreateCreditCard(createcreditcard models.CreditCardCreate) models.CreditCard {

	creditcard := models.CreditCard{
		ID: utils.IdGenerator(),
		CreditCardBase: models.CreditCardBase{
			Brand:      createcreditcard.Brand,
			Number:     createcreditcard.Number,
			CVV:        createcreditcard.CVV,
			MM:         createcreditcard.MM,
			YYYY:       createcreditcard.YYYY,
			CustomerID: createcreditcard.CustomerID}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with creditcard creation")

	result := conn.Create(&creditcard)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", creditcard)

	return creditcard
}

//TODO: Implement SearchCreditCards
func SearchCreditCards() []models.CreditCard {
	var creditcards []models.CreditCard

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with credit cards searching...")

	result := conn.Find(&creditcards)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", creditcards)

	return creditcards
}

//TODO: Implement SearchCreditCard
func SearchCreditCard(id string) (models.CreditCard, error) {
	var creditcard models.CreditCard
	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with credit card searching by id")

	result := conn.First(&creditcard, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return creditcard, result.Error
	}

	log.Println("Found record: ", creditcard)
	return creditcard, result.Error
}

//TODO: Implement UpdateCreditCard
func UpdateCreditCard(id string, creditcardUpdate map[string]interface{}) (models.CreditCard, error) {
	var creditcard models.CreditCard

	conn := db.DBConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with credit card update")

	if err := conn.First(&creditcard, "id = ?", id).Error; err != nil {
		log.Println("DB error ", err)
		return creditcard, err

	}

	if Brand := creditcardUpdate["Brand"].(string); Brand == "" {
		creditcardUpdate["Brand"] = creditcard.Brand
	}

	if Number := creditcardUpdate["Number"].(string); Number == "" {
		creditcardUpdate["Number"] = creditcard.Number
	}

	if CVV := creditcardUpdate["CVV"].(string); CVV == "" {
		creditcardUpdate["CVV"] = creditcard.CVV
	}

	if MM := creditcardUpdate["MM"].(string); MM == "" {
		creditcardUpdate["MM"] = creditcard.MM
	}

	if YYYY := creditcardUpdate["YYYY"].(string); YYYY == "" {
		creditcardUpdate["YYYY"] = creditcard.YYYY
	}

	if CustomerID := creditcardUpdate["CustomerID"].(string); CustomerID == "" {
		creditcardUpdate["CustomerID"] = creditcard.CustomerID
	}

	result := conn.Model(&creditcard).Updates(creditcardUpdate)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return creditcard, result.Error
	}

	return creditcard, result.Error
}

//TODO: Implement DeleteCreditCard
func DeleteCreditCard(id string) (models.CreditCard, error) {
	var creditcard models.CreditCard

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with credit card delete")

	result := conn.Delete(&creditcard, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return creditcard, result.Error
	}

	log.Println("CreditCard deleted: ", creditcard)

	return creditcard, result.Error
}
