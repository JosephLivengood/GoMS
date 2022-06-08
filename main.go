package main

import (
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	router := initRouter(initHandlers(initStores(db)))

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err.Error())
	}
}
