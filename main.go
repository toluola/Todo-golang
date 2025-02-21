package main

import (
    "log"
    "net/http"

    "todo-app/db"
    "todo-app/handlers"

    "github.com/gorilla/mux"
)

func main() {
    // Initialize database
    db.InitDB()

    // Create router
    r := mux.NewRouter()

    // Routes
    r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
    r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
    r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
    r.HandleFunc("/todos/{id}/done", handlers.MarkTodoDone).Methods("PUT")
    r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
    r.HandleFunc("/todos/{id}", handlers.GetTodo).Methods("GET")

    // Start server
    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
} 