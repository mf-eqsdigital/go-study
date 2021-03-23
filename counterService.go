package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	//"interfaces/bolk"
	// "model"
)

type CounterService struct {
	B Bolk
}

type Sargenor struct {
}

func (s *CounterService) GetCounter(w http.ResponseWriter, r *http.Request) {
	res := Get(w, r)

	fmt.Fprintf(w, "Counter from DB")
	json.NewEncoder(w).Encode(res)
}

func Get(w http.ResponseWriter, r *http.Request) Counter {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var counter Counter
	db.First(&counter)

	return counter
}

func (s *CounterService) UpdateCounter(w http.ResponseWriter, r *http.Request) {
	res := Update(w, r)
	fmt.Fprintf(w, "Successfully Updated Counter from %d to %d", res.Number-1, res.Number)
	json.NewEncoder(w).Encode(res)
}

func Update(w http.ResponseWriter, r *http.Request) Counter {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var counter Counter
	db.First(&counter)

	num := counter.Number
	counter.Number = num + 1

	db.Save(&counter)

	return counter
}
