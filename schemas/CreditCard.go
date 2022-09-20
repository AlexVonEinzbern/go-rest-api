package schemas

type CreditCard struct {
	Brand      string `example:"Visa"`
	Number     string `example:"4468467578025601"`
	CVV        string `example:"847"`
	MM         string `example:"07"`
	YYYY       string `example:"2025"`
	CustomerID string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
}

type CreditCardResponse struct {
	ID         string `example:"5cc1f36f-1287-4c88-63fb-601feb9634be"`
	Brand      string `example:"Visa"`
	Number     string `example:"4468467578025601"`
	CVV        string `example:"847"`
	MM         string `example:"07"`
	YYYY       string `example:"2025"`
	CustomerID string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
}
