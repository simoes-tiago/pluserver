package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"pluserver/service"
)

type handler struct {
	svc service.Service
}

func InitRouter(svc service.Service) *chi.Mux {
	h := handler{svc}

	router := chi.NewRouter()
	// Apply middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, hello!"))
	})

	router.Route("/users", func(u chi.Router) {
		u.Get("/", h.GetAllUsers)
		u.Post("/", h.CreateUser)
		u.Route("/{user}", func(r chi.Router) {
			r.Get("/", h.GetUser)
			r.Delete("/", h.DeleteUser)
			r.Patch("/", h.UpdateUser)
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
