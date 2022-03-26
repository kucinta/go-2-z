package main

import(
	"fmt"
	"log"
	"context"
	"database/sql"
	"text/template"
	"net/http"
	// Supports Older MSSQL Version 8 
	_ "github.com/minus5/gofreetds"
)

func main(){
	http.HandleFunc("/hello", hello)
	log.Println("Listening on http://0.0.0.0:8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}

type ViewData struct {
	Records []structAccount
}

type structAccount struct {
	AcctCode 	string
	Description	string
}


func myGet() ViewData {
	db := dbConn()
	defer db.Close()
	myRow, err := db.Query("SELECT AcctCode, Description FROM L_GLAccount ORDER BY AcctCode")
	checkErr(err)
	myAccount := structAccount{}
	myResults := []structAccount{}
	var AcctCode, Description string
	for myRow.Next() {
		err = myRow.Scan(&AcctCode, &Description)
		checkErr(err)
		myAccount.AcctCode = AcctCode
		myAccount.Description = Description
		myResults = append(myResults, myAccount)
		log.Println(AcctCode, Description)
	}
	vd :=ViewData{Records: myResults}
	return vd
}

func dbConn() (db *sql.DB) {
	dbHost := "192.168.1.201:1433"
	dbUser := "sa"
	dbPass := "welcome"
	dbName := "dbaoml"
	connStr :="host=" + dbHost + ";database=" + dbName + ";user=" + dbUser + ";pwd=" + dbPass + ";"
	dbDriver := "mssql"
	// mysql 
	// connStr := "dbUser+":"+dbPass+"@("+dbHost+")/"+dbName
	// dbDriver := "mysql"
	db, errdb := sql.Open(dbDriver, connStr)
	checkErr(errdb)
	ctx := context.Background()
	err := db.PingContext(ctx)
	checkErr(err)
	fmt.Println("Successfully connected to " + dbName)
	return db
}

func checkErr(err error) {
	if err != nil {
		log.Println("Error:", err.Error())
		panic(err)
	}
}

//var tmpl = template.Must(template.ParseGlob("templates/*.html"))

var myTmplvar = `
{{range .Records}}
    {{.AcctCode}} = {{.Description}}<br/>
  </div>
{{end}}
</p>
`

func hello(w http.ResponseWriter, r *http.Request) {
	vd:=myGet()
	usrTmpl,err := template.New("").Parse(myTmplvar)
	checkErr(err)
	w.Header().Set("Content-Type", "text/html")
	err2 := usrTmpl.Execute(w, vd)
	if err2 != nil {
		w.Write([]byte("<p>Internal Error</p>"))
		http.Error(w, err2.Error(), http.StatusInternalServerError)
	}
}