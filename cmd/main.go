package main

import (
	"encoding/json"
	"fmt"
	"github.com/MaksKazantsev/go-crud/internal/helper"
	"github.com/MaksKazantsev/go-crud/internal/model"
	"math/rand"
	"net/http"
	"strings"
)

// Book struct with ID, Title, Author with string DT
// Declaring a variable called books, with book DT
var books []model.Book

// This is a router function, because i was trying not to use libs, i had to write a switch and case to check if the method is right,
// and if it is, than we are calling a function for each case
func mainFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/books/":
			getBooks(w, r)
		default:
			getBook(w, r)
		}
	case http.MethodPost:
		createBook(w, r)
	case http.MethodPut:
		updateBook(w, r)
	case http.MethodDelete:
		deleteBook(w, r)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")       // Setting up a Content-Type - json
	if err := json.NewEncoder(w).Encode(books); err != nil { // Encoding the response and handling an error, it is important to handle all errors
		fmt.Println("Error!")
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	if strings.ContainsAny(id, "0123456789") {
		for _, item := range books {
			if fmt.Sprintf("%d", item.ID) == id {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bookItem model.Book
	_ = json.NewDecoder(r.Body).Decode(&bookItem)
	bookItem.ID = rand.Intn(100000)
	books = append(books, bookItem)
	json.NewEncoder(w).Encode(bookItem)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	if strings.ContainsAny(id, "0123456789") {
		for index, item := range books {
			if fmt.Sprintf("%d", item.ID) == id {
				var body model.Book
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				body.ID = item.ID
				books[index] = body
				json.NewEncoder(w).Encode(books)
				return
			}
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	if strings.ContainsAny(id, "0123456789") {
		for index, item := range books {
			if fmt.Sprintf("%d", item.ID) == id {
				books = append(books[:index], books[index+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(books)
	}
}

// Creating a value method for handler struct, basically this method serves an HTTP response and request and then colling a router function
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mainFunc(w, r)
}

// Handler Struct
type handler struct{}

func main() {
	fmt.Println("Starting server!")
	// Appending an example (not required)
	books = append(books, model.Book{ID: 1, Title: "Example", Author: "Example Author"})

	// Setting up some stuff about HTTP server, such as Address and HandlerFunction
	srv := http.Server{
		Addr:    ":8000",
		Handler: handler{},
	}

	// Launching server, using srv struct and handling an error
	err := srv.ListenAndServe()
	helper.PanicIfErr(err)
}
