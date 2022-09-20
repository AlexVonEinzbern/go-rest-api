package schemas

type Category struct {
	CategoryName string `example:"Category 1"`
	Description  string `example:"Description 1"`
}

type CategoryResponse struct {
	ID           string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
	CategoryName string `example:"Category 1"`
	Description  string `example:"Description 1"`
	Active       bool   `example:"false"`
}