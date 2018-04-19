package service

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/go-pg/pg"
	"github.com/jmoiron/sqlx"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/storage"
)

type AdService struct {
	Repo adServiceRepos
	db   *pg.DB
	val  *validator.Validate
}

type adServiceRepos struct {
	Ad       *storage.AdRepository
	Category *storage.CategoryRepository
}

func newAdService(db *sqlx.DB, pgdb *pg.DB, val *validator.Validate) *AdService {
	return &AdService{
		Repo: adServiceRepos{
			Ad:       storage.NewAdRepository(db, pgdb),
			Category: storage.NewCategoryRepository(db),
		},
		db:  pgdb,
		val: val,
	}
}

type AdCreateArgs struct {
	Name        string `validate:"required,min=2"`
	Description string `validate:"required,min=5"`
	UserUUID    string `validate:"required"` // TODO: uuid validattion?
	ComponentID string `validate:"required,numeric,min=1"`
}

func (as *AdService) Create(args AdCreateArgs) (*entity.Ad, error) {
	if err := as.val.Struct(args); err != nil {
		return nil, err
	}

	componentID, _ := strconv.ParseInt(args.ComponentID, 10, 64)
	userUUID, err := uuid.FromString(args.UserUUID)
	if err != nil {
		return nil, errors.Wrapf(err, "AdService.Create")
	}

	comp := &entity.Component{ID: componentID}
	err = as.db.Model(comp).Column("category_id").WherePK().First()
	if err != nil {
		return nil, errors.Wrapf(err, "TODO")
	}

	ad := &entity.Ad{
		UUID:        uuid.NewV4(),
		CreatedAt:   time.Now(),
		Name:        args.Name,
		Description: args.Description,
		CategoryID:  comp.CategoryID,
		ComponentID: componentID,
		UserUUID:    userUUID,
	}

	_, err = as.db.Model(ad).Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "TODO")
	}

	return ad, nil
}
