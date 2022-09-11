package controllers

import (
	"log"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/AlexVonEinzbern/go-rest-api/utils"
)

func CreateSubcategory(createsubcategory models.SubcategoryCreate) models.Subcategory {

	subcategory := models.Subcategory{
		ID: utils.IdGenerator(),
		SubcategoryBase: models.SubcategoryBase{
			SubCategoryName: createsubcategory.SubCategoryName,
			Description:     createsubcategory.Description,
			Active:          createsubcategory.Active,
			CategoryID:      createsubcategory.CategoryID}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with subcategory creation")

	result := conn.Create(&subcategory)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", subcategory)

	return subcategory
}

//TODO: Implement SearchSubcategories
func SearchSubcategories() []models.Subcategory {
	var subcategories []models.Subcategory

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with subcategory searching...")

	result := conn.Find(&subcategories)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", subcategories)

	return subcategories
}

//TODO: Implement SearchSubategoriesActive
func SearchSubcategoriesActive() []models.Subcategory {
	var subcategories []models.Subcategory

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Seaching subcategory by active...")

	result := conn.Where("Active = ?", true).Find(&subcategories)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", subcategories)

	return subcategories
}

//TODO: Implement UpdateSubcategory
func UpdateSubcategory(id string, subcategoryUpdate map[string]interface{}) (models.Subcategory, error) {
	var subcategory models.Subcategory

	conn := db.DBConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with subcategory update")

	if err := conn.First(&subcategory, "id = ?", id).Error; err != nil {
		log.Println("DB error ", err)
		return subcategory, err

	}

	if SubCategoryName := subcategoryUpdate["CategoryName"].(string); SubCategoryName == "" {
		subcategoryUpdate["SubCategoryName"] = subcategory.SubCategoryName
	}

	if Description := subcategoryUpdate["Description"].(string); Description == "" {
		subcategoryUpdate["Description"] = subcategory.Description
	}

	if CategoryID := subcategoryUpdate["CategoryID"].(string); CategoryID == "" {
		subcategoryUpdate["CategoryID"] = subcategory.CategoryID
	}

	result := conn.Model(&subcategory).Updates(subcategoryUpdate)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return subcategory, result.Error
	}

	return subcategory, result.Error
}

//TODO: Implement DeleteSubcategory

func DeleteSubcategory(id string) (models.Subcategory, error) {
	var subcategory models.Subcategory

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with subcategory delete")

	result := conn.Delete(&subcategory, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return subcategory, result.Error
	}

	log.Println("Subcategory deleted: ", subcategory)

	return subcategory, result.Error
}
