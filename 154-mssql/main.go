package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type L_Item struct {
	ItemCode string
	ItemName string
}

func main() {
	// TLS Handshake failed: tls: server selected unsupported - Go V1.18 diable TLS1.0 and 1.1 only 1.2
	// use encrypt=disable; to workaround it.

	dsn := "server=192.168.1.10;user id=sa;password=welcome;port=1433;database=dbomomcch;encrypt=disable;"
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT ItemCode, ItemName FROM L_Item")
	if err != nil {
		fmt.Println("Cannot query: ", err.Error())
		return
	}
	//myL_Item := L_Item{}
	defer rows.Close()
	for rows.Next() {
		var ItemCode, ItemName string
		err = rows.Scan(&ItemCode, &ItemName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(ItemName, ItemCode)
	}

}
