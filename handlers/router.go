package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	router.Route("/api", func(api chi.Router) {
		api.Route("/users", func(u chi.Router) {
			u.Get("/", h.GetAllUsers)
			u.Post("/", h.CreateUser)
			u.Route("/{user}", func(r chi.Router) {
				r.Get("/", h.GetUser)
				r.Delete("/", h.DeleteUser)
				r.Patch("/", h.UpdateUser)
			})
		})

		api.Route("/transactions", func(t chi.Router) {
			t.Get("/", h.GetAllTransactions)
			t.Post("/deposit", h.CreateDeposit)
			t.Post("/withdraw", h.CreateWithdraw)
			t.Post("/transfer", h.CreateTransfer)
			t.Route("/{id}", func(r chi.Router) {
				r.Get("/", h.GetTransaction)
				r.Delete("/", h.DeleteTransaction)
				r.Patch("/", h.UpdateTransaction)
			})
		})
	})
	return router
}
