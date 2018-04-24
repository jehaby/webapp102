package service

import (
	"sync"
	"time"

	"github.com/go-pg/pg"
	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/log"
)

var categoryService *CategoryService

type CategoryService struct {
	db  *pg.DB
	log *log.Logger

	lock  sync.RWMutex
	cache map[int64]*entity.Category
}

func NewCategoryService(pgdb *pg.DB, log *log.Logger) *CategoryService {
	if categoryService != nil {
		return categoryService
	}

	categoryService = &CategoryService{
		db:  pgdb,
		log: log,
	}

	go func() {
		t := time.NewTicker(time.Minute)
		for {
			categoryService.updateCache()
			<-t.C
		}
	}()

	return categoryService
}

func (cs *CategoryService) updateCache() {
	categories := []*entity.Category{}
	err := cs.db.Model(&categories).Select()
	if err != nil {
		cs.log.WithError(err).Error("CategoryService.updateCache: error from db")
	}

	cache := make(map[int64]*entity.Category)
	for _, cat := range categories {
		cache[cat.ID] = cat
	}
	cs.lock.Lock()
	cs.cache = cache
	cs.lock.Unlock()
}

func (cs *CategoryService) GetByID(id int64) *entity.Category {
	res, _ := cs.cache[id]
	return res
}

func (cs *CategoryService) CategoryExists(id int64) bool {
	cs.lock.RLock()
	_, ok := cs.cache[id]
	cs.lock.RUnlock()
	return ok
}
