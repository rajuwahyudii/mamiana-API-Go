package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mamiana")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	insert, err := db.Query("INSERT INTO user VALUES('raju wahyudi pratama','23','laki','90')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Hello world")
}
