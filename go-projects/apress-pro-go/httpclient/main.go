package main

import (
	"encoding/json"
	"fmt"
	// "io"
	"net/http"
	// "os"
	"time"
)

func main() {
	Printfln("Starting HTTP Server")
	go http.ListenAndServe(":5000", nil)

	time.Sleep(time.Second)
	response, err := http.Get("http://localhost:5000/json")
	if err == nil && response.StatusCode == http.StatusOK {
		defer response.Body.Close()
		data := []Product{}
		err = json.NewDecoder(response.Body).Decode(&data)
		if err == nil {
			for _, p := range data {
				Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
			}
		} else {
			Printfln("Decode error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	}

	Printfln("Press any key...")
	fmt.Scanln()
}
