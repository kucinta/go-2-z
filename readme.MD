# Go Versions
Go 1.18 [Release Notes](https://go.dev/doc/go1.18)
```code
V1.18 - 2022.03
V1.17 - 2021.08
...
V1.10 - 2018.02
V1.1 - 2013.05
V1. - 2012.03
```
# Go VSCode Workspace Setup
Inside VSCode settings.json, add the followings...
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
$ mkdir ~/go/go-2-z
$ cd ~/go/go-2-z
$ git init
$ git branch -a
$ git add .
$ git status
```
# Set Github username + password
```console
$ vim ~/.gitconfig
[user]
	email = <id>@<email-host>
	name = kucinta.com
[credential]
	helper = store
$ vim ~/.git-credentials
https://kucinta:<generated_token>@github.com
```
# Clone it
```
$ cd ~/go
$ git clone https://github.com/kucinta/go-2-z.git
$ cd ~/go/go-2-z
$ git status
```
# Update (Upload) to Github
```console
$ cd ~/go/go-2-z
$ touch test123.txt
$ git status
$ git pull
$ git add .
$ git commit -m "up by me"
$ git push
$ git status
```
# Update (Download) from Github
```console
$ git status
$ git pull
$ git status
```
# Create Go source file
```console
$ mkdir ~/go/go-2-z/051-test
$ cd ~/go/go-2-z/05
$ vim hello.go
```go
package main
func main() {
    println("Hello World")
}
```
# Build test and execute binary
```console
$ go build
$ ls hello* -lah
# -rwxr-xr-x  1 user user 1.7M Mar 21 09:52 hello
# -rw-r--r--  1 user user  113 Mar  8 13:26 hello.go
$ ./hello
# Hello World
$ rm hello
```
# Build smaller binary
By default Go build produce a binary file that contain debugging information and symbol tables. These can be removed to reduce the binary size.
```console
$ go build -ldflags "-w -s" hello.go
$ ls hello* -lah
# -rwxr-xr-x  1 user user 1.2M Mar 21 09:53 hello
# -rw-r--r--  1 user user  113 Mar  8 13:26 hello.go
```
The file is about 30% smaller.
# Run source without generating executable
```console
$ go run hello.go
```
# Compile source, install and execute binary in $GOPATH/bin
```console
$ go install hello.go
# or
$ go install .
# or
$ go install

$ $GOPATH/bin/hello
# or
$ ~/go/bin/hello
```
