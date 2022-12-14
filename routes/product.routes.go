package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Create Product
// @Description Create a Product
// @Accept  json
// @Produce  json
// @Tags Products
// @Param product body schemas.Product true "Product type"
// @Success 200 {object} schemas.ProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /products [post]
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
// @Tags Products
// @Success 200 {object} []schemas.ProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /products [get]
func SearchProducts(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchProducts())
}

// SearchActiveProducts godoc
// @Summary Search all active Products
// @Description Search all active products in the DataBase
// @Accept  json
// @Produce  json
// @Tags Products
// @Success 200 {object} []schemas.ProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /products/active [get]
func SearchActiveProducts(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchProductsActive())
}

// UpdateProduct godoc
// @Summary Update Product by id
// @Description Update a product by id
// @Accept  json
// @Produce  json
// @Tags Products
// @Param id path string true "Product ID" example(2ld1f12f-2227-8s08-18cc-222fdb9634x)
// @Param Product body schemas.Product true "Update Product"
// @Success 204 {object} schemas.ProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /products/{id} [patch]
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
// @Tags Products
// @Param id path string true "Product ID" example(2ld1f12f-2227-8s08-18cc-222fdb9634x)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := controllers.DeleteProduct(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, product)
}
