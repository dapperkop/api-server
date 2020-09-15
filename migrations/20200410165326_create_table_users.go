package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200410165326, Down20200410165326)
}

// Up20200410165326 func ...
func Up20200410165326(tx *sql.Tx) error {
	var query string

	query += "CREATE TABLE `users` ("
	query += "`id` INT UNSIGNED NOT NULL AUTO_INCREMENT, "
	query += "`username` VARCHAR(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL, "
	query += "`email` VARCHAR(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL, "
	query += "`password_hash` VARCHAR(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL, "
	query += "`is_active` BOOLEAN NOT NULL DEFAULT TRUE, "
	query += "`role` ENUM('admin','user') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'user', "
	query += "PRIMARY KEY (`id`), UNIQUE (`email`)"
	query += ") ENGINE=InnoDB CHARSET=utf8 COLLATE utf8_general_ci;"

	if _, err := tx.Exec(query); err != nil {
		return err
	}

	return nil
}

// Down20200410165326 func ...
func Down20200410165326(tx *sql.Tx) error {
	var query = "DROP TABLE `users`;"

	if _, err := tx.Exec(query); err != nil {
		return err
	}

	return nil
}
