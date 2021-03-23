package main

import (
	"net/http"
)

//go:generate mockgen -destination=../mocks/mock_bolk.go -package=main github.com/mf-eqsdigital/go-study/main Bolk

type Bolk interface {
	GetCounter(w http.ResponseWriter, r *http.Request)
	UpdateCounter(w http.ResponseWriter, r *http.Request)
}

// func GetCounter(w http.ResponseWriter, r *http.Request) Counter {
// 	db, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	var counter Counter
// 	db.First(&counter)

// 	return counter
// 	//fmt.Fprintf(w, "Counter from DB")
// 	//json.NewEncoder(w).Encode(counter)
// }
