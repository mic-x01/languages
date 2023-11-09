package main

import (
	"io"
)

func processData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)
	ConsumeData(pipeReader)

}
