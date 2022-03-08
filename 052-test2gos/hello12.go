package main

import (
	"fmt"
	"time"
)

// Cannot declare main() here as it will flagged as redeclared
//func main() {
//	PrintTime() 
//}

func PrintTime() {
	fmt.Println(time.Now())
}

func printHello1() {
	fmt.Println("... hello1")
	PrintTime() 
}

func printHello2() {
	fmt.Println("... hello2")
	PrintTime() 
}

func displayTime() {
	fmt.Println(time.Now())
}