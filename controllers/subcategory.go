package controllers

import (
	"log"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
)

func CreateSubcategory(createsubcategory models.SubcategoryCreate) models.Subcategory {

	subcategory := models.Subcategory{
		ID: createsubcategory.ID,
		SubcategoryBase: models.SubcategoryBase{
			SubCategoryName: createsubcategory.SubCategoryName,
			Description:     createsubcategory.Description,
			Active:          createsubcategory.Active}}

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
