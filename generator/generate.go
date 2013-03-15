package main

import (
	//"encoding/xml"
	"os"
	"text/template"
)

type Interface struct {
	Name  string
	Funcs []handler
}
type handler struct {
	Ident  string
	Args   []vars
	Return []vars
	Body   string
}
type vars struct {
	Ident string
	Type  string
}

func main() {
	t := template.Must(template.New("t1").Parse(BaseTemplate))

	t.Execute(os.Stdout,
		Interface{"Foo",
			[]handler{{
				"Bar",
				[]vars{{"Baz", "int"}, {"Boo", "float"}},
				[]vars{{"", "interface{}"}, {"", "*int"}}, ""}, {
				"Bee",
				[]vars{{"Boom", "float32"}},
				[]vars{{"", "error"}}, ""}}})
}
