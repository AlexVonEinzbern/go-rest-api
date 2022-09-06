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
			subcategories.GET("", routes.SearchSubcategories)
			subcategories.GET("/active", routes.SearchActiveSubcategory)
			subcategories.PATCH(":id", routes.UpdateSubcategory)
			subcategories.DELETE(":id", routes.DeleteSubcategoty)
		}
		categories := restapi.Group("/categories")
		{
			categories.POST("", routes.CreateCategory)
			categories.GET("", routes.SearchCategories)
			categories.GET("/active", routes.SearchActiveCategory)
			categories.PATCH(":id", routes.UpdateCategory)
			categories.DELETE(":id", routes.DeleteCategoty)
		}
		products := restapi.Group("/products")
		{
			products.POST("", routes.CreateProduct)
			products.GET("", routes.SearchProducts)
			products.GET("/active", routes.SearchActiveProducts)
			products.PATCH(":id", routes.UpdateProduct)
			products.DELETE(":id", routes.DeleteProduct)
		}
		suppliers := restapi.Group("/suppliers")
		{
			suppliers.POST("", routes.CreateSupplier)
			suppliers.GET("", routes.SearchSuppliers)
			suppliers.PATCH(":id", routes.UpdateSupplier)
			suppliers.DELETE(":id", routes.DeleteSupplier)
		}
		customers := restapi.Group("/customers")
		{
			customers.POST("", routes.CreateCustomer)
			customers.GET("", routes.SearchCustomers)
			customers.GET(":id", routes.SearchCustomer)
			customers.PATCH(":id", routes.UpdateCustomer)
			customers.DELETE(":id", routes.DeleteCustomer)
		}
		orders := restapi.Group("/orders")
		{
			orders.POST("", routes.CreateOrder)
			orders.GET("", routes.SearchOrders)
		}
		shippers := restapi.Group("/shippers")
		{
			shippers.POST("", routes.CreateShipper)
			shippers.GET("", routes.SearchShippers)
			shippers.PATCH(":id", routes.UpdateShipper)
			shippers.DELETE(":id", routes.DeleteShipper)
		}
		orderproducts := restapi.Group("/orderproducts")
		{
			orderproducts.POST("", routes.CreateOrderProduct)
			orderproducts.GET("", routes.SearchOrderProducts)
			orderproducts.GET(":id", routes.SearchOrderProduct)
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
