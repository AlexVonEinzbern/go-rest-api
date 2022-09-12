package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

//TODO: Implement CreateLogin

// @BasePath /api/v1

// CreateLogin godoc
// @Summary Create Login
// @Description Create a Login
// @Accept  json
// @Produce  json
// @Param login body models.LoginCreate true "Login type"
// @Success 200 {object} models.LoginResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/login [post]
func CreateLogin(c *gin.Context) {

	var createlogin models.LoginCreate

	err := c.BindJSON(&createlogin)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Login creation: ", createlogin)
	c.IndentedJSON(http.StatusOK, controllers.CreateLogin(createlogin))
}

//TODO: Implement SearchLogin

// SearchLogin godoc
// @Summary Search Login by Customer id
// @Description Search Login by customer id in the DataBase
// @Accept  json
// @Produce  json
// @Param id path string true "Login ID" default(cabckbalgaLJHALncas)
// @Success 200 {object} []models.LoginResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/login/{id} [get]
func SearchLogin(c *gin.Context) {

	id := c.Param("id")

	login, err := controllers.SearchLogin(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, login)
}

//TODO: Implement SearchLoginDate

// SearchLoginDate godoc
// @Summary Search Login by Date
// @Description Search Login by date in the DataBase
// @Accept  json
// @Produce  json
// @Param date path string true "Login date" default(2006-01-02)
// @Success 200 {object} []models.LoginResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/login/date/{date} [get]
func SearchLoginDate(c *gin.Context) {

	date := c.Param("date")

	login, err := controllers.SearchLoginDate(date)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, login)
}