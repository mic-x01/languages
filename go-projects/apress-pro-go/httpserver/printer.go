package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func Waitkey() {
	fmt.Println("Press any key...")
	fmt.Scanln()
}
