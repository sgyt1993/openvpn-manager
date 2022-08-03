package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dbPath = kingpin.Flag("db.path", "path do openvpn-user db").Default("./openvpn.db").Envar("DB_PATH").String()
)

func InitDb() {
	// boolean fields are integer because of sqlite does not support boolean: 1 = true, 0 = false
	_, err := GetDb().Exec(
		"CREATE TABLE IF NOT EXISTS role(" +
			"id integer not null primary key autoincrement, " +
			"role_name string UNIQUE)")

	CheckErr(err)
	fmt.Printf("Database initialized at %s\n", *dbPath)
}

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", *dbPath)
	CheckErr(err)
	if db == nil {
		panic("db is nil")
	}
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
