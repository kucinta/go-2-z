package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	fmt.Println("vim-go : 8067")
	http.HandleFunc("/", index)
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	//http.HandleFunc("/bad", http.NotFound)
	http.HandleFunc("/every-thing-else", http.NotFound)
	fmt.Println("Server started on: http://192.168.1.234:8067")
	http.ListenAndServe(":8067", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintf(w, "<h1>404 – Page Not Found!</h1><a href='/'>Go Home</a>")
		return
	}
	fmt.Fprintf(w, "<h1>Hello world!</h1><a href='/login'>Login</a> to <a href='/secret'>view /secret</a>.")
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "<h1>You are inside the secret garden!</h1><a href='/logout'>/logout to close garden</a>.")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
	fmt.Fprintln(w, "<p>You have Login</p><p><a href='/secret'>Visit Secret...</a></p>")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintf(w, "<p>Logout</p><p><a href='/'>Go Home</a></p>")
}
