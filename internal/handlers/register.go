package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Controller interface {
	RegisterHandlers(router chi.Router)
}

func Register(router chi.Router, controllers ...Controller) {
	router.Use(Logger)
	router.Use(Recovery)
	router.Mount("/debug", middleware.Profiler())

	for _, controller := range controllers {
		controller.RegisterHandlers(router)
	}
}
