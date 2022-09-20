package schemas

type Product struct {
	ProductName     string `example:"Product 1"`
	QuantityPerUnit string `example:"2"`
	UnitsInStock    uint   `example:"10"`
	UnitsOnOrder    uint   `example:"3"`
	ReorderLevel    uint   `example:"1"`
	Discontinued    bool   `example:"false"`
	Quantity        uint   `example:"10"`
	SupplierID      string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	CategoryID      string `example:"5cc1f36f-1287-4c88-63fb-601feb9634be"`
}

type ProductResponse struct {
	ID              string `example:"2ld1f12f-2227-8s08-18cc-222fdb9634xx"`
	ProductName     string `example:"Product 1"`
	QuantityPerUnit string `example:"32400"`
	UnitsInStock    uint   `example:"10"`
	UnitsOnOrder    uint   `example:"3"`
	ReorderLevel    uint   `example:"1"`
	Discontinued    bool   `example:"false"`
	Quantity        uint   `example:"10"`
	SupplierID      string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	CategoryID      string `example:"5cc1f36f-1287-4c88-63fb-601feb9634be"`
}
