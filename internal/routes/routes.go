package routes

import (
	"github.com/MaksKazantsev/go-crud/internal/server"
	"github.com/gorilla/mux"
)
import "github.com/MaksKazantsev/go-crud/internal/controllers"

const (
	BooksURL = "/books/"
	BookURL  = "/books/{ID}"
)

var RegisterRoutes = func(r *mux.Router, server *server.HTTPServer) {
	r.HandleFunc(BooksURL, controllers.GetBooks).Methods("GET")
	r.HandleFunc(BookURL, controllers.GetBook).Methods("GET")
	r.HandleFunc(BooksURL, controllers.CreateBook).Methods("POST")
	r.HandleFunc(BookURL, controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc(BookURL, controllers.UpdateBook).Methods("PUT")
}
