# Student API using Golang

This project is a simple RESTful API for managing student records. It is built using the Go programming language and utilizes SQLite for data storage. The API provides endpoints for creating, retrieving, updating, and deleting student records.

## Tech Stack

- **Go**: The primary programming language used for developing the API.
- **SQLite**: A lightweight, disk-based database used for storing student records.
- **Go-Playground Validator**: A library for validating request payloads.
- **Cleanenv**: A library for loading configuration from environment variables and configuration files.
- **Slog**: A structured logging library for Go.
- **HTTP**: The standard library's `net/http` package is used for handling HTTP requests and responses.

## Endpoints

- `POST /api/students`: Create a new student.
- `GET /api/students/{id}`: Retrieve a student by ID.
- `GET /api/students`: Retrieve a list of all students.
- `PUT /api/students/{id}`: Update a student by ID.
- `DELETE /api/students/{id}`: Delete a student by ID.

## Configuration

The application uses a configuration file to load settings such as the server address and the path to the SQLite database file. The configuration file should be specified using the `--config` flag or the `CONFIG_PATH` environment variable.

Example configuration file (`config/local.yaml`):

```yaml
env: development
storage_path: "./students.db"
http_server:
    address: ":8080"
```

## Running the Application

To run the application, use the following command:

```bash
set CGO_ENABLED=1
go run cmd/main.go --config your-config.yaml
```

## Learning Outcomes

By working on this project, you will learn:
- How to build a RESTful API using Go.
- How to use SQLite for data storage in a Go application.
- How to validate request payloads using the Go-Playground Validator library.
- How to load configuration from environment variables and configuration files using Cleanenv.
- How to handle HTTP requests and responses using the standard library's net/http package.
- How to implement structured logging using Slog.

## Project Structure

```text
student-api-golang/
├── cmd/
│   └── students-api/
│       └── main.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go
│   ├── storage/
│   │   ├── sqllite/
│   │   │   └── sqllite.go
│   │   └── storage.go
│   ├── types/
│   │   └── types.go
│   └── utils/
│       └── response/
│           └── response.go
├── .gitignore
├── go.mod
└── go.sum
```

## File Descriptions

[`main.go`](cmd/students-api/main.go)
This is the entry point of the application. It initializes the configuration, sets up the database connection, configures the HTTP router, and starts the HTTP server. It also handles graceful shutdown of the server.

[`config.go`](pkg/config/config.go)
This file contains the configuration loading logic. It uses the Cleanenv library to load configuration from environment variables and configuration files.

[`student.go`](pkg/http/handlers/student/student.go)
This file contains the HTTP handlers for managing student records. It includes handlers for creating, retrieving, updating, and deleting students.

[`sqllite.go`](pkg/storage/sqllite/sqllite.go)
This file contains the logic for interacting with the SQLite database. It includes functions for creating, retrieving, updating, and deleting student records.

[`storage.go`](pkg/storage/storage.go)
This file defines the Storage interface, which abstracts the database operations. It includes methods for creating, retrieving, updating, and deleting student records.

[`types.go`](pkg/types/types.go)
This file defines the data structures used in the application, such as the Student struct.

[`response.go`](pkg/utils/response/response.go)
This file contains utility functions for writing JSON responses and handling validation errors.

## Reference

This project was built using the tutorial by Coders Gyan. You can watch the tutorial [here](https://www.youtube.com/watch?v=OGhQhFKvMiM).