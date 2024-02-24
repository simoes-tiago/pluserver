package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()
	// Apply middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, hello!"))
	})

	router.Route("/users", func(u chi.Router) {
		u.Get("/", func(writer http.ResponseWriter, request *http.Request) {

		})
		u.Post("/", func(writer http.ResponseWriter, request *http.Request) {

		})
		u.Delete("/", func(writer http.ResponseWriter, request *http.Request) {

		})
		u.Route("/{id}", func(r chi.Router) {
			r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("Hello, hello!"))
			})
		})
	})

	router.Route("/transactions", func(t chi.Router) {
		t.Get("/", func(writer http.ResponseWriter, request *http.Request) {

		})
		t.Post("/", func(writer http.ResponseWriter, request *http.Request) {

		})
		t.Delete("/", func(writer http.ResponseWriter, request *http.Request) {

		})
		t.Route("/{id}", func(r chi.Router) {

		})
	})

	return router
}
