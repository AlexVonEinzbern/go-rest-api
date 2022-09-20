package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

//TODO: Implement CreateOrderProduct

// CreateOrderProduct godoc
// @Summary Create OrderProduct
// @Description Create a OrderProduct
// @Accept  json
// @Produce  json
// @Tags OrderProducts
// @Param orderproduct body schemas.OrderProduct true "OrderProduct type"
// @Success 200 {object} schemas.OrderProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /orderproducts [post]
func CreateOrderProduct(c *gin.Context) {

	var createorderproduct models.OrderProductCreate

	err := c.BindJSON(&createorderproduct)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with OrderProduct creation: ", createorderproduct)
	c.IndentedJSON(http.StatusOK, controllers.CreateOrderProduct(createorderproduct))
}

//TODO: Implement SearchOrderProducts

// SearchOrderProducts godoc
// @Summary Search all OrderProducts
// @Description Search all orderproducts in the DataBase
// @Accept  json
// @Produce  json
// @Tags OrderProducts
// @Success 200 {object} []schemas.OrderProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /orderproducts [get]
func SearchOrderProducts(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchOrderProducts())
}

//TODO: Implement SearchOrderProduct

// SearchOrderProduct godoc
// @Summary Search OrderProduct by id
// @Description Search OrderProduct by id in the DataBase
// @Accept  json
// @Produce  json
// @Tags OrderProducts
// @Param id path string true "OrderProduct ID" example(5dd1f36f-1627-4c88-98fb-601feb9634be)
// @Success 200 {object} []schemas.OrderProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /orderproducts/{id} [get]
func SearchOrderProduct(c *gin.Context) {

	id := c.Param("id")

	orderproduct, err := controllers.SearchOrderProduct(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, orderproduct)
}
