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
# install in $GOPATH/bin
```console
$ go install go-2-z/hello
# or
go install .
# or
go install
$ ~/go/bin/hello
```
# build
```console
$ go build
```
# execute
```console
$ ./hello
```
# run interactively
```console
$ go run hello.go
```
