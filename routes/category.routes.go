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

// CreateCategory godoc
// @Summary Create Category
// @Description Create a Category
// @Accept  json
// @Produce  json
// @Param category body models.CategoryCreate true "Category type"
// @Success 200 {object} models.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/categories [post]
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
// @Summary Search Categories
// @Description Search all catagories in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/categories [get]
func SearchAllCategories(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchCategories())
}

// SearchCategory godoc
// @Summary Search active Categories
// @Description Search active categories in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/categories/active [get]
func SearchActiveCategory(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchCategoriesActive())
}

// UpdateCategory godoc
// @Summary Update Category by id
// @Description Update a category by id
// @Accept  json
// @Produce  json
// @Param id path string true "Category ID" default(cabckbalgaLJHALncas)
// @Success 204 {object} models.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/categories/{id} [patch]
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
// @Param id path string true "Category ID" default(cabckbalgaLJHALncas)
// @Success 204 
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/categories/{id} [delete]
func DeleteCategoty(c *gin.Context) {
	id := c.Param("id")

	category, err := controllers.DeleteCategory(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, category)
}
