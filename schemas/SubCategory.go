package schemas

type Subcategory struct {
	SubCategoryName string `example:"Subcategory 1"`
	Description     string `example:"Description about subcategory 1"`
	CategoryID      string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
}

type SubcategoryResponse struct {
	ID              string `example:"2ld1f12f-2227-8s08-18cc-222fdb9634xx"`
	SubCategoryName string `example:"Subcategory 1"`
	Description     string `example:"Description about subcategory 1"`
	CategoryID      string `example:"5dd1f36f-1627-4c88-98fb-601feb9634be"`
}
