package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	Printfln("Starting HTTP Server")
	go http.ListenAndServe(":5000", nil)

	time.Sleep(time.Second)

	Printfln("Press any key...")
	fmt.Scanln()
}
