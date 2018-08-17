package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	context := getData()
	err := tmpl.Execute(w, context)
	if err != nil {
		log.Println("Error:", err)
	}

}

func makeServer() {
	muxRouter := mux.NewRouter()

	//route handlers
	muxRouter.HandleFunc("/", indexHandle).Methods("GET")
	muxRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	s := http.Server{
		Addr:    ":8000",
		Handler: handlers.LoggingHandler(os.Stdout, muxRouter),
	}
	log.Fatal(s.ListenAndServe())
}
