package entity

type Product struct {
	ID   int64 `json:",string"`
	Name string

	ManufacturerID int64
	Manufacturer   *Manufacturer

	CategoryID int64
	Category   *Category
}
