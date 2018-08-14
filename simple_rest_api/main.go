package main

//simple rest-api
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	// '*' allows us to access a pointer fields that  we will use when we create a struct variable
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
type RequestMessage struct {
	Status string `json:"status"`
}

// books
var books []Book

func main() {
	//test data books @todo implement database
	// we use a pointer field author of struct Book to struct Author
	books = append(books, Book{Id: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{Id: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Jane", Lastname: "Doe"}})

	//create router
	r := mux.NewRouter()
	// endpoints and router handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/get-book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/create-book", createBook).Methods("POST")
	r.HandleFunc("/api/update-book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/delete-book/{id}", deleteBook).Methods("DELETE")

	//lets start our server
	log.Fatal(http.ListenAndServe(":8000", r))

}
