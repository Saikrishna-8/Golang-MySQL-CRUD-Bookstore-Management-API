package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Fassak@08@tcp(localhost:3306)/mydb?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("mysql", "user:password@tcp(localhost:3306)/database")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
