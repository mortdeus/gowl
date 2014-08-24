package main

import (
	"encoding/xml"
	"os"
	"strings"
	"text/template"
)

func main() {
	generate()
}
func generate() {
	f, err := os.Open("wayland.xml")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	var proto Protocol
	if err := xml.NewDecoder(f).Decode(&proto); err != nil {
		panic(err)
	}
	
	fmtinator(&proto)
	for _, in := range proto.Interface {

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

func fmtinator(typ interface{}) {
	switch p := typ.(type) {
	case *Protocol:
		for i := range p.Interface {
			fmtinator(&p.Interface[i])
		}

	case *Interface:
		p.Name = stripAndCap(p.Name, false)
		for i := range p.Event {
			fmtinator(&p.Event[i])
		}
		for i := range p.Request {
			fmtinator(&p.Request[i])
		}
		for i := range p.Enum {
			fmtinator(&p.Enum[i])
		}
	case *Event:
		p.Name = stripAndCap(p.Name, false)
		for i := range p.Arg {
			fmtinator(&p.Arg[i])
		}
	case *Request:
		p.Name = stripAndCap(p.Name, false)
		for i := range p.Arg {
			fmtinator(&p.Arg[i])
		}
	case *Enum:
		p.Name = stripAndCap(p.Name, false)
		for i := range p.Entry {
			fmtinator(&p.Entry[i])
		}
	case *Arg:
		p.Name = stripAndCap(p.Name, true)
		p.Type = stripAndCap(p.Type, false)

	case *Entry:
		p.Name = stripAndCap(p.Name, false)
	}
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
