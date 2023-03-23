package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:surabaya220699@tcp(localhost:3306)/beauty_store")
	if err != nil {
		return nil, err
	}
	return db, nil
}
