package main

import (
	"GoMS/handlers"
	"GoMS/stores"
)

func initHandlers(storeMap stores.StoreMap) (handlerMap handlers.HandlerMap) {

	handlerMap.Ping = handlers.NewPingHandler(storeMap)

	return
}
