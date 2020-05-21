package model

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // blank import
)

var mysqlHost = os.Getenv("MYSQL_HOST")
var mysqlDatabase = os.Getenv("MYSQL_DATABASE")
var mysqlUser = os.Getenv("MYSQL_USER")
var mysqlPassword = os.Getenv("MYSQL_PASSWORD")
var mysqlURL string

// Model is
type Model struct {
	ID        int64     `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Db is
func Db() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", mysqlURL)
	return db, err
}

func init() {
	if mysqlHost == "" {
		mysqlHost = "127.0.0.1"
	}

	mysqlURL = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase)
}
