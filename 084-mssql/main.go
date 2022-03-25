package main

import (
	"database/sql"
	"fmt"
	"context"
	// Support new MSSQL Above Version 8
	_ "github.com/denisenkom/go-mssqldb" // the underscore indicates the package is used
)

func main() {
	connString := "sqlserver://sa:Welcome1@192.168.1.200:1433"
	db, err := sql.Open("sqlserver", connString)
	checkErr(err)
	ctx := context.Background()
	err = db.PingContext(ctx)
	checkErr(err)
	fmt.Printf("Connected!!\n")
	defer db.Close()
	myRow := db.QueryRow("SELECT AcctCode, Description FROM L_GLAccount WHERE AcctCode = '000-0001'")
	var firstName, lastName string
	err = myRow.Scan(&firstName, &lastName)
	checkErr(err)
	fmt.Printf("Results %s %s\n", firstName, lastName)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error open db:", err.Error())
		panic(err)
	}
	
}