package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// CreateSubcategory godoc
// @Summary Create Subcategory
// @Description Create a Subcategory
// @Accept  json
// @Produce  json
// @Tags Subcategories
// @Param subcategory body schemas.Subcategory true "Create a subcategory"
// @Success 200 {object} schemas.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /subcategories [post]
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
// @Tags Subcategories
// @Success 200 {object} []schemas.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /subcategories [get]
func SearchSubcategories(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchSubcategories())
}

// SearchActiveSubcategory godoc
// @Summary Search all active Subcategories
// @Description Search all active categories in the DataBase
// @Accept  json
// @Produce  json
// @Tags Subcategories
// @Success 200 {object} []schemas.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /subcategories/active [get]
func SearchActiveSubcategory(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchSubcategoriesActive())
}

// UpdateSubcategory godoc
// @Summary Update Subcategory by id
// @Description Update a subcategory by id
// @Accept  json
// @Produce  json
// @Tags Subcategories
// @Param id path string true "Subcategory ID" example(2ld1f12f-2227-8s08-18cc-222fdb9634xx)
// @Param Subcategory body schemas.Subcategory true "Update Subcategory"
// @Success 204 {object} schemas.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /subcategories/{id} [patch]
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
// @Tags Subcategories
// @Param id path string true "Subcategory ID" example(2ld1f12f-2227-8s08-18cc-222fdb9634xx)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /subcategories/{id} [delete]
func DeleteSubcategoty(c *gin.Context) {
	id := c.Param("id")

	subcategory, err := controllers.DeleteSubcategory(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, subcategory)
}
