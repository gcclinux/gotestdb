package main

import (
	"fmt"
	"testdb/conn"
)

func main() {

	db := conn.ConnectDB
	err := db().Ping()
	CheckError(err)

	rows, err := db().Query(`SELECT "user_id", "firstname", "surename", "age" FROM "testtable" `)
	CheckError(err)
	defer rows.Close()

	fmt.Println("-------------------------")
	fmt.Println("id| age| name surname")
	fmt.Println("-------------------------")
	for rows.Next() {
		var user_id string
		var firstname string
		var surename string
		var age string
		err = rows.Scan(&user_id, &firstname, &surename, &age)
		CheckError(err)
		fmt.Println(user_id, "|", age, "|", firstname, surename)
	}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
