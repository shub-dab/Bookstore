package config

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB


var cfg = mysql.Config{
	User:   "root",
	Passwd: "abcd",
	Net:    "tcp",
	Addr:   "127.0.0.1:3306",
	DBName: "bookstore",
	ParseTime: true,
}

func Connect() {
	d, err := gorm.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}