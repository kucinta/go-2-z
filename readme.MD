# Workspace Setup
Inside settings.json, add...
```code
"gopls": { "experimentalWorkspaceModule": true, }
```
# setup
```console
$ mkdir ~/go-pco3
$ cd ~/go-pco3
$ mkdir -p cmd/app pkg bin
$ go mod init github.com/eldus/go-pmo3
```
# create go source file
```console
$ nano cmd/app/main.go
```
```go
package main
func main() {
    println("Hello World")
}
```
# build
```console
$ go build -o bin/app ./cmd/app
```
# execute
```console
$ bin/app
```
# run interactively
```console
$ go run cmd/app/main.go
```
