package data

import (
	"github.com/jehaby/webapp102/entity"
	uuid "github.com/satori/go.uuid"
)

var Ads = []*entity.Ad{
	&entity.Ad{
		UUID:        uuid.FromStringOrNil("576c3184-eb8b-4558-9dca-7b4727b224e6"),
		Name:        "Shimano XT CS-M771-10 Cassette",
		Description: "in good condition!",
		UserUUID:    TestUser.UUID,
		CategoryID:  Categories.Cassette,
		BrandID:     Brands.Shimano,
		Condition:   entity.ConditionUsed,
		LocalityID:  Localities.Moscow,
		Price:       90000,
		Currency:    entity.CurrencyRUB,
		Weight:      567,
		Properties:  `{"speeds":"10sp","material":"steel"}`,
	},
	&entity.Ad{
		UUID:        uuid.FromStringOrNil("4cb05066-bb3d-4c55-8b0e-303f13c0fed0"),
		Name:        "Shimano Kassette CS-HG50-7",
		Description: "still works!",
		UserUUID:    TestUser.UUID,
		CategoryID:  Categories.Cassette,
		BrandID:     Brands.Shimano,
		Condition:   entity.ConditionUsed,
		LocalityID:  Localities.Moscow,
		Price:       30000,
		Currency:    entity.CurrencyRUB,
		Weight:      620,
		Properties:  `{"speeds":"7sp","material":"steel"}`,
	},
	&entity.Ad{
		UUID:        uuid.FromStringOrNil("7c97e952-d5fd-44c8-b3bc-cd8876f4899b"),
		Name:        "SRAM PG-970 9-fach Kassette",
		Description: "new!!!!",
		UserUUID:    TestUser.UUID,
		CategoryID:  Categories.Cassette,
		BrandID:     Brands.SRAM,
		Condition:   entity.ConditionNew,
		LocalityID:  Localities.Moscow,
		Price:       130000,
		Currency:    entity.CurrencyRUB,
		Weight:      350,
		Properties:  `{"speeds":"9sp","material":"aluminium"}`,
	},
	&entity.Ad{
		UUID:        uuid.FromStringOrNil("a38ea3ed-95ff-4509-9085-9676a16db152"),
		Name:        "SRAM PC 991 9-speed Chain",
		Description: "could be better",
		UserUUID:    TestUser.UUID,
		CategoryID:  Categories.Chain,
		BrandID:     Brands.SRAM,
		Condition:   entity.ConditionUsed,
		LocalityID:  Localities.SPB,
		Price:       25000,
		Currency:    entity.CurrencyRUB,
		Weight:      300,
		Properties:  `{"speeds":"9sp","material":"steel","links":114}`,
	},
	&entity.Ad{
		UUID:        uuid.FromStringOrNil("51dd8034-3618-4483-ab17-da72a6dca97a"),
		Name:        "Shimano CN-HG53 9-speed Chain",
		Description: "almost new!",
		UserUUID:    TestUser.UUID,
		CategoryID:  Categories.Chain,
		BrandID:     Brands.SRAM,
		Condition:   entity.ConditionUsedLikeNew,
		LocalityID:  Localities.SPB,
		Price:       55000,
		Currency:    entity.CurrencyRUB,
		Weight:      299,
		Properties:  `{"speeds":"9sp","material":"steel","links":114}`,
	},
}
