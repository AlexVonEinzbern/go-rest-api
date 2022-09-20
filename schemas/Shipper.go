package schemas

type Shipper struct {
	CompanyName string `example:"Servientrega"`
	Phone       string `example:"6697901"`
}

type ShipperResponse struct {
	ID          string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	CompanyName string `example:"Servientrega"`
	Phone       string `example:"6697901"`
}
