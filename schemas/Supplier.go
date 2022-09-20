package schemas

type Supplier struct {
	CompanyName string `example:"Arthur Clark"`
	Address     string `example:"49 Spring St"`
	City        string `example:"Billings"`
	Country     string `example:"United States"`
	PostalCode  string `example:"63104"`
	Phone       string `example:"6475240"`
	HomePage    string `example:"arthur.clark.com"`
}

type SupplierResponse struct {
	ID          string `example:"2ld1f12f-2227-8s08-18cc-222fdb9634xx"`
	CompanyName string `example:"Arthur Clark"`
	Address     string `example:"49 Spring St"`
	City        string `example:"Billings"`
	Country     string `example:"United States"`
	PostalCode  string `example:"63104"`
	Phone       string `example:"6475240"`
	HomePage    string `example:"arthur.clark.com"`
}
