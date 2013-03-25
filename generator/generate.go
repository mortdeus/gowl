package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"strings"
	"sync"
	"text/template"
)

var wg = new(sync.WaitGroup)

func main() {
	generate()
}
func generate() {
	buf := new(bytes.Buffer)
	p := new(Protocol)
	f, err := os.Open("wayland.xml")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	if _, err := buf.ReadFrom(f); err != nil && err != io.EOF {
		panic(err)
	}
	if err := xml.Unmarshal(buf.Bytes(), p); err != nil {
		panic(err)
	}
	for _, in := range p.Interface {
		wg := new(sync.WaitGroup)
		fmtinator(&in, wg)
		wg.Wait()
		f1, err := os.Create("../" + in.Name + ".go")
		if err != nil {
			panic(err)
		}
		tmpl := template.Must(template.New("pkg").Parse(pkgTemplate))
		if err := tmpl.Execute(f1, in); err != nil {
			panic(err)
		}

	}

}

func fmtinator(typ interface{}, wg *sync.WaitGroup) {
	f := func(s string) string {
		if strings.HasPrefix(s, "wl_") {
			s = s[3:]
		}
		if s == "class_" {
			s = "class"
		}
		str := strings.Split(s, "_")
		for i, in := range str {
			str[i] = strings.ToUpper(string(in[0])) + in[1:]

		}
		return strings.Join(str, "")
	}
	wg.Add(1)
	switch p := typ.(type) {
	case *Interface:
		p.Name = f(p.Name)
		wg1 := new(sync.WaitGroup)
		for i := range p.Event {
			go fmtinator(&p.Event[i], wg1)
		}
		for i := range p.Request {
			go fmtinator(&p.Request[i], wg1)
		}
		for i := range p.Enum {
			go fmtinator(&p.Enum[i], wg1)
		}
		wg1.Wait()
	case *Event:
		p.Name = f(p.Name)
		for i := range p.Arg {
			go fmtinator(&p.Arg[i], wg)
		}
	case *Request:
		p.Name = f(p.Name)
		for i := range p.Arg {
			go fmtinator(&p.Arg[i], wg)
		}
	case *Enum:
		p.Name = f(p.Name)
		for i := range p.Entry {
			go fmtinator(&p.Entry[i], wg)
		}
	case *Arg:
		p.Name = f(p.Name)

	case *Entry:
		p.Name = f(p.Name)
	}
	wg.Done()
}
