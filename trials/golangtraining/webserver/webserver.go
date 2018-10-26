package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func bye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/bye", bye)
	http.ListenAndServe(":8000", nil)
}
