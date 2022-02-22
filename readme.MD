# Workspace Setup
Inside settings.json, add...
```code
"gopls": { "experimentalWorkspaceModule": true, }
```
# setup
```console
$ mkdir ~/go-2-z
$ cd ~/go-2-z
$ go mod init go-2-z/hello
```
# create go source file
```console
$ nano hello.go
```
```go
package main
func main() {
    println("Hello World")
}
```
# build
```console
$ go build -o hello 
```
# execute
```console
$ bin/app
```
# run interactively
```console
$ go run cmd/app/main.go
```
