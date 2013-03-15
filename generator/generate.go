package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	//"text/template"
)

func main() {
	importProtocol()
}

func importProtocol() {
	f, err := os.Open("wayland.xml")
	if err != nil {
		panic(err)
	}
	proto := new(Protocol)
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(f); err != nil && err != io.EOF {
		panic(err)
	}

	if err := xml.Unmarshal(buf.Bytes(), proto); err != nil {
		panic(err)
	}
}
