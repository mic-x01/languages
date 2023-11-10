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

func writeReplaced(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(writer, str)
}

func main() {
	text := "It was a boat. A small boat."
	subs := []string{"boat", "kayak", "small", "huge"}
	var writer strings.Builder
	writeReplaced(&writer, text, subs...)
	fmt.Println(writer.String())
}
