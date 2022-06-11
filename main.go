package main

import (
	"net/http"

	. "GoMS/config"
)

func main() {
	config := GetConfig()

	router := initRouter(initHandlers(initStores(config)))

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err.Error())
	}
}
