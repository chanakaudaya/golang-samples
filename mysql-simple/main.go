package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID int `json:"id"`
	Name string `json:"name"`

}

func main() {
	fmt.Println("Go MySQL Example")

	// Connect to the database
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Root@1985"
	dbName := "opd_data"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// If there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// Defer the close till after the main function is finished
	defer db.Close()

	// Perform an insert operation
	insert, err := db.Query("INSERT INTO test VALUES (4, 'TEST44')")

	// If there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	// Read from the database
	results, err := db.Query("SELECT id, name FROM test")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(tag.Name)
	}


}