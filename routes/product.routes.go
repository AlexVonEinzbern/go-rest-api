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

// CreateProduct godoc
// @Summary Create Product
// @Description Create a Product
// @Accept  json
// @Produce  json
// @Param product body models.ProductCreate true "Product type"
// @Success 200 {object} models.ProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/products [post]
func CreateProduct(c *gin.Context) {

	var createproduct models.ProductCreate

	err := c.BindJSON(&createproduct)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Product creation: ", createproduct)
	c.IndentedJSON(http.StatusOK, controllers.CreateProduct(createproduct))
}

// SearchAllProducts godoc
// @Summary Search all Products
// @Description Search all products in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.ProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/products [get]
func SearchAllProducts(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchProducts())
}

// SearchActiveProducts godoc
// @Summary Search all active Products
// @Description Search all active products in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.ProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/products/active [get]
func SearchActiveProducts(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchProductsActive())
}

// UpdateProduct godoc
// @Summary Update Product by id
// @Description Update a product by id
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" default(cabckbalgaLJHALncas)
// @Success 204 {object} models.ProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/products/{id} [patch]
func UpdateProduct(c *gin.Context) {
	var updateproduct models.ProductBase

	id := c.Param("id")

	if err_bind := c.BindJSON(&updateproduct); err_bind != nil {
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: err_bind.Error()})
		return
	}

	log.Println("Data to be update: ", updateproduct)

	product, err := controllers.UpdateProduct(id, structs.Map(updateproduct))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a Product by id
// @Description Delete a product by id
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" default(cabckbalgaLJHALncas)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := controllers.DeleteProduct(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, product)
}
