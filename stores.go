package main

import (
	"GoMS/stores"
	"gorm.io/gorm"
)

func initStores(db *gorm.DB) (storeMap stores.StoreMap) {

	storeMap.Ping = stores.NewPingStore(db)

	return
}
