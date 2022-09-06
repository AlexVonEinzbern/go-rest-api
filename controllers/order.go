package controllers

import (
	"log"
	//"time"

	"github.com/AlexVonEinzbern/go-rest-api/db"
	"github.com/AlexVonEinzbern/go-rest-api/models"
)

//TODO: Implement CreateOrder
func CreateOrder(createorder models.OrderCreate) models.Order {

	order := models.Order{
		ID: createorder.ID,
		OrderBase: models.OrderBase{
			OrderDate:      createorder.OrderDate,
			RequiredDate:   createorder.RequiredDate,
			ShippedDate:    createorder.ShippedDate,
			ShipName:       createorder.ShipName,
			ShipAddress:    createorder.ShipAddress,
			ShipCity:       createorder.ShipCity,
			ShipPostalCode: createorder.ShipPostalCode,
			ShipCountry:    createorder.ShipCountry,
			CustomerID:     createorder.CustomerID,
			ShipperID:      createorder.ShipperID}}

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with order creation")

	result := conn.Create(&order)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", order)

	return order
}

//TODO: Implement SearchOrders
func SearchOrders() []models.Order {
	var orders []models.Order

	conn := db.DBConnection()
	sqlconn, err := conn.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlconn.Close()

	log.Println("Starting with orders searching...")

	result := conn.Find(&orders)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", orders)

	return orders
}

//TODO: Implement UpdateOrder
// func UpdateOrder(id string, orderUpdate map[string]interface{}) (models.Order, error) {
// 	var order models.Order

// 	conn := db.DBConnection()
// 	sqlconn, err := conn.DB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sqlconn.Close()

// 	log.Println("Starting with order update")

// 	if err := conn.First(&order, "id = ?", id).Error; err != nil {
// 		log.Println("DB error ", err)
// 		return order, err

// 	}

// 	if OrderDate := orderUpdate["OrderDate"].(time.Time); OrderDate.IsZero() {
// 		orderUpdate["OrderDate"] = order.OrderDate
// 	}

// 	if RequiredDate := orderUpdate["RequiredDate"].(time.Time); RequiredDate.IsZero() {
// 		orderUpdate["RequiredDate"] = order.RequiredDate
// 	}

//     if ShippedDate := orderUpdate["ShippedDate"].(time.Time); ShippedDate.IsZero() {
//         orderUpdate["ShippedDate"] = order.ShippedDate
//     }

//     if ShipName := orderUpdate["ShipName"].(string); ShipName == "" {
//         orderUpdate["ShipName"] = order.ShipName
//     }

//     if ShipAddress := orderUpdate["ShipAddress"].(string); ShipAddress == "" {
//         orderUpdate["ShipAddress"] = order.ShipAddress
//     }

//     if ShipCity := orderUpdate["ShipCity"].(string); ShipCity == "" {
//         orderUpdate["ShipCity"] = order.ShipCity
//     }

//     if ShipPostalCode := orderUpdate["ShipPostalCode"].(string); ShipPostalCode == "" {
//         orderUpdate["ShipPostalCode"] = order.ShipPostalCode
//     }

//     if ShipCountry := orderUpdate["ShipCountry"].(string); ShipCountry == "" {
//         orderUpdate["ShipCountry"] = order.ShipCountry
//     } 

//     if CustomerID := orderUpdate["CustomerID"].(string); CustomerID == "" {
//         orderUpdate["CustomerID"] = order.CustomerID
//     }

//     if ShipperID := orderUpdate["ShipperID"].(string); ShipperID == "" {
//         orderUpdate["ShipperID"] = order.ShipperID
//     }

// 	result := conn.Model(&order).Updates(orderUpdate)

// 	if result.Error != nil {
// 		log.Println("DB error ", result.Error)
// 		return order, result.Error
// 	}

// 	return order, result.Error
// }

//TODO: Implement DeleteOrder
// func DeleteOrder(id string) (models.Order, error) {
// 	var order models.Order

// 	conn := db.DBConnection()
// 	sqlconn, err := conn.DB()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer sqlconn.Close()

// 	log.Println("Starting with order delete")

// 	result := conn.Delete(&order, "id = ?", id)

// 	if result.Error != nil {
// 		log.Println("DB error", result.Error)
// 		return order, result.Error
// 	}

// 	log.Println("Order deleted: ", order)

// 	return order, result.Error
// }