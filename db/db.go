package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rubenv/sql-migrate"
)

var db *sql.DB

func init() {
	d, err := sql.Open("sqlite3", "./vcpod.db")
	if err != nil {
		panic(err)
	}
	migrations := &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "db/migrations",
	}
	if _, err := migrate.Exec(d, "sqlite3", migrations, migrate.Up); err != nil {
		panic(err)
	}
	db = d
}

func DB() *sql.DB {
	return db
}
