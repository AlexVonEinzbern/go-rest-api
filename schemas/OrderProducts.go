package schemas

type OrderProduct struct {
	UnitPrice float32 `example:"32400"`
	Quantity  uint    `example:"5"`
	Discount  float32 `example:"2400"`
	ProductID string  `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	OrderID   string  `example:"9cc1g36t-6537-4d10-11fs-771may9224vv"`
}

type OrderProductResponse struct {
	ID        string  `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	UnitPrice float32 `example:"32400"`
	Quantity  uint    `example:"5"`
	Discount  float32 `example:"2400"`
	ProductID string  `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	OrderID   string  `example:"9cc1g36t-6537-4d10-11fs-771may9224vv"`
}
