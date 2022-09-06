package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/controllers"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

//TODO: Implement CreateSupplier

// @BasePath /api/v1

// CreateSupplier godoc
// @Summary Create Supplier
// @Description Create a Supplier
// @Accept  json
// @Produce  json
// @Param category body models.SupplierCreate true "Supplier type"
// @Success 200 {object} models.SupplierResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/suppliers [post]
func CreateSupplier(c *gin.Context) {

	var createsupplier models.SupplierCreate

	err := c.BindJSON(&createsupplier)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Subcategory creation: ", createsupplier)
	c.IndentedJSON(http.StatusOK, controllers.CreateSupplier(createsupplier))
}

//TODO: Implement SearchAllSuppliers

// SearchAllSuppliers godoc
// @Summary Search all Suppliers
// @Description Search all suppliers in the DataBase
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.SupplierResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/suppliers [get]
func SearchAllSuppliers(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchAllSuppliers())
}

//TODO: Implement UpdateSupplier

// UpdateSupplier godoc
// @Summary Update Supplier by id
// @Description Update a supplier by id
// @Accept  json
// @Produce  json
// @Param id path string true "Supplier ID" default(cabckbalgaLJHALncas)
// @Success 204 {object} models.SupplierResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/suppliers/{id} [patch]
func UpdateSupplier(c *gin.Context) {
	var updatesupplier models.SupplierBase

	id := c.Param("id")

	if err_bind := c.BindJSON(&updatesupplier); err_bind != nil {
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: err_bind.Error()})
		return
	}

	log.Println("Data to be update: ", updatesupplier)

	supplier, err := controllers.UpdateSupplier(id, structs.Map(updatesupplier))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, supplier)
}

//TODO: Implement DeleteSupplier

// DeleteSupplier godoc
// @Summary Delete Supplier by id
// @Description Delete a supplier by id
// @Accept  json
// @Produce  json
// @Param id path string true "Supplier ID" default(cabckbalgaLJHALncas)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/suppliers/{id} [delete]
func DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	supplier, err := controllers.DeleteSupplier(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, supplier)
}
