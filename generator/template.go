package main

import (
	"encoding/xml"
)

type Protocol struct {
	XMLName   xml.Name    `xml:"protocol"`
	Copyright string      `xml:"copyright"`
	Interface []Interface `xml:"interface"`
}

type Interface struct {
	Name        string    `xml:"name,attr"`
	Description string    `xml:"description"`
	Request     []Request `xml:"request"`
	Event       []Event   `xml:"event"`
	Enum        []Enum    `xml:"enum"`
}

type Request struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description"`
	Arg         []Arg  `xml:"arg"`
}

type Event struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description"`
	Arg         []Arg  `xml:"arg"`
}

type Enum struct {
	Name        string  `xml:"name,attr"`
	Description string  `xml:"description"`
	Entry       []Entry `xml:"entry"`
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
	pkgTemplate string = `
	{{if .Description}}/*
	{{.Description}}
	*/{{end}}
	type {{.Name}} interface{
		{{range $.Request}}
			{{if .Description}}
				/*
				{{.Description}}
				*/
			{{end}}
			{{.Name}}({{range .Arg}}{{.Name}} {{.Type}},{{end}})
		
		{{end}}
		{{range $.Event}}
			{{if .Description}}
				/*
				{{.Description}}
				*/
			{{end}}
			{{.Name}}({{range .Arg}}{{.Name}} {{.Type}},{{end}})
			
		{{end}}	
	}

		{{range $e := $.Enum}}
			{{if .Description}}
				/*
				{{.Description}}
				*/
			{{end}}
			
			const(
				{{range .Entry}}{{$e.Name}}_{{.Name}} = {{.Value}};{{end}}
			)
		{{end}}

		
	
	`
)
