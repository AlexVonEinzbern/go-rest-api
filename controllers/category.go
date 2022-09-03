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

func SearchCategoriesActive() []models.Category {
	var categories []models.Category

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Seaching category by active...")

	result := conn.Where("Active = ?", true).Find(&categories)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", categories)

	return categories
}

func UpdateCategory(id string, categoryUpdate map[string]interface{}) (models.Category, error) {
	var category models.Category

	conn := db.DBConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with category update")

	if err := conn.First(&category, "id = ?", id).Error; err != nil {
		log.Println("DB error ", err)
		return category, err

	}

	if CategoryName := categoryUpdate["CategoryName"].(string); CategoryName == "" {
		categoryUpdate["CategoryName"] = category.CategoryName
	}

	if Description := categoryUpdate["Description"].(string); Description == "" {
		categoryUpdate["Description"] = category.Description
	}

	result := conn.Model(&category).Updates(categoryUpdate)
	//.Updates(characterUpdate)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return category, result.Error
	}

	return category, result.Error
}

func DeleteCategory(id string) (models.Category, error) {
	var category models.Category

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with category delete")

	result := conn.Delete(&category, "id = ?", id)

	if result.Error != nil {
		log.Println("DB error", result.Error)
		return category, result.Error
	}

	log.Println("Category deleted: ", category)

	return category, result.Error
}
