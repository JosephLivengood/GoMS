package main

import (
	"GoMS/handlers"
	"GoMS/stores"
)

func initHandlers(storeMap stores.StoreMap) (handlerMap handlers.HandlerMap) {

	handlerMap.Ping = handlers.NewPingHandler(storeMap)
	handlerMap.Auth = handlers.NewAuthHandler(storeMap)

	return
}
