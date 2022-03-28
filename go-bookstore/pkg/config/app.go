package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//d, err := gorm.Open("mysql", "root:@/go-bookstore?charset=utf8&parseTime=True&loc=Local")
	dsn := "root:@tcp(127.0.0.1:4000)/go_bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
