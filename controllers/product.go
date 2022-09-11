package controllers

import (
	"log"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/AlexVonEinzbern/go-rest-api/utils"
)

//TODO: Implement CreateProduct
func CreateProduct(createproduct models.ProductCreate) models.Product {

	product := models.Product{
		ID: utils.IdGenerator(),
		ProductBase: models.ProductBase{
			ProductName:     createproduct.ProductName,
			QuantityPerUnit: createproduct.QuantityPerUnit,
			UnitsInStock:    createproduct.UnitsInStock,
			UnitsOnOrder:    createproduct.UnitsOnOrder,
			ReorderLevel:    createproduct.ReorderLevel,
			Discontinued:    createproduct.Discontinued,
			Quantity:        createproduct.Quantity,
			SupplierID:      createproduct.SupplierID,
			CategoryID:      createproduct.CategoryID}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with product creation")

	result := conn.Create(&product)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", product)

	return product
}

//TODO: Implement SearchProducts
func SearchProducts() []models.Product {
	var products []models.Product

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with products searching...")

	result := conn.Find(&products)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", products)

	return products
}

//TODO: Implement SearchProductsActive
func SearchProductsActive() []models.Product {
	var products []models.Product

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Seaching products by active...")

	result := conn.Where("Active = ?", true).Find(&products)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", products)

	return products
}

//TODO: Implement UpdateProduct
func UpdateProduct(id string, productUpdate map[string]interface{}) (models.Product, error) {
	var product models.Product

	conn := db.DBConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with product update")

	if err := conn.First(&product, "id = ?", id).Error; err != nil {
		log.Println("DB error ", err)
		return product, err

	}

	if ProductName := productUpdate["ProductName"].(string); ProductName == "" {
		productUpdate["ProductName"] = product.ProductName
	}

	if QuantityPerUnit := productUpdate["QuantityPerUnit"].(string); QuantityPerUnit == "" {
		productUpdate["QuantityPerUnit"] = product.QuantityPerUnit
	}

	if UnitsInStock := productUpdate["UnitsInStock"].(uint); UnitsInStock == 0 {
		productUpdate["UnitsInStock"] = product.UnitsInStock
	}

	if UnitsOnOrder := productUpdate["UnitsOnOrder"].(uint); UnitsOnOrder == 0 {
		productUpdate["UnitsOnOrder"] = product.UnitsOnOrder
	}

	if ReorderLevel := productUpdate["ReorderLevel"].(uint); ReorderLevel == 0 {
		productUpdate["ReorderLevel"] = product.ReorderLevel
	}

	if Quantity := productUpdate["Quantity"].(uint); Quantity == 0 {
		productUpdate["Quantity"] = product.Quantity
	}

	if SupplierID := productUpdate["SupplierID"].(string); SupplierID == "" {
		productUpdate["SupplierID"] = product.SupplierID
	}

	if CategoryID := productUpdate["CategoryID"].(string); CategoryID == "" {
		productUpdate["CategoryID"] = product.CategoryID
	}

	result := conn.Model(&product).Updates(productUpdate)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return product, result.Error
	}

	return product, result.Error
}

//TODO: Implement DeleteProduct
func DeleteProduct(id string) (models.Product, error) {
	var product models.Product

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with product delete")

	result := conn.Delete(&product, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return product, result.Error
	}

	log.Println("Product deleted: ", product)

	return product, result.Error
}
