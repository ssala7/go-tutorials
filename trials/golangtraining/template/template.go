package main

import (
	"os"
	"text/template"
)

type Person struct {
	//exported field since it begins with a capital letter
	Name string
}

func main() {
	//create a new template with some name
	t := template.New("hello template")

	//parse some content and generate a template,
	//which is an internal representation
	t, _ = t.Parse("hello {{.Name}}!")

	//define an instance with required field
	p := Person{Name: "Mary"}

	//merge template ‘t’ with content of ‘p’
	t.Execute(os.Stdout, p)
}
