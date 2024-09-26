# Task Manager API

A simple and efficient RESTful(Representational State Transfer) API built with Go for managing a collection of tasks. This API supports CRUD(Create, Read, Update, Delete) operations, allowing users to create, read, update, and delete tasks seamlessly.

## Features

- **Create Tasks**: Add new tasks with title, description, and completion status.
- **Read Tasks**: Retrieve all tasks or a specific task by ID.
- **Update Tasks**: Modify the details of existing tasks.
- **Delete Tasks**: Remove tasks from the collection.

## Technologies Used

- **Go**: The programming language used to build the API.
- **net/http**: Standard library for handling HTTP requests and responses.

## Getting Started

### Prerequisites

- Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ClemNTTS/TaskManagerAPI
   cd TaskManagerAPI
   ```

2. Run the server:

   ```bash
   go run main.go 3000
   ```

3. The API will start on `http://localhost:3000`. On 8080 port if no arguments.

## API Endpoints

### Create a Task

```http
POST /tasks
```

**Request Body**:

```json
{
  "id": "3",
  "title": "Learn Go",
  "content": "Make a cool API",
  "complete": false
}
```

### Get All Tasks

```http
GET /tasks
```

### Get a Single Task

```http
GET /tasks/{id}
```

### Update a Task

```http
PUT /tasks/{id}
```

**Request Body**:

```json
{
  "id": "1",
  "title": "New Task",
  "content": "Content updated",
  "complet": true
}
```

### Delete a Task

```http
DELETE /tasks/{id}
```

## Testing the API

You can use tools like **Postman** or **cURL** to interact with the API. Here are some example cURL commands:

- **Create a Task**:

  ```bash
  curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"id":"1","title":"Study Go","content":"Make a cool API","complete":false}'
  ```

- **Get All Tasks**:

  ```bash
  curl http://localhost:8080/tasks
  ```

- **Get a Task by ID**:

  ```bash
  curl http://localhost:8080/tasks/1
  ```

- **Update a Task**:

  ```bash
  curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title":"Learn C++","content":"Make a cool Game","complete":true}'
  ```

- **Delete a Task**:
  ```bash
  curl -X DELETE http://localhost:8080/tasks/1
  ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
