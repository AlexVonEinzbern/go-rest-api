package main

import (
	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/AlexVonEinzbern/go-rest-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title go-rest-api
// @version 1
// @description ecommerce api-rest with basic CRUD operations (Create, Read, Update and Delete) products, suppliers, etc. Manage automatically invoice and shipping operations
// @termsOfService http://swagger.io/terms/

// @contact.name AlexVonEinzbern
// @contact.url https://github.com/AlexVonEinzbern
// @contact.email mario.diaz@correounivalle.edu.co

// @license.name Apache 2.0

func setupRouter() *gin.Engine {
	db.Migrate()
	r := gin.Default()
	restapi := r.Group("/go-rest-api")
	{
		subcategories := restapi.Group("/subcategories")
		{
			subcategories.POST("", routes.CreateSubcategory)
		}
		categories := restapi.Group("/categories")
		{
			categories.POST("", routes.CreateCategory)
			categories.GET("", routes.SearchAllCategories)
			categories.GET("/active", routes.SearchActiveCategory)
			categories.PATCH(":id", routes.UpdateCategory)
			categories.DELETE(":id", routes.DeleteCategoty)
		}
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") //The URL pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
