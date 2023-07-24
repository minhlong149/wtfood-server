package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	dbUri := os.Getenv("DATABASE_URI")

	var err error
	Db, err = sql.Open("postgres", dbUri)

	if err != nil {
		log.Fatal(err)
	}

	if err = Db.Ping(); err != nil {
		log.Fatal(err)
	}
}
