package router

import (
	"net/http"

	"github.com/dapperkop/blank/apiserver/httpserver"
	"github.com/dapperkop/api-server/app/handler"
	"github.com/dapperkop/api-server/app/middleware"
)

// Setup func ...
func Setup() {
	var (
		api    = httpserver.Router.PathPrefix("/api").Subrouter()
		login  = api.PathPrefix("/login").Subrouter()
		logout = api.PathPrefix("/logout").Subrouter()
	)

	login.Use(middleware.IsGuest)
	login.HandleFunc("/", handler.Login).Name("login").Methods(http.MethodPost)

	logout.Use(middleware.IsAuth)
	logout.HandleFunc("/", handler.Logout).Name("logout").Methods(http.MethodGet)
}
