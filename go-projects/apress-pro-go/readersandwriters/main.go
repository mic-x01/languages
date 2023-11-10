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
	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)
	fmt.Print(writer.String())
}
