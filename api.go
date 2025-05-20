package handlers

import (
	"github.com/AbdulQuayyum/go-test/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/user", func(router chi.Router) {

		router.Use(middleware.Authorization)

		router.Get("/shawties", GetShawtiesAmount)
	})
}
