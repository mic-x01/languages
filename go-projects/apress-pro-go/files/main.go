package main

import (
	// "fmt"
	"os"
	"path/filepath"
	// "path/filepath"
	// "time"
	// "encoding/json"
)

func callback(path string, dir os.DirEntry, dirErr error) (err error) {
	info, _ := dir.Info()
	Printfln("Path %v, Size: %v", path, info.Size())
	return
}

func main() {
	path, err := os.Getwd()
	if err == nil {
		err = filepath.WalkDir(path, callback)
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
}
