package schemas

import "time"

type Order struct {
	OrderDate      time.Time `example:"1990-01-02"`
	RequiredDate   time.Time `example:"1990-01-02"`
	ShippedDate    time.Time `example:"1990-01-04"`
	ShipName       string    `example:"Servientrega"`
	ShipAddress    string    `example:"6597 Westheimer Rd"`
	ShipCity       string    `example:"Billings"`
	ShipPostalCode string    `example:"63104"`
	ShipCountry    string    `example:"United States"`
	CustomerID     string    `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	ShipperID      string    `example:"9cc1g36t-6537-4d10-11fs-771may9224vv"`
}

type OrderResponse struct {
	ID             string    `example:"2ld1f12f-2227-8s08-18cc-222fdb9634xx"`
	OrderDate      time.Time `example:"1990-01-02"`
	RequiredDate   time.Time `example:"1990-01-02"`
	ShippedDate    time.Time `example:"1990-01-04"`
	ShipName       string    `example:"not null"`
	ShipAddress    string    `example:"6597 Westheimer Rd"`
	ShipCity       string    `example:"Billings"`
	ShipPostalCode string    `example:"63104"`
	ShipCountry    string    `example:"United States"`
	CustomerID     string    `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	ShipperID      string    `example:"9cc1g36t-6537-4d10-11fs-771may9224vv"`
}
