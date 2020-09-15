package middleware

import (
	"net/http"

	"github.com/dapperkop/api-server/app/helpers"
)

// IsAuth func ...
func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, auth := helpers.Auth(r); auth {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	})
}
