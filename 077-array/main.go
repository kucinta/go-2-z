package main

import "fmt"

type car struct {
	make  string
	model string
	year  int
}

func Collect(list []car, f func(car) string) []string {
	result := make([]string, len(list))
	for i, item := range list {
		result[i] = f(item)
	}
	return result
}

func main() {
	var cars = []car{
		{make: "Subaru", model: "Outback", year: 2007},
		{make: "Saab", model: "Sonett III", year: 1973},
		{make: "Lexus", model: "SC 430", year: 2009},
		{make: "Volvo", model: "V90", year: 1992},
	}

	for cars() {

	}
	getMake := func(c car) string {
		return c.make
	}
	fmt.Println(Collect(cars, getMake))

	getModel := func(c car) string {
		return c.model
	}
	fmt.Println(Collect(cars, getModel))

}
