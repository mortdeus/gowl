package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"reflect"
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
		fmt.Fprintln(os.Stderr, "Error:", err.Error())
		os.Exit(-1)
	}

	var proto Protocol
	if err := xml.NewDecoder(f).Decode(&proto); err != nil {
		fmt.Fprintln(os.Stderr, "Decode Error:", err.Error())
		os.Exit(-1)
	}

	fmtinator(&proto)

	of := os.Stdout
	fmt.Fprintln(of, "package gowl")

	for _, iface := range proto.Interface {
		tmpl := template.Must(template.New("pkg").Parse(pkgTemplate))
		if err := tmpl.Execute(of, iface); err != nil {
			fmt.Fprintln(os.Stderr, "Template Error:", err.Error())
			os.Exit(-1)
		}
	}
}

func fmtinator(typ interface{}) {
	switch p := typ.(type) {
	case *Protocol:
		for _, i := range p.Interface {
			fmtinator(i)
		}

	case *Interface:
		p.Name = stripAndCap(p.Name, false)
		for _, i := range p.Event {
			fmtinator(i)
		}
		for _, i := range p.Request {
			fmtinator(i)
		}
		for _, i := range p.Enum {
			fmtinator(i)
		}
	case *Event:
		p.Name = stripAndCap(p.Name, false)
		for _, i := range p.Arg {
			fmtinator(i)
		}
	case *Request:
		p.Name = stripAndCap(p.Name, false)
		for _, i := range p.Arg {
			fmtinator(i)
		}
	case *Enum:
		p.Name = stripAndCap(p.Name, false)
		for _, i := range p.Entry {
			fmtinator(i)
		}
	case *Arg:
		p.Name = stripAndCap(p.Name, true)
		p.Type = stripAndCap(p.Type, false)

	case *Entry:
		p.Name = stripAndCap(p.Name, false)

	default:
		fmt.Fprintf(os.Stderr, "Error: Trying to format unknown type (%s)\n", reflect.TypeOf(p))
		os.Exit(-1)
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
