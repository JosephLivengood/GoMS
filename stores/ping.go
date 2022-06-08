package stores

import (
	"GoMS/models"
	"gorm.io/gorm"
)

//go:generate mockgen -source=ping.go -destination=mocks/ping.go -package=mocks

func NewPingStore(db *gorm.DB) PingStore {
	_ = db.AutoMigrate(&models.Ping{})
	return pingStore{db}
}

type pingStore struct {
	db *gorm.DB
}

type PingStore interface {
	Create(createdBy string) (*models.Ping, error)
	GetAll() ([]models.Ping, error)
	GetById(id int) (*models.Ping, error)
}

func (p pingStore) Create(createdBy string) (*models.Ping, error) {
	ping := &models.Ping{CreatedBy: createdBy}
	err := p.db.Create(ping).Error

	return ping, err
}

func (p pingStore) GetAll() ([]models.Ping, error) {
	var pings []models.Ping
	err := p.db.Find(&pings).Error

	return pings, err
}

func (p pingStore) GetById(id int) (*models.Ping, error) {
	ping := &models.Ping{ID: id}
	err := p.db.First(ping).Error

	return ping, err
}
