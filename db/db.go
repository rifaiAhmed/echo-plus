package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {

	dsn := "root:@tcp(127.0.0.1:3306)/kredit_plus?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	if err != nil {
		panic("dsn invalid")
	}

}

func CreateCon() *gorm.DB {
	return db
}
