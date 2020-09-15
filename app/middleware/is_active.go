package middleware

import (
	"net/http"

	"github.com/dapperkop/api-server/app/helpers"
)

// IsActive func ...
func IsActive(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, auth := helpers.Auth(r); auth && user.IsActive {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	})
}
