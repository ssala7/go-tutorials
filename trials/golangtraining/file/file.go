package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
	Author string
	Name   string
}

func main() {
	filename := `books.json`
	p := Book{Author: "Suresh", Name: "HelloWorld"}
	b, _ := json.Marshal(p)
	ioutil.WriteFile(filename, b, 0755)
	f, _ := ioutil.ReadFile(filename)
	fmt.Println(string(f))
}
