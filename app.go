package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(dbname string) {

	a.DB = DB_Connect(dbname)
	seed_DB(a.DB)

	a.Router = mux.NewRouter()
	defer a.DB.Close()
}

func (a *App) Run(addr string) {}

func seed_DB(db *gorm.DB) {
	if !db.HasTable(&Counter{}) {
		initialMigration(db)
	}
}

func (a *App) handleRequests() {
	a.Router.HandleFunc("/counter", getCounter).Methods("GET")
	a.Router.HandleFunc("/counter", updateCounter).Methods("PUT")
}
