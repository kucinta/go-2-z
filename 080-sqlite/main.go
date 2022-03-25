package main

import (
    "database/sql"
    "fmt"
    "strconv"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
	myDb, _ := sql.Open("sqlite3", "./bogo.db")
	myQuery, _ := myDb.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
    myQuery.Exec()
	myQuery, _ = myDb.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
    myQuery.Exec("Rob", "Gronkowski")
	myQuery.Exec("KY", "Lee")
	myRecord, _ := myDb.Query("SELECT id, firstname, lastname FROM people")
    var id int
    var firstname string
    var lastname string
    for myRecord.Next() {
        myRecord.Scan(&id, &firstname, &lastname)
        fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
    }
}