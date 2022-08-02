package main

import (
	"bytes"
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
		Implement: "User",
	}

	buf := bytes.NewBuffer(nil)

	if err := tpl.Execute(buf, def); err != nil {
		log.Fatal("Failed to execute template: ", err)
	}
	ioutil.WriteFile("test.go", buf.Bytes(), 0644)
}
