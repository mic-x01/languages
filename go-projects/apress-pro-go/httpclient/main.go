package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	Printfln("Starting HTTP Server")
	go http.ListenAndServe(":5000", nil)

	time.Sleep(time.Second)
	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])
	if err == nil {
		reqURL, err := url.Parse("http://localhost:5000/echo")
		if err == nil {
			req := http.Request{
				Method: http.MethodPost,
				URL:    reqURL,
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(builder.String())),
			}
			response, err := http.DefaultClient.Do(&req)
			if err == nil && response.StatusCode == http.StatusOK {
				io.Copy(os.Stdout, response.Body)
				defer response.Body.Close()
			} else {
				Printfln("Error: %v", err.Error())
			}
		} else {
			Printfln("Parse Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}

	Printfln("Press any key...")
	fmt.Scanln()
}
