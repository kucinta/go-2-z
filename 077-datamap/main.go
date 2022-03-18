package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	//Test1()
	//Test2()
	Test3()
}
func Test3() {
	userInfo2s := make([]map[string]interface{}, 10)
	for i := 0; i < 10; i++ {
		userInfo2s[i] = make(map[string]interface{}, 10)
		userInfo2s[i]["id"] = 201902181425
		userInfo2s[i]["name"] = "name..."
		userInfo2s[i]["age"] = "Age..."
	}
	fmt.Println(userInfo2s)	
}
type MyStruct struct {
    Id int64
    Name string
    Price float32
}

func Test2() {

    dict := make(map[string]interface{})
    dict["id"] = 201902181425       // int64
    dict["name"] = "jackytse"       // string
    dict["price"] = 123.456 + 100   // float32

    jsonbody, err := json.Marshal(dict)
	checkErr(err)
	// jsonbody in bytes
	fmt.Printf("%#v\n", jsonbody)
	// convert bytes to string
	fmt.Printf("%#v\n", string(jsonbody))


    myStruct := MyStruct{}
	
	err = json.Unmarshal(jsonbody, &myStruct)
	checkErr(err)

    fmt.Printf("%#v\n", myStruct)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


func Test1() {
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