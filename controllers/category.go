package controllers

import (
	"log"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
)

func CreateCategory(createcategory models.CategoryCreate) models.Category {

	category := models.Category{
		ID: createcategory.ID,
		CategoryBase: models.CategoryBase{
			CategoryName: createcategory.CategoryName,
			Description:  createcategory.Description,
			Active:       createcategory.Active}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with category creation")

	result := conn.Create(&category)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", category)

	return category
}

func SearchCategories() []models.Category {
	var categories []models.Category

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with category searching...")

	result := conn.Find(&categories)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", categories)

	return categories
}

func SearchCategory(id string) (models.Category, error) {
	var category models.Category

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Seaching category by id...")

	result := conn.First(&category, "id = ?", id)

	if result.Error != nil {
		log.Fatal(result.Error)
		return category, result.Error
	}

	log.Println("Found record: ", category)

	return category, result.Error
}
