package db

import (
	"log"
	"os"

	"github.com/AlexVonEinzbern/go-rest-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {

	//err := godotenv.Load(".env")

	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

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

	log.Println("Automigration working...")
	db.AutoMigrate(
		&models.Category{},
		&models.Shipper{},
		&models.Supplier{},
		&models.Subcategory{},
		&models.Customer{},
		&models.CreditCard{},
		&models.Login{},
		&models.Order{},
		&models.Payment{},
		&models.Product{},
		&models.OrderProduct{},
		&models.APIError{})
}
