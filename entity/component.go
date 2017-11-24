package entity

type Component struct {
	ID           uint32
	Name         string
	Manufacturer *Manufacturer
	Category     *Category
}
