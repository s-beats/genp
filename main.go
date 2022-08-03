package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"text/template"
)

type Arg struct {
	Name string
	Type string
}

type Return struct {
	Type string
}

type MethodDefinition struct {
	Name    string
	Args    []*Arg
	Returns []*Return
}

type Definition struct {
	Interface string
	Methods   []*MethodDefinition
	Implement string
}

func main() {
	tpl, err := template.ParseFiles("test.tmpl")
	if err != nil {
		log.Fatal("Failed to generate template: ", err)
	}

	def := &Definition{
		Interface: "User",
		Methods: []*MethodDefinition{
			{
				Name:    "Create",
				Args:    []*Arg{{"name", "string"}},
				Returns: []*Return{{"error"}},
			},
		},
		Implement: "user",
	}

	buf := bytes.NewBuffer(nil)

	if err := tpl.Execute(buf, def); err != nil {
		log.Fatal("Failed to execute template: ", err)
	}

	bytes, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal("Failed to format source: ", err)
	}

	ioutil.WriteFile("test.go", bytes, 0644)
}
