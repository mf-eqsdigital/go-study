package main

import (
	"github.com/jinzhu/gorm"
)

func DB_Connect() *gorm.DB {
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err.Error())
	}
	return db
}
