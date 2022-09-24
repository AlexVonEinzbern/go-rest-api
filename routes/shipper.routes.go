package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//TODO: Implement CreateShipper

// CreateShipper godoc
// @Summary Create Shipper
// @Description Create a Shipper
// @Accept  json
// @Produce  json
// @Tags Shippers
// @Param shipper body schemas.Shipper true "Shipper type"
// @Success 200 {object} schemas.ShipperResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /shippers [post]
func CreateShipper(c *gin.Context) {

	var createshipper models.ShipperCreate

	err := c.BindJSON(&createshipper)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Shipper creation: ", createshipper)
	c.IndentedJSON(http.StatusOK, controllers.CreateShipper(createshipper))
}

//TODO: Implement SearchShippers

// SearchShippers godoc
// @Summary Search all Shippers
// @Description Search all shippers in the DataBase
// @Accept  json
// @Produce  json
// @Tags Shippers
// @Success 200 {object} []schemas.ShipperResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /shippers [get]
func SearchShippers(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchShippers())
}

//TODO: Implement UpdateShippers

// UpdateShipper godoc
// @Summary Update Shipper by id
// @Description Update a shipper by id
// @Accept  json
// @Produce  json
// @Tags Shippers
// @Param id path string true "Shipper ID" example(5dd1f36f-1627-4c88-98fb-601feb9634be)
// @Param Shipper body schemas.Shipper true "Update Shipper"
// @Success 204 {object} schemas.ShipperResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /shippers/{id} [patch]
func UpdateShipper(c *gin.Context) {
	var updateshipper models.ShipperBase

	id := c.Param("id")

	if err_bind := c.BindJSON(&updateshipper); err_bind != nil {
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: err_bind.Error()})
		return
	}

	log.Println("Data to be update: ", updateshipper)

	shipper, err := controllers.UpdateShipper(id, structs.Map(updateshipper))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, shipper)
}

//TODO: Implement DeleteShipper

// DeleteShipper godoc
// @Summary Delete Shipper by id
// @Description Delete a shipper by id
// @Accept  json
// @Produce  json
// @Tags Shippers
// @Param id path string true "Shipper ID" example(5dd1f36f-1627-4c88-98fb-601feb9634be)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /shippers/{id} [delete]
func DeleteShipper(c *gin.Context) {
	id := c.Param("id")

	shipper, err := controllers.DeleteShipper(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, shipper)
}
