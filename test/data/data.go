package data

import (
	"github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/entity"
)

var TestUser = entity.User{
	UUID:     uuid.FromStringOrNil("e12087ab-23b9-4d97-8b61-e7016e4e956b"),
	Name:     "urf",
	Email:    "u@j.com",
	Password: "$2a$10$R2iIpKeBPb12wcF3cZnzDuzlWKbM4fyFQo01S2d5eiNEXMO.8t7cS",
}

var Brands = struct {
	Shimano, SRAM, Merida int64
}{
	Shimano: 1,
	SRAM:    2,
	Merida:  3,
}

var Categories = struct {
	Chain, Fork, Cassette, FrameHardtailXC int64
}{
	Chain:           1401,
	Cassette:        1461,
	FrameHardtailXC: 641,
}

var Localities = struct {
	Moscow, SPB int64
}{
	Moscow: 1,
	SPB:    2,
}
