package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Db() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "")
	return db, err
}
