package main

import (
	"encoding/json"
	"io"
	// "fmt"
	"strings"
)

func main() {
	reader := strings.NewReader(`	
	{"Name":"Kayak","Category":"Watersports","Price":279}	
	{"Name":"Lifejacket","Category":"Watersports"}	
	{"name":"Canoe","category":"Watersports","price":100,"inStock":true}	
	`)
	decoder := json.NewDecoder(reader)
	for {
		var val Product
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v",
				val.Name, val.Category, val.Price)
		}
	}
}
