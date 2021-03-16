package main

import (
	"github.com/jinzhu/gorm"
)

func DB_Connect(dbname string) *gorm.DB {
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	DB, err := gorm.Open("sqlite3", dbname)
	if err != nil {
		panic(err.Error())
	}
	return DB
}
