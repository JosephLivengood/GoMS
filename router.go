package main

import (
	"time"

	"GoMS/handlers"
	localMiddleware "GoMS/middleware"
	"GoMS/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func initRouter(handlerMap handlers.HandlerMap) (router chi.Router) {
	router = chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/_health"))
	router.Use(middleware.Timeout(60 * time.Second))

	router.Use(localMiddleware.JsonResponse)

	router.Mount("/ping", routers.PingRouter(handlerMap))

	return
}
