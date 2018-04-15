package entity

type Component struct {
	ID   uint32 `json:",string"`
	Name string

	ManufacturerID uint16
	Manufacturer   *Manufacturer

	CategoryID int64
	Category   *Category
}
