package controllers

import (
	"log"
	"time"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/AlexVonEinzbern/go-rest-api/utils"
)

//TODO: Implement CreateCustomer
func CreateCustomer(createcustomer models.CustomerCreate) models.Customer {

	// parse the input string into a time.Time value
	t, err := time.Parse("2006-01-02", createcustomer.Birthday)
	if err != nil {
		log.Fatal(err)
	}

	// convert the time.Time value to a string in the desired layout
	date_birthday := t.Format("2006-01-02T15:04:05Z07:00")

	customer := models.Customer{
		ID: utils.IdGenerator(),
		CustomerBase: models.CustomerBase{
			Name:       createcustomer.Name,
			Email:      createcustomer.Email,
			DNI:        createcustomer.DNI,
			Address:    createcustomer.Address,
			City:       createcustomer.City,
			Birthday:   date_birthday,
			PostalCode: createcustomer.PostalCode,
			Country:    createcustomer.Country,
			Phone:      createcustomer.Phone,
			UserName:   createcustomer.UserName,
			Password:   createcustomer.Password}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with customer creation")

	result := conn.Create(&customer)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", customer)

	return customer
}

//TODO: Implement SearchCustomers
func SearchCustomers() []models.Customer {
	var customers []models.Customer

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with customers searching...")

	result := conn.Find(&customers)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", customers)

	return customers
}

//TODO: Implement SearchCustomer
func SearchCustomer(id string) (models.Customer, error) {
	var customer models.Customer
	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with customer searching by id")

	result := conn.First(&customer, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return customer, result.Error
	}

	log.Println("Found record: ", customer)
	return customer, result.Error
}

//TODO: Implement UpdateCustomer
func UpdateCustomer(id string, customerUpdate map[string]interface{}) (models.Customer, error) {
	var customer models.Customer

	conn := db.DBConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with customer update")

	if err := conn.First(&customer, "id = ?", id).Error; err != nil {
		log.Println("DB error ", err)
		return customer, err

	}

	if Name := customerUpdate["Name"].(string); Name == "" {
		customerUpdate["Name"] = customer.Name
	}

	if Email := customerUpdate["Email"].(string); Email == "" {
		customerUpdate["Email"] = customer.Email
	}

	if DNI := customerUpdate["DNI"].(string); DNI == "" {
		customerUpdate["DNI"] = customer.DNI
	}

	if Address := customerUpdate["Address"].(string); Address == "" {
		customerUpdate["Address"] = customer.Address
	}

	if City := customerUpdate["City"].(string); City == "" {
		customerUpdate["City"] = customer.City
	}

	if Birthday := customerUpdate["Birthday"].(time.Time); Birthday.IsZero() {
		customerUpdate["Birthday"] = customer.Birthday
	}

	if PostalCode := customerUpdate["PostalCode"].(string); PostalCode == "" {
		customerUpdate["PostalCode"] = customer.PostalCode
	}

	if Country := customerUpdate["Country"].(string); Country == "" {
		customerUpdate["Country"] = customer.Country
	}

	if Phone := customerUpdate["Phone"].(string); Phone == "" {
		customerUpdate["Phone"] = customer.Phone
	}

	if UserName := customerUpdate["UserName"].(string); UserName == "" {
		customerUpdate["UserName"] = customer.UserName
	}

	if Password := customerUpdate["Password"].(string); Password == "" {
		customerUpdate["Password"] = customer.Password
	}

	result := conn.Model(&customer).Updates(customerUpdate)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return customer, result.Error
	}

	return customer, result.Error
}

//TODO: Implement DeleteCustomer
func DeleteCustomer(id string) (models.Customer, error) {
	var customer models.Customer

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with customer delete")

	result := conn.Delete(&customer, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return customer, result.Error
	}

	log.Println("Customer deleted: ", customer)

	return customer, result.Error
}
