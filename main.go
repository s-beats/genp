package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"
)

type MethodDefinition struct {
	Name    string
	Args    []string
	Returns []string
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

	m := map[string]interface{}{
		"Name": "test",
		"Age":  30,
	}

	buf := bytes.NewBuffer(nil)

	if err := tpl.Execute(buf, m); err != nil {
		log.Fatal("Failed to execute template: ", err)
	}
	ioutil.WriteFile("test.go", buf.Bytes(), 0644)
}
