# Go VSCode Workspace Setup
Inside settings.json, add...
```code
{
  ...
  "gopls": { "experimentalWorkspaceModule": true, },
  ...
}
```
# $GOPATH Setup
```
$ nano ~/.bashrc
# --- add ---
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
# --- end ---
$ source .bashrc
$ echo $GOPATH
$ echo $PATH
```
# Initial Project Setup
```console
$ mkdir ~/go-2-z
$ cd ~/go-2-z
$ go mod init go-2-z/hello
$ git init
$ git branch -a
$ git add .
$ git status
```
# Update github
```console
$ git status
$ git add .
$ git commit -m "up 244"
$ git push
$ git status
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
# Compile source, install and execute binary in $GOPATH/bin
```console
$ go install go-2-z/hello
# or
$ go install .
# or
$ go install

$ $GOPATH/bin/hello
# or
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
