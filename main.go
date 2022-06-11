package main

import (
	"fmt"
	"net/http"

	. "GoMS/config"
)

func main() {
	config := GetConfig()

	stores := initStores(config)
	handlers := initHandlers(stores)
	router := initRouter(handlers)

	err := http.ListenAndServe(fmt.Sprint(":", config.Fields.Server.Port), router)
	if err != nil {
		panic(err.Error())
	}
}
