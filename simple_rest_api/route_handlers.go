package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

/*
-route handlers
-takes 2 must params - response,request
---response http:ResponseWriter
---request *http:Request
*/

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	//set http header content-type to json
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		log.Fatal("Error:", err)
		_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Server Error"})
	}
	return

}

//get a book by id
func getBook(w http.ResponseWriter, r *http.Request) {
	//set http header content-type to json
	w.Header().Set("Content-Type", "application/json")
	//get params of request
	params := mux.Vars(r)
	//get book item by id
	for _, book := range books {
		if book.Id == params["id"] {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				log.Fatal("Error:", err)
				_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Server Error"})
			}
			return
		}
	}
	//return an empty array if id is not found
	//json.NewEncoder(w).Encode(&Book{})
	//lets return an error message

	_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Book with id:" + params["id"] + " does not exist"})
	return
}

//create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//create a book item
	var book Book
	//decode request to book item
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Unable to decode POST request"})
		return
	}
	book.Id = generatePseudoRandomBookId()
	books = append(books, book)
	_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Book created successfully"})
	return
}

//update a book by id
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//get book by id
	for i, book := range books {
		if book.Id == params["id"] {
			err := json.NewDecoder(r.Body).Decode(&books[i])
			if err != nil {
				log.Fatal("Error: ", err)
				_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Unable to decode request"})
				return
			}
			_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Book updated successfully"})
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Book with id:" + params["id"] + " does not exist"})
	return
}

//delete book by id
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, book := range books {
		if book.Id == params["id"] {
			//delete book from books
			//memory leak can occur if books has pointers @todo implement fix for memory leak
			books = append(books[:i], books[i+1:]...)
			_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Book deleted successfully"})
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&RequestMessage{Status: "Book with id:" + params["id"] + " does not exist"})
	return

}

func generatePseudoRandomBookId() string {
	//Note: generates random id of range 1 - 100,000 hence ids may not be unique
	//@todo : id incrememnt method instead of random ids
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return strconv.Itoa(rng.Intn(100000))
}
