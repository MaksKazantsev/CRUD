package routes

import "github.com/gorilla/mux"
import "github.com/MaksKazantsev/go-crud/internal/controllers"

var RegisterRoutes = func(r *mux.Router) {
	r.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{ID}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{ID}", controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/{ID}", controllers.UpdateBook).Methods("PUT")
}
