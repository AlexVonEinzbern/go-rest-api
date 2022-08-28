package main

import (

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

// @title go-rest-api
// @version 0.1
// @description ecommerce api-rest with basic CRUD operations (Create, Read, Update and Delete) products, suppliers, etc. Manage automatically invoice and shipping operations
// @termsOfService http://swagger.io/terms/

// @contact.name AlexVonEinzbern
// @contact.url https://github.com/AlexVonEinzbern
// @contact.email mario.diaz@correounivalle.edu.com

// @license.name Apache 2.0

func setupRouter() *gin.Engine{
	db.Migrate()
	r := gin.Default()

	//continue...
	
}

func main() {

	db.DBConnection()


}