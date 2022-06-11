package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseConnections struct {
	Primary *gorm.DB
}

func populateDBConnections(config *Config) {
	config.DB = &DatabaseConnections{
		Primary: initPrimary(config),
	}
}

func initPrimary(config *Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.Fields.Database.Primary.ConnectionPath), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
