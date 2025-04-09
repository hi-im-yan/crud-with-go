# Learning Go Project

This is a beginner-friendly project to learn Go, the Chi router, PostgreSQL, and database migrations. The project focuses on building a simple web API with user management features, including inserting, updating, and retrieving users from a PostgreSQL database.

## Features
- RESTful API using Go and the Chi router (`github.com/go-chi/chi`).
- PostgreSQL as the database.
- Handling database operations using `pgx` (`github.com/jackc/pgx/v5`).
- Environment variable configuration using `godotenv`.
- Database migrations with `golang-migrate`.
- Error handling for database constraints (e.g., unique email constraint).

## Setup
### Prerequisites
- Go installed (`>=1.20` recommended)
- Docker & Docker Compose
- PostgreSQL client (optional for manual database access)

### Clone the Repository
```sh
git clone https://github.com/hi-im-yan/crud-with-go.git
cd crud-with-go
```

### Create a `.env` File
Create a `.env` file in the project root with:
```
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=crud
DB_HOST=localhost
DB_PORT=5432
```

### Start PostgreSQL with Docker
```sh
docker-compose up -d
```
This runs a PostgreSQL container accessible at `localhost:5432`.

### Database Migrations
```sh
migrations will automatically run when the server starts
```

### Start the API
```sh
go run main.go
```

## API Endpoints
Once the server is up, visit [http://localhost:8080/swagger/](http://localhost:8080/swagger/)

### Error Handling
- `400 Bad Request`: Invalid request format.
- `404 Not Found`: User not found.
- `409 Conflict`: Email already in use.
- `500 Internal Server Error`: Unexpected errors.

## Learning Goals
This project helps understand:
- Routing in Go with `chi`.
- Connecting to PostgreSQL with `pgx`.
- Handling database migrations.
- Writing structured error responses.

Happy coding! 🚀

