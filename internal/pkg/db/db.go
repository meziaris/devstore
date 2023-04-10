package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(DBDriver string, DBConnection string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(DBDriver, DBConnection)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Print("database connection establish")
	return db, nil
}
