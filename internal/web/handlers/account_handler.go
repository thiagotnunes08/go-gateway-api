package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thiagotnunes08/go-gateway-api/internal/dto"
	"github.com/thiagotnunes08/go-gateway-api/internal/service"
)

type AccountHandler struct {
	acccountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{acccountService: accountService}
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateAccountInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.acccountService.CreateAccount(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)

}

func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-KEY")

	if apiKey == "" {
		http.Error(w, "API Key is required", http.StatusUnauthorized)
		return
	}

	output, err := h.acccountService.FindByAPIKey(apiKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
