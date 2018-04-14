package entity

type Component struct {
	ID   uint32
	Name string

	ManufacturerID uint16
	Manufacturer   *Manufacturer

	CategoryID int64
	Category   *Category
}
