package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "time"

    "todo-app/db"
    "todo-app/models"

    "github.com/gorilla/mux"
)

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo models.Todo
    json.NewDecoder(r.Body).Decode(&todo)

    if todo.Title == "" {
        jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Title is required"})
        return
    }

    err := db.DB.QueryRow(
        "INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING id, created_at, updated_at",
        todo.Title, todo.Description,
    ).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)

    if err != nil {
        jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }

    jsonResponse(w, http.StatusOK, todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, title, description, done, created_at, updated_at FROM todos")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var todo models.Todo
        err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        todos = append(todos, todo)
    }

    json.NewEncoder(w).Encode(todos)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    var todo models.Todo
    json.NewDecoder(r.Body).Decode(&todo)

    _, err := db.DB.Exec(
        "UPDATE todos SET title=$1, description=$2, updated_at=$3 WHERE id=$4",
        todo.Title, todo.Description, time.Now(), id,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    todo.ID = id
    json.NewEncoder(w).Encode(todo)
}

func MarkTodoDone(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    _, err := db.DB.Exec(
        "UPDATE todos SET done=true, updated_at=$1 WHERE id=$2",
        time.Now(), id,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    _, err := db.DB.Exec("DELETE FROM todos WHERE id=$1", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    var todo models.Todo
    err := db.DB.QueryRow(
        "SELECT id, title, description, done, created_at, updated_at FROM todos WHERE id=$1", id,
    ).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt)

    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Todo not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(todo)
} 