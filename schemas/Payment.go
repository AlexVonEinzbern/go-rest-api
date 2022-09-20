package schemas

type Payment struct {
	OrderID      string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	CreditCardID string `example:"5cc1f36f-1287-4c88-63fb-601feb9634be"`
}

type PaymentResponse struct {
	ID           string `example:"2ld1f12f-2227-8s08-18cc-222fdb9634xx"`
	OrderID      string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	CreditCardID string `example:"5cc1f36f-1287-4c88-63fb-601feb9634be"`
}
