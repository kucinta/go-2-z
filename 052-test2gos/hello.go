package main
import "fmt"

/*
Run once:
$ go mod init 052-test2gos
$ go mod tidy
Cannot run: 
$ go run hello.go
Must run as: 
$ go run .
OR 
$ go run hello.go hello12.go
*/

func main() {
	fmt.Println("Hello, 2 Gos World.")
	PrintTime()
	printHello1()
	printHello2()
}