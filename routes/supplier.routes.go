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

// CreateSupplier godoc
// @Summary Create Supplier
// @Description Create a Supplier
// @Accept  json
// @Produce  json
// @Tags Suppliers
// @Param supplier body schemas.Supplier true "Supplier type"
// @Success 200 {object} schemas.SupplierResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /suppliers [post]
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
// @Tags Suppliers
// @Success 200 {object} []schemas.SupplierResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /suppliers [get]
func SearchSuppliers(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, controllers.SearchSuppliers())
}

//TODO: Implement UpdateSupplier

// UpdateSupplier godoc
// @Summary Update Supplier by id
// @Description Update a supplier by id
// @Accept  json
// @Produce  json
// @Tags Suppliers
// @Param id path string true "Supplier ID" example(2ld1f12f-2227-8s08-18cc-222fdb9634x)
// @Success 204 {object} schemas.SupplierResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /suppliers/{id} [patch]
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
// @Tags Suppliers
// @Param id path string true "Supplier ID" example(2ld1f12f-2227-8s08-18cc-222fdb9634x)
// @Success 204
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /suppliers/{id} [delete]
func DeleteSupplier(c *gin.Context) {
	id := c.Param("id")

	supplier, err := controllers.DeleteSupplier(id)

	if err != nil {
		c.IndentedJSON(http.StatusOK, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, supplier)
}
