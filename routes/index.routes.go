package routes

import (
	"log"
	"net/http"

	"github.com/AlexVonEinzbern/go-rest-api/crud"
	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// CreateSubcategory godoc
// @Summary Create Subcategory
// @Description create a Subcategory
// @Accept  json
// @Produce  json
// @Param subcategory body models.SubcategoryCreate true "Create a subcategory"
// @Success 200 {object} models.SubcategoryResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /go-rest-api/subcategory [post]
func CreateSubcategory(c *gin.Context) {

	var createsubcategory models.SubcategoryCreate

	err := c.BindJSON(&createsubcategory)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with Subcategory creation: ", createsubcategory)
	c.IndentedJSON(http.StatusOK, crud.CreateSubcategory(createsubcategory))
}
