package db

import (
	"log"
	"os"

	"github.com/AlexVonEinzbern/go-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	var error error

	Host := os.Getenv("DB_HOST")
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASSWORD")
	Dbname := os.Getenv("DB_NAME")

	DSN := "host=" + Host + " user=" + User + " password=" + Password + " dbname=" + Dbname + " port=5432 sslmode=disable"

	db, error := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}

	return db
}

func Migrate() {
	db := DBConnection()
	sqlDB, error := db.DB()

	if error != nil {
		log.Fatal(error)
	}

	defer sqlDB.Close()

	log.Println("Automigration working....")
	db.AutoMigrate(
		&models.Category{},
		&models.CreditCard{},
		&models.Customers{},
		&models.Login{},
		&models.OrderProduct{},
		&models.Orders{},
		&models.Payment{},
		&models.Products{},
		&models.Shippers{},
		&models.Subcategory{},
		&models.Suppliers{})
}
