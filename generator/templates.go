package main

import (
	"encoding/xml"
	"strings"
)

type Protocol struct {
	XMLName   xml.Name `xml:"Protocol"`
	Copyright string   `xml:"Copyright"`
	Interface []Interface
}

type Interface struct {
	Name        string `xml:"name,attr"`
	Description string
	Request     []Request
	Event       []Event
	Enum        []Enum
}

type Request struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"Description"`
	Arg         []Arg
}

type Event struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"Description"`
	Arg         []Arg
}

type Enum struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"Description"`
	Entry       []Entry
}

type Entry struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Arg struct {
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr"`
	Interface string `xml:"interface,attr"`
	AllowNull string `xml:"allow-null,attr"`
}

var (
	pkgTemplate string = `package gowl
	
	{{if .Description}}/*
	{{.Description}}
	*/	{{end}}
	type {{.Name}} struct{}

		{{template "Enum"}}
		{{template "Request"}}
		{{template "Event"}}
	

	{{define "Header"}}
	func (*{{$.Name}}) .Name({{range .Args}}{{.Name}}{{.Type}},{{end}})
	{{end}}
	
	{{define "Request"}}
	{{range .Request}}
	{{template "Header"}}{

	}
	{{end}}
	{{end}}

	{{define "Event"}}
		{{range .Event}}
			{{template "Header"}}{

			}
		{{end}}
	{{end}}
	
	{{define "Enum"}}		
		{{range .Enum}}
			{{if .Description}}/*
			{{.Description}}
			*/{{end}}
			var {{.Name}} int
			const(
				{{range .Entry}}
				{{.Name}} = {{.Value}} 
				{{end}}
			)
		{{end}}
	{{end}}

	`
)

func stripAndCap(p *Protocol) {
	f := func(s string) string {
		if strings.HasPrefix(s, "wl_") {
			s = s[3:]

		}
		str := strings.Split(s, "_")

		for i, in := range str {
			str[i] = strings.ToUpper(string(in[0])) + in[1:]
		}
		return strings.Join(str, "")
	}
	for i, in := range p.Interface {
		p.Interface[i].Name = f(in.Name)
	}

}
