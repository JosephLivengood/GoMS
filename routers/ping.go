package routers

import (
	"GoMS/handlers"
	"github.com/go-chi/chi/v5"
)

func PingRouter(handlers handlers.HandlerMap) (router chi.Router) {
	router = chi.NewRouter()

	router.Get("/", handlers.Ping.Create)
	router.Get("/all", handlers.Ping.GetAll)
	router.Get("/{pingID}", handlers.Ping.GetByID)

	return
}
