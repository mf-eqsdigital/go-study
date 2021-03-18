package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mf-eqsdigital/go-study/services"
)

type App struct {
	Router         *mux.Router
	DB             *gorm.DB
	CounterService *services.CounterService
}

func (a *App) Initialize(dbname string) {

	a.DB = DB_Connect(dbname)
	seed_DB(a.DB)

	a.Router = mux.NewRouter()
	a.handleRequests()

	defer a.DB.Close()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func seed_DB(db *gorm.DB) {
	if !db.HasTable(&Counter{}) {
		initialMigration(db)
	}
}

func (a *App) handleRequests() {
	a.Router.HandleFunc("/counter", a.CounterService.GetCounter).Methods("GET")
	a.Router.HandleFunc("/counter", a.CounterService.UpdateCounter).Methods("PUT")
}
