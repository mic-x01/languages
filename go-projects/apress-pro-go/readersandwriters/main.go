package main

import (
	// "bufio"
	"fmt"
	// "io"
	"encoding/json"
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

// func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
// 	fmt.Fprintf(writer, template, vals...)
// }
//
// func writeReplaced(writer io.Writer, str string, subs ...string) {
// 	replacer := strings.NewReplacer(subs...)
// 	replacer.WriteString(writer, str)
// }

func main() {
	var b bool = true
	var str string = "Hello"
	var fval float64 = 99.99
	var ival int = 200
	var pointer *int = &ival
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	for _, val := range []interface{}{b, str, fval, ival, pointer} {
		encoder.Encode(val)
	}
	fmt.Print(writer.String())
}
