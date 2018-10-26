// package main
//
// import (
// 	"os"
// 	"text/template"
// )
//
// type Book struct {
// 	//exported field since it begins with a capital letter
// 	Author string
// 	Name   string
// }
//
// func main() {
// 	//create a new template with some name
// 	t := template.New("Author : template , Book Name : template")
//
// 	//parse some content and generate a template,
// 	//which is an internal representation
// 	t, _ = t.Parse("Author  : {{.Author}} , Book Name : {{.Name}}!")
//
// 	//define an instance with required field
// 	p := Book{Author: "Suresh", Name: "Terraform"}
//
// 	//merge template ‘t’ with content of ‘p’
// 	t.Execute(os.Stdout, p)
// }

package main

import (
	"os"
	"text/template"
)

func main() {
	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
	}

	var strs []string
	strs = append(strs, "test1")
	strs = append(strs, "test2")

	tmpl, err := template.New("test").Funcs(funcMap).Parse(`{{range $index, $element := .}}
  Number: {{inc $index}}, Text:{{$element}}
{{end}}`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, strs)
	if err != nil {
		panic(err)
	}
}
