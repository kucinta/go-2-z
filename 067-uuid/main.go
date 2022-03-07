package main

import (
	"fmt"
	"crypto/rand"
	"log"
)

/*
UUID Generator uses the rand.Read function from package crypto/rand
to generate a basic UUID.

A universally unique identifier (UUID), or globally unique identifier (GUID), is
a 128-bit number used to identify information. A UUID is for practical purposes
unique: the probability that it will be duplicated is very close to zero.

UUIDs don’t depend on a central authority or on coordination between those 
generating them. The string representation of a UUID consists of 32 hexadecimal
digits displayed in 5 groups separated by 4 hyphens. 

The length of each group is: 8-4-4-4-12. 

UUIDs are fixed length. For example: 
123e4567-e89b-12d3-a456-426655440000
b2108443-36c0-1493-610f-d140725f9638
*/

func main() {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println(uuid)
}