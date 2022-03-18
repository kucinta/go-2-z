package main

import (
	"database/sql"
	"fmt"

	md "github.com/JohannesKaufmann/html-to-markdown"
	_ "github.com/go-sql-driver/mysql"
)



func main() {


		myId := 135951
		fmt.Println(myId)
		db1, err := sql.Open("mysql", "root:welcome@tcp(192.168.1.222:3306)/mydbgo")
		checkErr(err)
		var post_content string
		err = db1.QueryRow("SELECT post_content FROM dbedu_posts where ID = ?", myId).Scan(&post_content)
		checkErr(err)
		fmt.Println(post_content)
		// update
		stmt, err := db1.Prepare("update dbedu_posts set post_content_filtered=? where ID=?")
		checkErr(err)
		converter := md.NewConverter("", true, nil)
		//html := `<html><h1>hello</h1><h3>world</h3><p>This is <strong>Important</strong></p><p>The End</p></html>`
		myContent, err := converter.ConvertString(post_content)
		fmt.Println(myContent)
		checkErr(err)
		res, err := stmt.Exec(myContent, myId)
		checkErr(err)
		affect, err := res.RowsAffected()
		checkErr(err)
		fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}