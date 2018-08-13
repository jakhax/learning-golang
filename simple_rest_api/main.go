package main

//simple rest-api
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id    string `json:"Id"`
	Isbn  string `json:"Isbn"`
	Title string `json:"Title"`
	// '*' allows us to access a pointer fields that  we will use when we create a struct variable
	Author *Author `json:"Author"`
}

type Author struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

/*
-route handlers
-takes 2 must params - response,request
---response http:ResponseWriter
---request *http:Request
*/

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

//get a book by id
func getBook(w http.ResponseWriter, r *http.Request) {

}

//create a book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//update a book by id
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//delete book by id
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//create router
	r := mux.NewRouter()
	// endpoints and router handlers
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	//lets start our server
	log.Fatal(http.ListenAndServe(":8000", r))

}
