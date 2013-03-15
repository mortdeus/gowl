package main

var BaseTemplate string = "package wl\n" +
	"type {{.Name}} struct{}\n" +
	"{{range .Funcs}} func (*{{$.Name}})  {{.Ident}}" +
	"({{range .Args}}{{.Ident}} {{.Type}},{{end}})" +
	"({{range .Return}}{{.Type}},{{end}})" +
	"{\n" +
	"{{.Body}}" +
	"}\n" +
	"{{end}}"
