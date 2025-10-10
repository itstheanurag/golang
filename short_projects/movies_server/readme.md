# Movies Server (Golang REST API)

This project is a simple RESTful API server for managing a list of movies, written in Go. It demonstrates basic CRUD (Create, Read, Update, Delete) operations using the [Gorilla Mux](https://github.com/gorilla/mux) router.

## Features

- List all movies (`GET /movies`)
- Get a movie by ID (`GET /movies/{id}`)
- Add a new movie (`POST /movies`)
- Update an existing movie (`PUT /movies/{id}`)
- Delete a movie (`DELETE /movies/{id}`)

## Movie Data Structure

Each movie has the following structure:

```json
{
  "id": "string",
  "isbn": "string",
  "title": "string",
  "director": {
    "first_name": "string",
    "last_name": "string"
  }
}
```

## Getting Started

### Prerequisites

- Go 1.18 or higher installed

### Installation

1. Clone the repository or copy the project files.
2. Install dependencies:
   ```sh
   go get github.com/gorilla/mux
   ```

### Running the Server

Navigate to the `movies_server` directory and run:

```sh
go run main.go
```

The server will start on [http://localhost:8000](http://localhost:8000).

## API Endpoints

| Method | Endpoint       | Description          |
| ------ | -------------- | -------------------- |
| GET    | `/movies`      | Get all movies       |
| GET    | `/movies/{id}` | Get a movie by ID    |
| POST   | `/movies`      | Add a new movie      |
| PUT    | `/movies/{id}` | Update a movie by ID |
| DELETE | `/movies/{id}` | Delete a movie by ID |

### Example: Add a Movie

```sh
curl -X POST http://localhost:8000/movies \
  -H "Content-Type: application/json" \
  -d '{"isbn":"123456789","title":"New Movie","director":{"first_name":"John","last_name":"Doe"}}'
```

### Example: Get All Movies

```sh
curl http://localhost:8000/movies
```

## Notes

- The server uses an in-memory slice to store movies; data will reset when the server restarts.
- The server is for demonstration and learning purposes.

## License

This project is for educational use.
