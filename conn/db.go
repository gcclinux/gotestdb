package conn

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	Host, Port, User, Pass, DB := Getconf()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Pass, DB)
	db, _ := sql.Open("postgres", psqlInfo)

	return db
}

func CloseDB() {
	db := ConnectDB
	defer db().Close()
}
