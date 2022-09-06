package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// CreateSubcategory godoc
// @Summary Create Subcategory
// @Description Create a Subcategory
// @Accept  json
// @Produce  json
// @Param subcategory body models.SubcategoryCreate true "Create a subcategory"
// @Success 200 {object} models.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/subcategories [post]
func CreateSubcategory(c *gin.Context) {

	var createsubcategory models.SubcategoryCreate

	err := c.BindJSON(&createsubcategory)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Subcategory creation: ", createsubcategory)
	c.IndentedJSON(http.StatusOK, controllers.CreateSubcategory(createsubcategory))
}

// SearchAllSubcategories godoc
// @Summary Search all Subcategories
// @Description Search all subcatagories in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/subcategories [get]
func SearchAllSubcategories(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchSubcategories())
}

// SearchActiveSubcategory godoc
// @Summary Search all active Subcategories
// @Description Search all active categories in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/subcategories/active [get]
func SearchActiveSubcategory(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchSubcategoriesActive())
}

// UpdateSubcategory godoc
// @Summary Update Subcategory by id
// @Description Update a subcategory by id
// @Accept  json
// @Produce  json
// @Param id path string true "Subcategory ID" default(cabckbalgaLJHALncas)
// @Success 204 {object} models.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/subcategories/{id} [patch]
func UpdateSubcategory(c *gin.Context) {
	var updatesubcategory models.SubcategoryBase

	id := c.Param("id")

	if err_bind := c.BindJSON(&updatesubcategory); err_bind != nil {
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: err_bind.Error()})
		return
	}

	log.Println("Data to be update: ", updatesubcategory)

	subcategory, err := controllers.UpdateSubcategory(id, structs.Map(updatesubcategory))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, subcategory)
}

// DeleteSubcategory godoc
// @Summary Delete a Subcategory by id
// @Description Delete a subcategory by id
// @Accept  json
// @Produce  json
// @Param id path string true "Subcategory ID" default(cabckbalgaLJHALncas)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/subcategories/{id} [delete]
func DeleteSubcategoty(c *gin.Context) {
	id := c.Param("id")

	subcategory, err := controllers.DeleteSubcategory(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, subcategory)
}
