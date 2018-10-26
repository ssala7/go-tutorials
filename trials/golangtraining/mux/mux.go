package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Author string `json:Author,omitempty`
	Name   string `json:Name,omitempty`
	Id     string `json:id,omitempty`
}

var Books []Book

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Books)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	var book Book                   // create a var of Book
	err := json.Unmarshal(b, &book) // Unmarshal with a pointer
	if err != nil {
		fmt.Println(err)
		return
	}
	Books = append(Books, book)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range Books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	Books = append(Books, Book{Name: "Write presentation", Author: "Suresh", Id: "12345"})
	Books = append(Books, Book{Name: "Write presentation", Author: "Suresh", Id: "123"})
	Books = append(Books, Book{Name: "Write presentation", Author: "Vivek", Id: "456"})
	// Routes consist of a path and a handler function.
	r.HandleFunc("/books/{id}", GetOne).Methods("GET")
	r.HandleFunc("/books", GetAll).Methods("GET")
	r.HandleFunc("/books", Insert).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
