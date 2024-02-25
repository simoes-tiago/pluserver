package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"pluserver/domain"
	"strconv"
)

func (h *handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := h.svc.GetAllTransactions()
	b, err := json.Marshal(transactions)
	if err != nil {
		log.Println("error getting transactions:", err)
	}
	w.Write(b)
}

func (h *handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	transactionId, err := strconv.Atoi(chi.URLParam(r, "id"))
	transaction := h.svc.GetTransaction(uint(transactionId))

	if transaction.ID == 0 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	b, err := json.Marshal(transaction)
	if err != nil {
		log.Println("error getting transaction:", err)
	}
	w.Write(b)
}

func (h *handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "not found", http.StatusUnprocessableEntity)
		return
	}
	err = h.svc.DeleteTransaction(uint(transactionId))
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("deleted"))
}

func (h *handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction domain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "bad input", http.StatusUnprocessableEntity)
		return
	}
	transactionId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "not found", http.StatusUnprocessableEntity)
		return
	}
	err = h.svc.UpdateTransaction(uint(transactionId), transaction)
	if err != nil {
		http.Error(w, "not found", http.StatusBadRequest)
		return
	}

	w.Write([]byte("updated"))
}

func (h *handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var transaction domain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "bad input", http.StatusUnprocessableEntity)
		return
	}
	transaction.Type = domain.Transfer
	err := h.svc.CreateTransaction(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("created %v", r.Body)))
}

func (h *handler) CreateDeposit(w http.ResponseWriter, r *http.Request) {
	var transaction domain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "bad input", http.StatusUnprocessableEntity)
		return
	}
	transaction.Type = domain.Deposit
	err := h.svc.CreateTransaction(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("created %v", r.Body)))
}

func (h *handler) CreateWithdraw(w http.ResponseWriter, r *http.Request) {
	var transaction domain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "bad input", http.StatusUnprocessableEntity)
		return
	}
	transaction.Type = domain.Withdraw
	err := h.svc.CreateTransaction(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("created %v", r.Body)))
}
