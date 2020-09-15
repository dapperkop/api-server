package helpers

import (
	"errors"
	"net/http"

	"github.com/dapperkop/api-server/app/model"
)

// Auth func ...
func Auth(r *http.Request) (model.User, bool) {
	var (
		auth  bool
		token string
		user  model.User
	)

	token = r.Header.Get("X-Session-Token")

	if token == "" {
		return user, auth
	}

	user, auth = model.GetUserByToken(token)

	return user, auth
}

// AuthAttempt func ...
func AuthAttempt(credentials model.Credentials) (model.Session, error) {
	var (
		found bool
		user  model.User
	)

	user, found = model.GetUserByCredentials(credentials)

	if !found {
		user, found = model.LDAPUserByCredentials(credentials)
	}

	var session = model.NewSession(user)

	if !found {
		return *session, errors.New("Email or password incorrect")
	}

	session.Save()

	return *session, nil
}
