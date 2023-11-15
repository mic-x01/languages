package main

import (
	"encoding/json"
	// "fmt"
	// "io"
	"strings"
)

func main() {
	reader := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	m := map[string]float64{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&m)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Map: %T, %v", m, m)
		for k, v := range m {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}
}
