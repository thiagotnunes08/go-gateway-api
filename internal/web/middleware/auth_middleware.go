package middleware

import (
	"net/http"

	"github.com/thiagotnunes08/go-gateway-api/internal/domain"
	"github.com/thiagotnunes08/go-gateway-api/internal/service"
)

type AuthMiddleware struct {
	acccountService *service.AccountService
}

func NewAuthMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{acccountService: accountService}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apiKey := r.Header.Get("X-API-KEY")

		if apiKey == "" {
			http.Error(w, "X-API-KEY is required", http.StatusUnauthorized)
			return
		}

		_, err := m.acccountService.FindByAPIKey(apiKey)

		if err != nil {

			if err == domain.ErrAccountNotFound {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)

	})
}
