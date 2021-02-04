package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {
	fmt.Println("Go ORM Tutorial")

	db := DB_Connect()
	seed_DB(db)

	handleRequests()
	defer db.Close()
}

func seed_DB(db *gorm.DB) {
	if !db.HasTable(&Counter{}) {
		initialMigration(db)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/counter", getCounter).Methods("GET")
	myRouter.HandleFunc("/counter", updateCounter).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
