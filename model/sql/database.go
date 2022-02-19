package sql

import (
	"database/sql"
	"log"
	"os"
	"qltTestApi/utils"
)

func InitDb() *sql.DB {
	db, err := utils.ConnectDb("sqlite3", "file:database.db?sqlite_foreign_keys=on")
	if err != nil {
		panic(err)
	}
	db.SQLDB.Ping()

	userQuery, err := os.ReadFile("model/sql/payments.sql")
	if err != nil {
		log.Println(err)
	}
	if _, err := db.SQLDB.Exec(string(userQuery)); err != nil {
		log.Println(err)
	}

	return db.SQLDB
}
