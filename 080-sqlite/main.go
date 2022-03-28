package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	myDb, _ := sql.Open("sqlite3", "./bogo.db")
	tableStructure := `
    CREATE TABLE IF NOT EXISTS article (
        id INTEGER PRIMARY KEY,
        title TEXT,
        content TEXT,
        date_created TEXT DEFAULT CURRENT_TIMESTAMP
    )
    `
	myQuery, _ := myDb.Prepare(tableStructure)
	myQuery.Exec()
	tableStructure = `
    CREATE TABLE IF NOT EXISTS comment (
        id INTEGER PRIMARY KEY,
		id_article INTEGER NOT NULL,
        remark TEXT,
        date_created TEXT DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(id_article) REFERENCES article(id)
		ON UPDATE CASCADE ON DELETE CASCADE
    )
    `
	myQuery, _ = myDb.Prepare(tableStructure)
	myQuery.Exec()
	myQuery, _ = myDb.Prepare("PRAGMA foreign_keys = ON;")
	myQuery.Exec()
	myQuery, _ = myDb.Prepare("INSERT INTO article (title, content) VALUES (?, ?)")
	myQuery.Exec("Rob", "Gronkowski")
	myQuery.Exec("KY", "Lee")
	myRecord, _ := myDb.Query("SELECT * FROM article")
	var id int
	var title string
	var content string
	var date_created string
	for myRecord.Next() {
		myRecord.Scan(&id, &title, &content, &date_created)
		fmt.Println(strconv.Itoa(id) + ": " + title + " " + content + " " + date_created)
	}
	myQuery, _ = myDb.Prepare("INSERT INTO comment (id_article, remark) VALUES (1, ?)")
	myQuery.Exec("My remarks")
}
