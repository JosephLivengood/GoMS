package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"GoMS/stores"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func NewPingHandler(storeMap stores.StoreMap) PingHandler {
	return pingHandler{
		storeMap.Ping,
	}
}

type pingHandler struct {
	pingStore stores.PingStore
}

type PingHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

func (p pingHandler) Create(w http.ResponseWriter, r *http.Request) {
	ping, err := p.pingStore.Create(r.RemoteAddr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a, _ := json.Marshal(ping)
	w.WriteHeader(http.StatusCreated)
	w.Write(a)
}

func (p pingHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pings, err := p.pingStore.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a, _ := json.Marshal(pings)
	w.Write(a)
}

func (p pingHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	inputPingID := chi.URLParam(r, "pingID")

	pingID, err := strconv.Atoi(inputPingID)
	if err != nil {
		// inputPingID was not an int, won't be found
		http.Error(w, "Ping not found", http.StatusBadRequest)
		return
	}

	ping, err := p.pingStore.GetById(pingID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Ping not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	a, _ := json.Marshal(ping)
	w.Write(a)
}
