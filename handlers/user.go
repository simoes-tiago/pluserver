package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"pluserver/domain"
)

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.svc.GetAllUser()
	b, err := json.Marshal(users)
	if err != nil {
		log.Println("error getting users:", err)
	}
	w.Write(b)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	user := h.svc.GetUser(chi.URLParam(r, "user"))

	if user.Username == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		log.Println("error getting user:", err)
	}
	w.Write(b)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "bad input", http.StatusUnprocessableEntity)
		return
	}
	err := h.svc.CreateUser(user)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("created %v", user)))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := h.svc.DeleteUser(chi.URLParam(r, "user"))
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("deleted"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "bad input", http.StatusUnprocessableEntity)
		return
	}
	err := h.svc.UpdateUser(chi.URLParam(r, "user"), user)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("updated"))
}