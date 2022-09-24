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

// CreateCustomer godoc
// @Summary Create Customer
// @Description Create a Customer
// @Accept  json
// @Produce  json
// @Tags Customers
// @Param customer body schemas.Customer true "Customer type"
// @Success 200 {object} schemas.CustomerResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /customers [post]
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
// @Tags Customers
// @Success 200 {object} []schemas.CustomerResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /customers [get]
func SearchCustomers(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchCustomers())
}

//TODO: Implement SearchCustomer

// SearchCustomer godoc
// @Summary Search Customer by id
// @Description Search Customer by id in the DataBase
// @Accept  json
// @Produce  json
// @Tags Customers
// @Param id path string true "Customer ID" example(5dd1f36f-1627-4c88-98fb-601feb9634be)
// @Success 200 {object} []schemas.CustomerResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /customers/{id} [get]
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
// @Tags Customers
// @Param id path string true "Customer ID" example(5dd1f36f-1627-4c88-98fb-601feb9634be)
// @Param Customer body schemas.Customer true "Update Customer"
// @Success 204 {object} schemas.CustomerResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /customers/{id} [patch]
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
// @Tags Customers
// @Param id path string true "Customer ID" example(5dd1f36f-1627-4c88-98fb-601feb9634be)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /customers/{id} [delete]
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	customer, err := controllers.DeleteCustomer(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, customer)
}
