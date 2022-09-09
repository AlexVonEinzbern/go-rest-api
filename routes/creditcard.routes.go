package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//TODO: Implement CreateCreditCard

// @BasePath /api/v1

// CreateCreditCard godoc
// @Summary Create creditcard
// @Description Create a CreditCard
// @Accept  json
// @Produce  json
// @Param creditcard body models.CreditCardCreate true "CreditCard a subcategory"
// @Success 200 {object} models.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/creditcard [post]
func CreateCreditCard(c *gin.Context) {

	var createcreditcard models.CreditCardCreate

	err := c.BindJSON(&createcreditcard)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with CreditCard creation: ", createcreditcard)
	c.IndentedJSON(http.StatusOK, controllers.CreateCreditCard(createcreditcard))
}

//TODO: Implement SearchCreditCards

// SearchCreditCards godoc
// @Summary Search all SearchCreditCards
// @Description Search all credit cards in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.CreditCardResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/creditcard [get]
func SearchCreditCards(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchCreditCards())
}

//TODO: Implement SearchCreditCard

// SearchCreditCard godoc
// @Summary Search CreditCard by id
// @Description Search CreditCard by id in the DataBase
// @Accept  json
// @Produce  json
// @Param id path string true "CreditCard ID" default(cabckbalgaLJHALncas)
// @Success 200 {object} []models.CreditCardResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/creditcard/{id} [get]
func SearchCreditCard(c *gin.Context) {

    id := c.Param("id")

    creditcard, err := controllers.SearchCreditCard(id)

    if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, creditcard)
}

//TODO: Implement UpdateCreditCard

// UpdateCreditCard godoc
// @Summary Update CreditCard by id
// @Description Update a credit card by id
// @Accept  json
// @Produce  json
// @Param id path string true "CreditCard ID" default(cabckbalgaLJHALncas)
// @Success 204 {object} models.CreditCardResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/creditcard/{id} [patch]
func UpdateCreditCard(c *gin.Context) {
	var updatecreditcard models.CreditCardBase

	id := c.Param("id")

	if err_bind := c.BindJSON(&updatecreditcard); err_bind != nil {
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: err_bind.Error()})
		return
	}

	log.Println("Data to be update: ", updatecreditcard)

	creditcard, err := controllers.UpdateCreditCard(id, structs.Map(updatecreditcard))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, creditcard)
}

//TODO: Implement DeleteCreditCard

// DeleteCreditCard godoc
// @Summary Delete CreditCard by id
// @Description Delete a credit card by id
// @Accept  json
// @Produce  json
// @Param id path string true "CreditCard ID" default(cabckbalgaLJHALncas)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/creditcard/{id} [delete]
func DeleteCreditCard(c *gin.Context) {
	id := c.Param("id")

	creditcard, err := controllers.DeleteCreditCard(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, creditcard)
}