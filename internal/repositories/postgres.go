package repositories

import (
	"gorm.io/gorm"
)

type Postgres struct {
	dbGorm *gorm.DB
}

func NewPostgres(dbGorm *gorm.DB) *Postgres {
	return &Postgres{
		dbGorm: dbGorm,
	}
}
