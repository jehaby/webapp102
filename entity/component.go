package entity

type Product struct {
	ID   int64 `json:",string"`
	Name string

	BrandID int64
	Brand   *Brand

	CategoryID int64
	Category   *Category
}
