package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thiagotnunes08/go-gateway-api/internal/domain"
	"github.com/thiagotnunes08/go-gateway-api/internal/dto"
	"github.com/thiagotnunes08/go-gateway-api/internal/service"
)

type InvoicetHandler struct {
	invoiceService *service.InvoiceService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService) *InvoicetHandler {
	return &InvoicetHandler{invoiceService: invoiceService}
}

func (h *InvoicetHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateInvoiceInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.APIKey = r.Header.Get("X-API-KEY")

	output, err := h.invoiceService.Create(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)

}

func (h *InvoicetHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-KEY")

	if apiKey == "" {
		http.Error(w, "X-API-KEY is required", http.StatusBadRequest)
		return
	}

	output, err := h.invoiceService.GetByID(id, apiKey)

	if err != nil {
		switch err {

		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return

		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return

		case domain.ErrUnauthorizedAcces:
			http.Error(w, err.Error(), http.StatusForbidden)
			return

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)

}

func (h *InvoicetHandler) ListByAccount(w http.ResponseWriter, r *http.Request) {

	apiKey := r.Header.Get("X-API-KEY")

	if apiKey == "" {
		http.Error(w, "X-API-KEY is required", http.StatusUnauthorized)
		return
	}

	output, err := h.invoiceService.ListByAccountAPIKey(apiKey)

	if err != nil {
		switch err {

		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
