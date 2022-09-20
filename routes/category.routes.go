package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Create Category
// @Description Create a Category
// @Accept  json
// @Produce  json
// @Tags Categories
// @Param category body schemas.Category true "Category type"
// @Success 200 {object} schemas.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /categories [post]
func CreateCategory(c *gin.Context) {

	var createcategory models.CategoryCreate

	err := c.BindJSON(&createcategory)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Subcategory creation: ", createcategory)
	c.IndentedJSON(http.StatusOK, controllers.CreateCategory(createcategory))
}

// SearchCategories godoc
// @Summary Search all Categories
// @Description Search all catagories in the DataBase
// @Accept  json
// @Produce  json
// @Tags Categories
// @Success 200 {object} []schemas.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /categories [get]
func SearchCategories(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchCategories())
}

// SearchCategory godoc
// @Summary Search all active Categories
// @Description Search active categories in the DataBase
// @Accept  json
// @Produce  json
// @Tags Categories
// @Success 200 {object} []schemas.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /categories/active [get]
func SearchActiveCategory(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchCategoriesActive())
}

// UpdateCategory godoc
// @Summary Update Category by id
// @Description Update a category by id
// @Accept  json
// @Produce  json
// @Tags Categories
// @Param id path string true "Category ID" example(5dd1f36f-1627-4c88-98fb-601feb9634be)
// @Success 204 {object} schemas.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /categories/{id} [patch]
func UpdateCategory(c *gin.Context) {
	var updatecategory models.CategoryBase

	id := c.Param("id")

	if err_bind := c.BindJSON(&updatecategory); err_bind != nil {
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: err_bind.Error()})
		return
	}

	log.Println("Data to be update: ", updatecategory)

	category, err := controllers.UpdateCategory(id, structs.Map(updatecategory))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete Category by id
// @Description Delete a category by id
// @Accept  json
// @Produce  json
// @Tags Categories
// @Param id path string true "Category ID" example(5dd1f36f-1627-4c88-98fb-601feb9634be)
// @Success 204 
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /categories/{id} [delete]
func DeleteCategoty(c *gin.Context) {
	id := c.Param("id")

	category, err := controllers.DeleteCategory(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, category)
}
