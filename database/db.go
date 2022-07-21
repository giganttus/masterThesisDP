package database

import (
	"github.com/go-pg/pg"
	"os"
)

func Connect() *pg.DB {

	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
		panic("PostgreSQL is down")
	}
	return db
}
