package main

import (
	// "bufio"
	"fmt"
	"io"
	"strings"
)

// func scanFromReader(reader io.Reader, template string,
// 	vals ...interface{}) (int, error) {
// 	return fmt.Fscanf(reader, template, vals...)
// }
//
// func scanSingle(reader io.Reader, val interface{}) (int, error) {
// 	return fmt.Fscan(reader, val)
// }

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}

func main() {
	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"
	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
	fmt.Println(writer.String())
}
