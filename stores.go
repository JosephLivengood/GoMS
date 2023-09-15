package main

import (
	"GoMS/config"
	"GoMS/stores"
)

func initStores(config *config.Config) (storeMap stores.StoreMap) {

	storeMap.Ping = stores.NewPingStore(config.DB.Primary)
	storeMap.User = stores.NewUserStore(config.DB.Primary)

	return
}
