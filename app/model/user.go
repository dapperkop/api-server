package model

import (
	"database/sql"
	"fmt"

	"github.com/dapperkop/blank/database"
	"github.com/dapperkop/blank/logger"
	"github.com/go-ldap/ldap/v3"
)

// Credentials type ...
type Credentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User type ...
type User struct {
	ID           int    `db:"id" json:"id"`
	Username     string `db:"username" json:"username"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"-"`
	IsActive     bool   `db:"is_active" json:"is_active"`
	Role         string `db:"role" json:"role"`
}

// GetUserByCredentials func ...
func GetUserByCredentials(credentials Credentials) (User, bool) {
	var (
		err   error
		found bool
		query string
		row   *sql.Row
		user  User
	)

	query += "SELECT `id`, `username`, `email`, `password_hash`, `is_active`, `role` FROM `users` "
	query += "WHERE (`username` = ? OR `email` = ?) AND `password_hash` = MD5(?);"

	row = database.DB.QueryRow(query, credentials.Username, credentials.Email, credentials.Password)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.IsActive, &user.Role)

	switch err {
	case sql.ErrNoRows:
		found = false
	case nil:
		found = true
	default:
		logger.Logger.Fatalln(err)
	}

	return user, found
}

// LDAPUserByCredentials func ...
func LDAPUserByCredentials(credentials Credentials) (User, bool) {
	var (
		conn *ldap.Conn
		err  error
	)

	conn, err = ldap.Dial("tcp", "domain.com:389")

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	defer conn.Close()

	// err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})

	// if err != nil {
	// 	logger.Logger.Fatalln(err)
	// }

	var (
		user     = new(User)
		username = "uid=" + credentials.Username + ",dc=domain,dc=dom"
		password = credentials.Password
	)

	err = conn.Bind(username, password)

	if err != nil {
		return *user, false
	}

	user.Username = credentials.Username
	user.Email = fmt.Sprintf("%s@domain.com", credentials.Username)
	user.PasswordHash = credentials.Password
	user.IsActive = true
	user.Role = "user"

	user.Save()

	return *user, true
}

// Save func ...
func (user *User) Save() {
	var (
		err          error
		lastInsertID int64
		query        string
		result       sql.Result
	)

	if user.ID == 0 {
		query += "INSERT INTO `users` (`id`, `username`, `email`, `password_hash`, `is_active`, `role`) "
		query += "VALUES (NULL, ?, ?, MD5(?), ?, ?);"

		result, err = database.DB.Exec(
			query,
			user.Username,
			user.Email,
			user.PasswordHash,
			user.IsActive,
			user.Role,
		)
	} else {
		query += "UPDATE `users` SET "
		query += "`username` = ?, "
		query += "`email` = ?, "
		query += "`password_hash` = MD5(?), "
		query += "`is_active` = ?, "
		query += "`role` = ? "
		query += "WHERE `users`.`id` = ?;"

		result, err = database.DB.Exec(
			query,
			user.Username,
			user.Email,
			user.PasswordHash,
			user.IsActive,
			user.Role,
			user.ID,
		)
	}

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	if user.ID == 0 {
		lastInsertID, err = result.LastInsertId()

		if err != nil {
			logger.Logger.Fatalln(err)
		}

		user.ID = int(lastInsertID)
	}
}
