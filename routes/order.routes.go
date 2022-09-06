package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	//"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//TODO: Implement CreateOrder

// @BasePath /api/v1

// CreateOrder godoc
// @Summary Create Order
// @Description Create a Order
// @Accept  json
// @Produce  json
// @Param category body models.OrderCreate true "Order type"
// @Success 200 {object} models.OrderResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/orders [post]
func CreateOrder(c *gin.Context) {

	var createorder models.OrderCreate

	err := c.BindJSON(&createorder)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Order creation: ", createorder)
	c.IndentedJSON(http.StatusOK, controllers.CreateOrder(createorder))
}

//TODO: Implement SearchOrders

// SearchOrders godoc
// @Summary Search all Orders
// @Description Search all orders in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.OrderResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/orders [get]
func SearchOrders(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchOrders())
}

//TODO: Implement UpdateOrder

// UpdateOrder godoc
// @Summary Update Order by id
// @Description Update a order by id
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID" default(cabckbalgaLJHALncas)
// @Success 204 {object} models.OrderResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/orders/{id} [patch]
// func UpdateOrder(c *gin.Context) {
// 	var updateorder models.OrderBase

// 	id := c.Param("id")

// 	if err_bind := c.BindJSON(&updateorder); err_bind != nil {
// 		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: err_bind.Error()})
// 		return
// 	}

// 	log.Println("Data to be update: ", updateorder)

// 	order, err := controllers.UpdateOrder(id, structs.Map(updateorder))

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, order)
// }

//TODO: Implement DeleteOrder

// DeleteOrder godoc
// @Summary Delete Order by id
// @Description Delete a order by id
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID" default(cabckbalgaLJHALncas)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/orders/{id} [delete]
// func DeleteOrder(c *gin.Context) {
// 	id := c.Param("id")

// 	order, err := controllers.DeleteOrder(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusNoContent, order)
// }