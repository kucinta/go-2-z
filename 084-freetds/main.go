package main

import(
	"fmt"
	"context"
	"database/sql"
	// Supports Older MSSQL Version 8 
	_ "github.com/minus5/gofreetds"
)
func main(){
	connStr :="host=192.168.1.201:1433;database=dbaoml;user=sa;pwd=welcome;"
	db, errdb := sql.Open("mssql", connStr)
	checkErr(errdb)
	ctx := context.Background()
	err := db.PingContext(ctx)
	checkErr(err)
	fmt.Printf("Connected!\n")
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
