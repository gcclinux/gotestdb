package proc

import (
	"testdb/conn"

	_ "github.com/lib/pq"
)

func Insert(name, surename, age string) {

	db := conn.ConnectDB
	err := db().Ping()
	CheckError(err)

	stmt := `INSERT INTO public.testtable(firstname, surename, age) VALUES ($1, $2, $3)`
	_, err = db().Exec(stmt, name, surename, age)
	if err != nil {
		panic(err)
	}
}
