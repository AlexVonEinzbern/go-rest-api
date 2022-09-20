package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

//TODO: Implement CreatePayment

// CreatePayment godoc
// @Summary Create Payment
// @Description Create a Payment
// @Accept  json
// @Produce  json
// @Tags Payments
// @Param payment body schemas.Payment true "Payment type"
// @Success 200 {object} schemas.PaymentResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /payments [post]
func CreatePayment(c *gin.Context) {

	var createpayment models.PaymentCreate

	err := c.BindJSON(&createpayment)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Payment creation: ", createpayment)
	c.IndentedJSON(http.StatusOK, controllers.CreatePayment(createpayment))
}

//TODO: Implement SearchPayment

// SearchPayment godoc
// @Summary Search Payment by Customer id
// @Description Search Payment by id in the DataBase
// @Accept  json
// @Produce  json
// @Tags Payments
// @Param id path string true "Payment ID" example(2ld1f12f-2227-8s08-18cc-222fdb9634xx)
// @Success 200 {object} []schemas.PaymentResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /payments/{id} [get]
func SearchPayment(c *gin.Context) {

	id := c.Param("id")

	payment, err := controllers.SearchPayment(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, payment)
}

//TODO: Implement SearchPaymentDate

// SearchPaymentDate godoc
// @Summary Search Payments by Date
// @Description Search Payments by date in the DataBase
// @Accept  json
// @Produce  json
// @Tags Payments
// @Param date path string true "Payment date" example(2006-01-02)
// @Success 200 {object} []schemas.PaymentResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /payments/date/{date} [get]
func SearchPaymentDate(c *gin.Context) {

	date := c.Param("date")

	payment, err := controllers.SearchLoginDate(date)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, payment)
}
