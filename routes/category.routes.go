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
// @Param subcategory body models.CategoryCreate true "Create a Category"
// @Success 200 {object} models.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/category [post]
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
// @Param ListCategories query array false "List all Categories"
// @Success 200 {object} []models.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go.rest-api/category [get]
func SearchCategories(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchCategories())
}

// SearchCategory godoc
// @Summary Search Category
// @Description Search category by id in the DataBase
// @Accept  json
// @Produce  json
// @Param id path int true "Category id"
// @Success 200 {object} []models.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go.rest-api/category/{id} [get]
func SearchCategory(c *gin.Context) {
	id := c.Param("id")

	category, err := controllers.SearchCategory(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary Update Category by id
// @Description Update a category by id
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body models.CategoryCreate true "Update category"
// @Success 204 {object} models.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/category/{id} [patch]
func UpdateCategory(c *gin.Context) {
	var updatecategory models.CategoryBase

	id := c.Param("id")

	log.Println("Category to be update: ", updatecategory)

	category, err := controllers.UpdateCategory(id, structs.Map(updatecategory))
}
