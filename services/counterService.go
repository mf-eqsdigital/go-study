package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"bolk"
	"model"
)

type CounterService struct {
	b bolk.Bolk
}

func (u *CounterService) UpdateCounter(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var counter model.Counter
	db.First(&counter)

	num := counter.Number
	counter.Number = num + 1

	db.Save(&counter)
	fmt.Fprintf(w, "Successfully Updated Counter from %d to %d", num, counter.Number)
	json.NewEncoder(w).Encode(counter)
}

func (u *CounterService) GetCounter(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var counter model.Counter
	db.First(&counter)

	fmt.Fprintf(w, "Counter from DB")
	json.NewEncoder(w).Encode(counter)
}
