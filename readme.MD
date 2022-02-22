# Go Workspace Setup
Inside settings.json, add...
```code
{
  ...
  "gopls": { "experimentalWorkspaceModule": true, },
  ...
}
```
# Initial Projct Setup
```console
$ mkdir ~/go-2-z
$ cd ~/go-2-z
$ go mod init go-2-z/hello
```
# Create Go source file
```console
$ nano hello.go
```
```go
package main
func main() {
    println("Hello World")
}
```
# Compile source and install binary in $GOPATH/bin
```console
$ go install go-2-z/hello
# or
go install .
# or
go install
$ ~/go/bin/hello
```
# Build test binary 
```console
$ go build
```
# Execute and delete test binary
```console
$ ./hello
# Hello World
$ rm hello
```
# Run source without generating executable
```console
$ go run hello.go
```
