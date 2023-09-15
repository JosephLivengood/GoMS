package stores

import (
	"GoMS/models"
	"gorm.io/gorm"
)

//go:generate mockgen -source=user.go -destination=mocks/user.go -package=mocks

func NewUserStore(db *gorm.DB) UserStore {
	_ = db.AutoMigrate(&models.User{})
	return userStore{db}
}

type userStore struct {
	db *gorm.DB
}

type UserStore interface {
	Save(user *models.User) error
}

func (u userStore) Save(user *models.User) error {
	if user == nil {
		return nil
	}
	// Better error code handling to wrapped error
	return u.db.Create(user).Error
}
