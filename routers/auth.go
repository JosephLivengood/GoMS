package routers

import (
	"GoMS/handlers"
	"github.com/go-chi/chi/v5"
)

func AuthRouter(handlers handlers.HandlerMap) chi.Router {
	router := chi.NewRouter()

	router.Get("/me", handlers.Auth.Create)

	return router
}
