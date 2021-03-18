package main

import (
	// "encoding/json"
	// "fmt"
	// "net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Counter struct {
	gorm.Model
	Number int
}

func initialMigration(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&Counter{
		Model:  gorm.Model{},
		Number: 0,
	})

	// Create reccord
	db.Create(&Counter{Number: 0})
}

// func (a *App)updateCounter(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	var counter Counter
// 	db.First(&counter)

// 	num := counter.Number
// 	counter.Number = num + 1

// 	db.Save(&counter)
// 	fmt.Fprintf(w, "Successfully Updated Counter from %d to %d", num, counter.Number)
// 	json.NewEncoder(w).Encode(counter)
// }

// func (a *App) getCounter(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	var counter Counter
// 	db.First(&counter)

// 	fmt.Fprintf(w, "Counter from DB")
// 	json.NewEncoder(w).Encode(counter)
// }
