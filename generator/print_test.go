package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	//"text/template"
	"testing"
)

func TestPrintProtocol(t *testing.T) {
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
	for _, i := range proto.Interface {
		fmt.Println("\\******************INTERFACE***********************\\")
		fmt.Printf("\n\t\t[%s]\n%s\n", i.Name, i.Description)

		if len(i.Request) >= 1 {
			fmt.Println("\\******************REQUEST***********************\\")
			for _, j := range i.Request {
				fmt.Printf("\n\t\t[%s]\n%s\nArgs:\t%s\n", j.Name, j.Description, j.Arg)
			}
		}

		if len(i.Event) >= 1 {
			fmt.Println("\\******************EVENT**************************\\")

			for _, k := range i.Event {
				fmt.Printf("\n\t\t[%s]\n%s\nArgs:\t%s\n", k.Name, k.Description, k.Arg)
			}
		}

		if len(i.Enum) >= 1 {
			fmt.Println("\\*******************ENUM****************************\\")
			for _, l := range i.Enum {
				fmt.Printf("\n\t\t[%s]\n%s\nEntry:\t%s\n", l.Name, l.Description, l.Entry)
			}
		}
	}
}
