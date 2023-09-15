package main

import (
	"fmt"
	"net/http"

	conf "GoMS/config"
)

func main() {
	config := conf.GetConfig()

	stores := initStores(config)
	handlers := initHandlers(stores)
	router := initRouter(handlers)

	err := http.ListenAndServe(fmt.Sprint(":", config.Fields.Server.Port), router)
	if err != nil {
		panic(err.Error())
	}
}
