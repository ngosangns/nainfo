This file will provide an overview of the auth service, including a brief description and instructions on how to set up and run the service.

```markdown
# Auth Service

The Auth Service handles all authentication-related functionalities such as login and registration.

## Getting Started

### Prerequisites

- Go 1.20+
- MySQL
- Environment variables set up

### Environment Variables

Create a `.env` file in the `auth-service` directory with the following contents:

```
MYSQL_DSN="user:password@tcp(localhost:3306)/dbname"
JWT_SECRET="your_secret_key"
AUTH_SERVICE_ADDRESS=":8000"
```

### Running the Service

1. Ensure you have MySQL running and accessible with the credentials provided in the `.env` file.

2. Navigate to the `auth-service` directory:

    ```sh
    cd auth-service
    ```

3. Run the service:

    ```sh
    go run cmd/auth-service/main.go
    ```

### API Endpoints

The following are the available API endpoints for the Auth Service:

- [POST /login](api.md#post-login)
- [POST /register](api.md#post-register)

## Architecture

The Auth Service is built using Domain-Driven Design (DDD) and Clean Architecture principles. For more details, see [architecture.md](architecture.md).

## Tests

To run tests:

```sh
go test ./...
```