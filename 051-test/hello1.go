package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("... hello1")
	fmt.Println(time.Now())
}

func displayTime() {
	fmt.Println("... hello2")
	fmt.Println(time.Now())
}
