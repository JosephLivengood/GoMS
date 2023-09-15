package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"GoMS/models"
	"GoMS/stores"
)

func NewAuthHandler(storeMap stores.StoreMap) AuthHandler {
	return authHandler{
		storeMap.User,
	}
}

type authHandler struct {
	UserStore stores.UserStore
}

type AuthHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func (a authHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Email: "a",
		Password: "b",
		LastLogin: time.Now(),
		CreatedAt: time.Now(),
	}
	err := a.UserStore.Save(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, _ := json.Marshal(user)
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}
