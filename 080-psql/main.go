package main

import(
	"fmt"
  	//"gorm.io/driver/postgres"
  	//"gorm.io/gorm"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {	
	insertRecord()
	deleteRecord()
	updateRecord()
}

func dbConn() (db *sql.DB) {
	dsn := "host=192.168.1.237 user=user1 password=123 dbname=postgres port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	checkErr(err)
	fmt.Println("db ok, ready!")
	return db
}

func deleteRecord() {
	db := dbConn()
	sqlStatement := `
	DELETE FROM contacts WHERE id = $1`
	id := 6
	res, err := db.Prepare(sqlStatement)
	checkErr(err)
	res1, err := res.Exec(id)
	checkErr(err)
	count, err := res1.RowsAffected()
	checkErr(err)
	fmt.Println(count, " record ID is deleted:", id)
	defer db.Close()
}

func updateRecord() {
	db := dbConn()
	sqlStatement := `
	UPDATE contacts SET FullName=$2, Email=$3 WHERE id = $1 
	RETURNING id,Email;
	`
	id := 4
	var email1 string
var id1 int

	 err := db.QueryRow(sqlStatement, id, "Ah Seng","123 ABC St").Scan(&id1,&email1)
	checkErr(err)
	//count, err := res.RowsAffected()
	//checkErr(err)
	fmt.Println(" record ID is updated:", id1, email1)
	defer db.Close()
}


func insertRecord() {
	db := dbConn()
	sqlStatement := `
	INSERT INTO contacts (PhoneNumber, FullName, Address, Email)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, "12349999", "me siam", "123 any st","123@me.com").Scan(&id)
	checkErr(err)
	fmt.Println("New record ID is:", id)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
