package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/MaksKazantsev/go-crud/internal/model"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var books []model.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")       // Setting up a Content-Type - json
	if err := json.NewEncoder(w).Encode(books); err != nil { // Encoding the response and handling an error, it is important to handle all errors
		fmt.Println("Error!")
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
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

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bookItem model.Book
	_ = json.NewDecoder(r.Body).Decode(&bookItem)
	bookItem.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, bookItem)
	json.NewEncoder(w).Encode(bookItem)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
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

func DeleteBook(w http.ResponseWriter, r *http.Request) {
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
