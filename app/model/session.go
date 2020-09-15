package model

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"

	"github.com/dapperkop/blank/database"
	"github.com/dapperkop/blank/logger"
)

// Session type ...
type Session struct {
	ID    int
	Token string
	User  User
}

var sessions = loadSessions()

func generateToken() string {
	rand.Seed(time.Now().UnixNano())

	var (
		b     strings.Builder
		chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		err   error
		n     = len(chars)
	)

	for i := 0; i < 256; i++ {
		_, err = b.WriteRune(chars[rand.Intn(n)])

		if err != nil {
			logger.Logger.Fatalln(err)
		}
	}

	return b.String()
}

// GetUserByToken func ...
func GetUserByToken(token string) (User, bool) {
	var (
		user  User
		found bool
	)

	user, found = sessions[token]

	return user, found
}

func loadSessions() map[string]User {
	var (
		err      error
		query    string
		rows     *sql.Rows
		sessions = make(map[string]User)
	)

	query += "SELECT `sessions`.`id`, `sessions`.`token`, "
	query += "`users`.`id`, `users`.`username`, `users`.`email`, "
	query += "`users`.`password_hash`, `users`.`is_active`, `users`.`role` "
	query += "FROM `sessions` INNER JOIN `users` ON `sessions`.`user_id` = `users`.`id`;"

	rows, err = database.DB.Query(query)

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			logger.Logger.Fatalln(err)
		}
	}()

	for rows.Next() {
		var session Session

		err = rows.Scan(
			&session.ID,
			&session.Token,
			&session.User.ID,
			&session.User.Username,
			&session.User.Email,
			&session.User.PasswordHash,
			&session.User.IsActive,
			&session.User.Role,
		)

		if err != nil {
			logger.Logger.Fatalln(err)
		}

		sessions[session.Token] = session.User
	}

	return sessions
}

// NewSession func ...
func NewSession(user User) *Session {
	var session = new(Session)

	session.Token = generateToken()
	session.User = user

	return session
}

// Save func ...
func (session *Session) Save() {
	var (
		err          error
		lastInsertID int64
		query        string
		result       sql.Result
	)

	if session.ID == 0 {
		query += "INSERT INTO `sessions` (`id`, `token`, `user_id`) "
		query += "VALUES (NULL, ?, ?);"

		result, err = database.DB.Exec(
			query,
			session.Token,
			session.User.ID,
		)
	} else {
		query += "UPDATE `sessions` SET "
		query += "`token` = ?, "
		query += "`user_id` = ? "
		query += "WHERE `sessions`.`id` = ?;"

		result, err = database.DB.Exec(
			query,
			session.Token,
			session.User.ID,
			session.ID,
		)
	}

	if err != nil {
		logger.Logger.Fatalln(err)
	}

	if session.ID == 0 {
		lastInsertID, err = result.LastInsertId()

		if err != nil {
			logger.Logger.Fatalln(err)
		}

		session.ID = int(lastInsertID)
	}

	sessions[session.Token] = session.User
}
