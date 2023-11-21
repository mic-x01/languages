package main

import (
	// "fmt"
	"os"
	"path/filepath"
	// "path/filepath"
	// "time"
	// "encoding/json"
)

func main() {
	path, err := os.Getwd()
	if err == nil {
		matches, err := filepath.Glob(filepath.Join(path, "*.json"))
		if err == nil {
			for _, m := range matches {
				Printfln("Match: %v", m)
			}
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
}
