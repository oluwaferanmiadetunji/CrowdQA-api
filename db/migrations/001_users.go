package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func Init() {
	goose.AddMigration(Up001, Down001)
}

func Up001(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE users SET username='admin' WHERE username='root';")
	if err != nil {
		return err
	}
	return nil
}

func Down001(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE users SET username='root' WHERE username='admin';")
	if err != nil {
		return err
	}
	return nil
}