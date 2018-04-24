package test

import (
	"github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/entity"
)

var testUser = entity.User{
	UUID:     uuid.FromStringOrNil("e12087ab-23b9-4d97-8b61-e7016e4e956b"),
	Name:     "urf",
	Email:    "u@j.com",
	Password: "$2a$10$R2iIpKeBPb12wcF3cZnzDuzlWKbM4fyFQo01S2d5eiNEXMO.8t7cS",
}

var brands = struct {
	Shimano, SRAM int64
}{
	Shimano: 1,
	SRAM:    2,
}
