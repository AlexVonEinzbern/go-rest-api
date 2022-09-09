package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

//TODO: Implement CreateOrderProduct

// @BasePath /api/v1

// CreateOrderProduct godoc
// @Summary Create OrderProduct
// @Description Create a OrderProduct
// @Accept  json
// @Produce  json
// @Param orderproduct body models.OrderProductCreate true "OrderProduct type"
// @Success 200 {object} models.OrderProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/orderproducts [post]
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
// @Success 200 {object} []models.OrderProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/orderproducts [get]
func SearchOrderProducts(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchOrderProducts())
}

//TODO: Implement SearchOrderProduct

// SearchOrderProduct godoc
// @Summary Search OrderProduct by id
// @Description Search OrderProduct by id in the DataBase
// @Accept  json
// @Produce  json
// @Param id path string true "OrderProduct ID" default(cabckbalgaLJHALncas)
// @Success 200 {object} []models.OrderProductResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/orderproducts/{id} [get]
func SearchOrderProduct(c *gin.Context) {

	id := c.Param("id")

	orderproduct, err := controllers.SearchOrderProduct(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, orderproduct)
}
