package entity

type Property struct {
	ID         int
	CategoryID int64
	Name       string
	Type       PropertyType
	Required   bool
	Values     []string
}

type PropertyType string

const (
	AdPropertyTypeRANGE  = "RANGE"
	AdPropertyTypeVALUES = "VALUES"
	AdPropertyTypeBOOL   = "BOOL"
)
