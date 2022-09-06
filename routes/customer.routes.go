package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//TODO: Implement CreateCustomer

// @BasePath /api/v1

// CreateCustomer godoc
// @Summary Create Customer
// @Description Create a Customer
// @Accept  json
// @Produce  json
// @Param category body models.CustomerCreate true "Customer type"
// @Success 200 {object} models.CustomerResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/customers [post]
func CreateCustomer(c *gin.Context) {

	var createcustomer models.CustomerCreate

	err := c.BindJSON(&createcustomer)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Customer creation: ", createcustomer)
	c.IndentedJSON(http.StatusOK, controllers.CreateCustomer(createcustomer))
}

//TODO: Implement SearchCustomers

// SearchAllCustomers godoc
// @Summary Search all Customers
// @Description Search all customers in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.CustomerResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/customers [get]
func SearchCustomers(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchCustomers())
}

//TODO: Implement SearchCustomer
// SearchCustomer godoc
// @Summary Search Customer by id
// @Description Search Customer by id in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.CategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/categories/active [get]
func SearchCustomer(c *gin.Context) {

    id := c.Param("id")

    customer, err := controllers.SearchCustomer(id)

    if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, customer)
}

//TODO: Implement UpdateCustomer

// UpdateCustomer godoc
// @Summary Update Customer by id
// @Description Update a customer by id
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID" default(cabckbalgaLJHALncas)
// @Success 204 {object} models.CustomerResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/customers/{id} [patch]
func UpdateCustomer(c *gin.Context) {
	var updatecustomer models.CustomerBase

	id := c.Param("id")

	if err_bind := c.BindJSON(&updatecustomer); err_bind != nil {
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: err_bind.Error()})
		return
	}

	log.Println("Data to be update: ", updatecustomer)

	customer, err := controllers.UpdateCustomer(id, structs.Map(updatecustomer))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, customer)
}

//TODO: Implement DeleteCustomer

// DeleteCustomer godoc
// @Summary Delete Customer by id
// @Description Delete a customer by id
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID" default(cabckbalgaLJHALncas)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/customers/{id} [delete]
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	customer, err := controllers.DeleteCustomer(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, customer)
}
