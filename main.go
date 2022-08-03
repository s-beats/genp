package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/s-beats/genp/internal"
)

func main() {
	tpl, err := template.ParseFiles("test.tmpl")
	if err != nil {
		log.Fatal("Failed to generate template: ", err)
	}

	def := &internal.Definition{
		Interface: "User",
		Methods: []*internal.MethodDefinition{
			{
				Name: "Create",
				Args: []*internal.Arg{{
					Name: "name",
					Type: "string",
				}},
				Returns: []*internal.Return{{
					Type: "error",
				}},
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
