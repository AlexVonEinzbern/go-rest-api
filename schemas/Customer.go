package schemas

import "time"

type Customer struct {
	Name       string    `example:"Ramona Ellis"`
	Email      string    `example:"ramona.ellis@example.com"`
	DNI        string    `example:"105883636"`
	Address    string    `example:"6597 Westheimer Rd"`
	City       string    `example:"Billings"`
	Birthday   time.Time `example:"1990-01-02"`
	PostalCode string    `example:"63104"`
	Country    string    `example:"United States"`
	Phone      string    `example:"7097900"`
	UserName   string    `example:"yellowpeacock117"`
	Password   string    `example:"addison"`
}

type CustomerResponse struct {
	ID         string    `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	Name       string    `example:"Ramona Ellis"`
	Email      string    `example:"ramona.ellis@example.com"`
	DNI        string    `example:"105883636"`
	Address    string    `example:"6597 Westheimer Rd"`
	City       string    `example:"Billings"`
	Birthday   time.Time `example:"1990-01-02"`
	PostalCode string    `example:"63104"`
	Country    string    `example:"United States"`
	Phone      string    `example:"7097900"`
	UserName   string    `example:"yellowpeacock117"`
	Password   string    `example:"addison"`
}
