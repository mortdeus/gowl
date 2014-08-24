package main

import (
	"encoding/xml"
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
	p := new(Protocol)
	f, err := os.Open("wayland.xml")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	if err := xml.NewDecoder(f).Decode(p); err != nil {
		panic(err)
	}
	for _, in := range p.Interface {
		wg := new(sync.WaitGroup)
		fmtinator(&in, wg)
		wg.Wait()
		f1, err := os.OpenFile(
			"zprotocol"+".go", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		if s, err := f1.Stat(); err != nil {
			panic(err)
		} else if s.Size() <= 0 {
			if _, err := f1.WriteString("package gowl"); err != nil {
				panic(err)
			}
		}

		tmpl := template.Must(template.New("pkg").Parse(pkgTemplate))
		if err := tmpl.Execute(f1, in); err != nil {
			panic(err)
		}
	}
}

func fmtinator(typ interface{}, wg *sync.WaitGroup) {
	wg.Add(1)
	switch p := typ.(type) {
	case *Interface:
		p.Name = stripAndCap(p.Name, false)
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
		p.Name = stripAndCap(p.Name, false)
		for i := range p.Arg {
			go fmtinator(&p.Arg[i], wg)
		}
	case *Request:
		p.Name = stripAndCap(p.Name, false)
		for i := range p.Arg {
			go fmtinator(&p.Arg[i], wg)
		}
	case *Enum:
		p.Name = stripAndCap(p.Name, false)
		for i := range p.Entry {
			go fmtinator(&p.Entry[i], wg)
		}
	case *Arg:
		p.Name = stripAndCap(p.Name, true)
		p.Type = stripAndCap(p.Type, false)

	case *Entry:
		p.Name = stripAndCap(p.Name, false)
	}
	wg.Done()
}

func stripAndCap(s string, mixedCase bool) string {
	if strings.HasPrefix(s, "wl_") {
		s = s[3:]
	}
	if s == "class_" {
		s = "class"
	}
	if s == "uint" || s == "int" {
		return string(s[:] + "32")
	}
	if s == "string" {
		return s
	}
	str := strings.Split(s, "_")
	for i, in := range str {
		if mixedCase && i == 0 {
			if in == "type" {
				str[i] = "type_"
			}
			if in == "interface" {
				str[i] = "interface_"
			}
			continue
		}
		str[i] = strings.ToUpper(string(in[0])) + in[1:]
	}
	return strings.Join(str, "")
}
