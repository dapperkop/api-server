package middleware

import (
	"net/http"

	"github.com/dapperkop/api-server/app/helpers"
)

// IsAdmin func ...
func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, auth := helpers.Auth(r); auth && user.IsActive && user.Role == "admin" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	})
}
