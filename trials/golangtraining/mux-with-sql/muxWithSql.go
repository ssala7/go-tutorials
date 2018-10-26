package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Book struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Name   string `json:"name"`
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	//	r.HandleFunc("/books/{id}", GetOne).Methods("GET")
	r.HandleFunc("/books", GetAll).Methods("GET")
	r.HandleFunc("/books", Insert).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:PassWord123@tcp(localhost:32777)/booksdb")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var books Book
		// for each row, scan the result into our tag composite object
		err = rows.Scan(&books.Id, &books.Author, &books.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		w.Write([]byte(books.Author))
	}

}

func Insert(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:PassWord123@tcp(localhost:32777)/booksdb")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(r.Body)
	var book Book // create a var of Book
	json.Unmarshal(b, &book)
	q := `insert into books(author, name) values('%s', '%s')`
	f := fmt.Sprintf(q, book.Author, book.Name)
	fmt.Println("Going to run query: ", f)
	insert, err := db.Query(f)
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
