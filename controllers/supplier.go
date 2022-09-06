package controllers

import (
	"log"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
)

//TODO: Implement CreateSupplier
func CreateSupplier(createsupplier models.SupplierCreate) models.Supplier {

	supplier := models.Supplier{
		ID: createsupplier.ID,
		SupplierBase: models.SupplierBase{
			CompanyName: createsupplier.CompanyName,
			Address:     createsupplier.Address,
			City:        createsupplier.City,
			Country:     createsupplier.Country,
			PostalCode:  createsupplier.PostalCode,
			Phone:       createsupplier.Phone,
			HomePage:    createsupplier.HomePage}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with subcategory creation")

	result := conn.Create(&supplier)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", supplier)

	return supplier
}

//TODO: Implement SearchSuppliers
func SearchSuppliers() []models.Supplier {
	var suppliers []models.Supplier

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with suppliers searching...")

	result := conn.Find(&suppliers)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", suppliers)

	return suppliers
}

//TODO: Implement UpdateSupplier
func UpdateSupplier(id string, supplierUpdate map[string]interface{}) (models.Supplier, error) {
	var supplier models.Supplier

	conn := db.DBConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with supplier update")

	if err := conn.First(&supplier, "id = ?", id).Error; err != nil {
		log.Println("DB error ", err)
		return supplier, err

	}

	if CompanyName := supplierUpdate["CompanyName"].(string); CompanyName == "" {
		supplierUpdate["CompanyName"] = supplier.CompanyName
	}

	if Address := supplierUpdate["Address"].(string); Address == "" {
		supplierUpdate["Address"] = supplier.Address
	}

	if City := supplierUpdate["City"].(string); City == "" {
		supplierUpdate["City"] = supplier.City
	}

	if Country := supplierUpdate["Country"].(string); Country == "" {
		supplierUpdate["Country"] = supplier.Country
	}

	if PostalCode := supplierUpdate["PostalCode"].(string); PostalCode == "" {
		supplierUpdate["PostalCode"] = supplier.PostalCode
	}

	if Phone := supplierUpdate["Phone"].(string); Phone == "" {
		supplierUpdate["Phone"] = supplier.Phone
	}

	if HomePage := supplierUpdate["HomePage"].(string); HomePage == "" {
		supplierUpdate["HomePage"] = supplier.HomePage
	}

	result := conn.Model(&supplier).Updates(supplierUpdate)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return supplier, result.Error
	}

	return supplier, result.Error
}

//TODO: Implement DeleteSupplier
func DeleteSupplier(id string) (models.Supplier, error) {
	var supplier models.Supplier

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with supplier delete")

	result := conn.Delete(&supplier, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return supplier, result.Error
	}

	log.Println("Supplier deleted: ", supplier)

	return supplier, result.Error
}
