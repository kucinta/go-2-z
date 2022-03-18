package main

import "fmt"

func main() {
	// interface{} is an 'empty interface'
	// can handle all kinds of variables
	userInfo := make(map[string]interface{})
	userInfo["id"] = 123
	userInfo["name"] = "Zhang San"
	userInfo["age"] = 25 + 25
	fmt.Println(userInfo)

	userInfo2s := make([]map[string]string, 5)
	for i := 0; i < 3; i++ {
		userInfo2s[i] = make(map[string]string, 5)
		userInfo2s[i]["id"] = "123"
		userInfo2s[i]["name"] = "name..."
		userInfo2s[i]["age"] = "Age..."
	}
	for _, abc := range userInfo2s {
		fmt.Println(abc["age"])
	}
	fmt.Println(userInfo2s)
	userInfo2s[1]["id"] = "ID123"
	userInfo2s[1]["age"] = "99"
	userInfo2s[1]["name"] = "ABC"
	fmt.Println(userInfo2s[1])

	foods := map[string]interface{}{
		"bacon": "delicious",
		"price": 1 + 11,
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}
	fmt.Println(foods["price"])
}
