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
# Build test and execute binary
```console
$ go build
$ ls hello* -la
# -rwxr-xr-x 1 user user 1766414 Feb 22 14:56 hello
# -rw-r--r-- 1 user user      75 Feb 22 09:00 hello.go
$ ./hello
# Hello World
$ rm hello
```
# Run source without generating executable
```console
$ go run hello.go
```
