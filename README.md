

# Gin CRUD API

A simple CRUD API built using the [Gin framework](https://github.com/gin-gonic/gin) in Go. This API allows basic operations to manage books, including creating, reading, updating, and deleting books. 

## Features

- **Create**: Add a new book to the collection.
- **Read**: Get all books or a specific book by its ID.
- **Update**: Modify an existing book's information.
- **Delete**: Remove a book from the collection.

## Prerequisites

Before running the application, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.18 or later)
- [Gin Framework](https://github.com/gin-gonic/gin)
- [Postman](https://www.postman.com/) or [curl](https://curl.se/) for API testing (optional)

## Setup

### 1. Clone the repository
Clone the project to your local machine:

```bash
git clone https://github.com/your-username/gin-crud-api.git
cd gin-crud-api
```

### 2. Install Dependencies

Run the following command to install the required Go modules:

```bash
go mod tidy
```

This will fetch the necessary dependencies (such as `gin-gonic/gin`).

### 3. Setup Environment Variables (Optional)

If you want to use environment variables (e.g., for configuring your application), create a `.env` file in the root of the project. The `config.go` will load environment variables using `github.com/joho/godotenv`.

Example `.env` file:

```
PORT=8080
```

### 4. Run the Application

You can start the application using the following command:

```bash
go run main.go
```

This will start the API server on `http://localhost:8080`.

## API Endpoints

### 1. **Create a Book** (POST `/books`)

Adds a new book to the collection.

#### Request
```bash
POST /books
Content-Type: application/json
```

#### Request Body
```json
{
  "id": "3",
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "year": "1925"
}
```

#### Response
```json
{
  "message": "Book created successfully"
}
```

### 2. **Get All Books** (GET `/books`)

Retrieves all books in the collection.

#### Request
```bash
GET /books
```

#### Response
```json
[
  {
    "id": "1",
    "title": "1984",
    "author": "George Orwell",
    "year": "1949"
  },
  {
    "id": "2",
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "year": "1960"
  }
]
```

### 3. **Get a Book by ID** (GET `/books/:id`)

Retrieves a book by its `ID`.

#### Request
```bash
GET /books/1
```

#### Response
```json
{
  "id": "1",
  "title": "1984",
  "author": "George Orwell",
  "year": "1949"
}
```

### 4. **Update a Book** (PUT `/books/:id`)

Updates an existing book by its `ID`.

#### Request
```bash
PUT /books/1
Content-Type: application/json
```

#### Request Body
```json
{
  "title": "Nineteen Eighty-Four",
  "author": "George Orwell",
  "year": "1949"
}
```

#### Response
```json
{
  "message": "Book updated successfully"
}
```

### 5. **Delete a Book** (DELETE `/books/:id`)

Removes a book from the collection by its `ID`.

#### Request
```bash
DELETE /books/1
```

#### Response
```json
{
  "message": "Book removed successfully"
}
```

---

## File Structure

```plaintext
gin-crud-api/
├── config/            # Configuration files
│   └── config.go      # Load environment variables
├── controllers/       # Route handlers
│   └── book_controller.go  # Controller logic for book routes
├── models/            # Data models
│   └── book.go        # Book struct definition
├── services/          # Business logic
│   └── book_service.go    # Logic to manage book operations
├── repositories/      # Data access layer
│   └── book_repo.go       # Methods for querying and modifying book data
├── middlewares/       # Custom middleware (e.g., auth, logging)
│   └── auth_middleware.go   # Example middleware (JWT)
├── routes/            # Route definitions
│   └── book_routes.go # Define book-related routes
├── utils/             # Utility functions (e.g., standardized response format)
│   └── response.go    # Helper functions for consistent responses
├── main.go            # Application entry point
├── go.mod             # Go modules file
└── go.sum             # Go checksum file
```

## Error Handling

The API provides consistent error messages with HTTP status codes:

- **404** - Book not found
- **400** - Invalid request (e.g., missing or incorrect parameters)
- **500** - Internal server error

Example error response:

```json
{
  "error": "Book not found"
}
```

## Testing the API

You can test the API using [Postman](https://www.postman.com/) or `curl`.

### Example `curl` commands:

1. **Create a book:**
   ```bash
   curl -X POST http://localhost:8080/books -d '{"id":"3", "title":"The Great Gatsby", "author":"F. Scott Fitzgerald", "year":"1925"}' -H "Content-Type: application/json"
   ```

2. **Get all books:**
   ```bash
   curl http://localhost:8080/books
   ```

3. **Get a book by ID:**
   ```bash
   curl http://localhost:8080/books/1
   ```

4. **Update a book:**
   ```bash
   curl -X PUT http://localhost:8080/books/1 -d '{"title":"Nineteen Eighty-Four", "author":"George Orwell", "year":"1949"}' -H "Content-Type: application/json"
   ```

5. **Delete a book:**
   ```bash
   curl -X DELETE http://localhost:8080/books/1
   ```

