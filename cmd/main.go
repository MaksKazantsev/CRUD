package main

import (
	"encoding/json"
	"fmt"
	"github.com/MaksKazantsev/go-crud/internal/helper"
	"github.com/MaksKazantsev/go-crud/internal/model"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

// Book struct with ID, Title, Author with string DT
// Declaring a variable called books, with book DT
var books []model.Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")       // Setting up a Content-Type - json
	if err := json.NewEncoder(w).Encode(books); err != nil { // Encoding the response and handling an error, it is important to handle all errors
		fmt.Println("Error!")
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if strings.ContainsAny(params["ID"], "0123456789") {
			if item.ID == params["ID"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		} else {
			http.Error(w, "Error - no string in URL!", http.StatusBadRequest)
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bookItem model.Book
	_ = json.NewDecoder(r.Body).Decode(&bookItem)
	bookItem.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, bookItem)
	json.NewEncoder(w).Encode(bookItem)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if strings.ContainsAny(params["ID"], "0123456789") {
		for index, item := range books {
			if item.ID == params["ID"] {
				books = append(books[:index], books[index+1:]...)
				var book model.Book
				_ = json.NewDecoder(r.Body).Decode(&book)
				book.ID = params["ID"]
				books = append(books, book)
				json.NewEncoder(w).Encode(book)
				return
			}
		}
	} else {
		http.Error(w, "Error - no string in URL!", http.StatusBadRequest)
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if strings.ContainsAny(params["ID"], "0123456789") {
		for index, item := range books {
			if item.ID == params["ID"] {
				books = append(books[:index], books[index+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(books)
	}
}

func main() {
	fmt.Println("Starting server!")
	r := mux.NewRouter()
	r.HandleFunc("/books/", getBooks).Methods("GET")
	r.HandleFunc("/books/{ID}", getBook).Methods("GET")
	r.HandleFunc("/books/", createBook).Methods("POST")
	r.HandleFunc("/books/{ID}", deleteBook).Methods("DELETE")
	r.HandleFunc("/books/{ID}", updateBook).Methods("PUT")

	books = append(books, model.Book{ID: "1", Title: "Example", Author: "Example Author"})

	srv := http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	// Launching server, using srv struct and handling an error

	err := srv.ListenAndServe()
	helper.PanicIfErr(err)
}
