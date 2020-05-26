package model

import (
	"fmt"
	"log"
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
var testDB *gorm.DB

// Model is
type Model struct {
	ID        int64     `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Db is
func Db() (*gorm.DB, error) {
	test := os.Getenv("TRANSACTIONAL_TEST")
	if test == "true" {
		return TestDB(), nil
	}
	return gorm.Open("mysql", mysqlURL)
}

// TestDB is
func TestDB() *gorm.DB {
	if testDB == nil {
		SetTestDB()
	}
	return testDB
}

// SetTestDB is
func SetTestDB() {
	mysqlDatabase = "coffeehub_test"
	mysqlURL = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase)
	db, err := gorm.Open("mysql", mysqlURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set testDB")
	testDB = db
}

// StartTx is
func StartTx(db *gorm.DB) *gorm.DB {
	tx := db.Begin()
	testDB = tx

	return tx
}

func init() {
	if mysqlHost == "" {
		mysqlHost = "127.0.0.1"
	}

	mysqlURL = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase)
}
