# Todo App

This is a simple Todo application built with Go, Gorilla Mux, and PostgreSQL. The application allows users to create, read, update, and delete todo items. It also provides an endpoint to mark a todo as done.

## Features

- Create a new todo
- Retrieve all todos
- Retrieve a single todo by ID
- Update a todo
- Mark a todo as done
- Delete a todo

## Endpoints

### Create a Todo

- **URL:** `/todos`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "title": "Sample Todo",
    "description": "This is a sample todo item."
  }
  ```
- **Response:** Returns the created todo item with its ID and timestamps.

### Get All Todos

- **URL:** `/todos`
- **Method:** `GET`
- **Response:** Returns a list of all todos.

### Get a Single Todo

- **URL:** `/todos/{id}`
- **Method:** `GET`
- **Response:** Returns the todo item with the specified ID.

### Update a Todo

- **URL:** `/todos/{id}`
- **Method:** `PUT`
- **Request Body:**
  ```json
  {
    "title": "Updated Title",
    "description": "Updated description."
  }
  ```
- **Response:** Returns the updated todo item.

### Mark Todo as Done

- **URL:** `/todos/{id}/done`
- **Method:** `PUT`
- **Response:** Returns a success status if the todo is marked as done.

### Delete a Todo

- **URL:** `/todos/{id}`
- **Method:** `DELETE`
- **Response:** Returns a success status if the todo is deleted.

## Setup Instructions

1. **Install Go:**
   Ensure you have Go installed on your system. You can download it from [golang.org/dl](https://golang.org/dl/).

2. **Clone the Repository:**
   ```bash
   git clone https://github.com/toluola/Todo-golang.git
   cd todo-app
   ```

3. **Set Up PostgreSQL:**
   - Ensure PostgreSQL is installed and running.
   - Create a database named `todo_db`.
   - Update the connection string in `db/db.go` with your PostgreSQL credentials.

4. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

5. **Run the Application:**
   ```bash
   go run main.go
   ```

6. **Access the Application:**
   The server will start on port 8080. You can access the endpoints using a tool like Postman or curl.

## Code Reference

- Main application entry point: `main.go`

- Handlers for todo operations: `handlers/todo_handler.go`

- Database initialization: `db/db.go`

- Todo model definition: `models/todo.go`