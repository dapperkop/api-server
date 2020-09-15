package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dapperkop/api-server/app/helpers"
	"github.com/dapperkop/api-server/app/model"
)

// Login func ...
func Login(w http.ResponseWriter, r *http.Request) {
	var (
		credentials model.Credentials
		err         error
	)

	err = json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		helpers.Response(
			errors.New("Bad Request"),
			http.StatusBadRequest,
			w,
			r,
		)

		return
	}

	var session model.Session

	session, err = helpers.AuthAttempt(credentials)

	if err != nil {
		helpers.Response(
			err,
			http.StatusUnauthorized,
			w,
			r,
		)

		return
	}

	helpers.Response(
		struct {
			Token string     `json:"token"`
			User  model.User `json:"user"`
		}{
			Token: session.Token,
			User:  session.User,
		},
		http.StatusOK,
		w,
		r,
	)
}

// Logout func ...
func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout")
}
