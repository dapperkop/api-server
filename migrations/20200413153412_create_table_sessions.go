package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200413153412, Down20200413153412)
}

// Up20200413153412 func ...
func Up20200413153412(tx *sql.Tx) error {
	var query string

	query += "CREATE TABLE `sessions` ("
	query += "`id` INT UNSIGNED NOT NULL AUTO_INCREMENT, "
	query += "`token` VARCHAR(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL, "
	query += "`user_id` INT UNSIGNED NOT NULL, "
	query += "PRIMARY KEY (`id`), UNIQUE (`token`)"
	query += ") ENGINE=InnoDB CHARSET=utf8 COLLATE utf8_general_ci;"

	if _, err := tx.Exec(query); err != nil {
		return err
	}

	return nil
}

// Down20200413153412 func ...
func Down20200413153412(tx *sql.Tx) error {
	var query = "DROP TABLE `sessions`;"

	if _, err := tx.Exec(query); err != nil {
		return err
	}

	return nil
}
