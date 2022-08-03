package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"

	"golang.org/x/tools/imports"

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
				Args: []*internal.Arg{
					{
						Name: "name",
						Type: "string",
					},
					{
						Name: "endedAt",
						Type: "time.Time",
					},
				},
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

	bytes, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		log.Fatal("Failed to format and adjust imports source: ", err)
	}

	ioutil.WriteFile("test.go", bytes, 0644)
}

var _ = imports.Options{}
