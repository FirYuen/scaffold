package dao

import (
	"context"
	"gorm.io/gorm"
)

var (
	ErrNotFound = gorm.ErrRecordNotFound
)

// Dao mysql struct
type Dao struct {
	db        *gorm.DB
	//userCache *cache.Cache
}

// New new a Dao and return
func New(db *gorm.DB) *Dao {
	return &Dao{
		db:        db,
		//userCache: cache.NewUserCache(),
	}
}

// Ping ping mysql
func (d *Dao) Ping(c context.Context) error {
	return nil
}

// Close release mysql connection
func (d *Dao) Close() {

}
