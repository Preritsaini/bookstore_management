package routes

import (
	"awesomeProject2/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")                // Get all books
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")   // Get book by ID
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")            // Create a book
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")    // Update a book
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE") // Delete a book
}
