This file will provide an overview of the architecture used in the auth service.

```markdown
# Auth Service Architecture

The Auth Service is designed following principles of Domain-Driven Design (DDD) and Clean Architecture. This structure aims to maintain separation of concerns and allow the system to be extendable and maintainable.

## Layers

### 1. Domain

The domain layer contains the business logic of the application.

- **Model**: Defines the core entities of the system (e.g., User).
- **Repository**: Defines interfaces for persistence operations required by the domain.

### 2. Application

The application layer contains the use cases and services that orchestrate domain logic.

- **Service**: Implements application business logic using domain entities and interfaces.
- **Handler**: Defines HTTP handlers that interact with the outside world (e.g., controllers in MVC.

### 3. Infrastructure

The infrastructure layer contains implementation details and external interfaces.

- **Persistence**: Implements repository interfaces for actual data storage (e.g., MySQL).
- **Router**: Sets up HTTP routes and binds them to the appropriate handlers.

## Workflow

1. **HTTP Request**: An incoming HTTP request is received by the HTTP router.
2. **Handler**: The router forwards the request to the appropriate handler in the application layer.
3. **Service**: The handler calls the necessary service methods to process the request.
4. **Domain Logic**: Service methods may interact with domain entities and repositories to execute business logic.
5. **Response**: The handler returns the appropriate HTTP response to the client.

Below is a diagram illustrating the architecture:

```
      ┌───────────────────────────────────┐
      │              Client                │
      └───────────────────────────────────┘
                        │
                        ▼
      ┌───────────────────────────────────┐
      │              Router                │
      └───────────────────────────────────┘
                        │
                        ▼
      ┌───────────────────────────────────┐
      │              Handler               │
      └───────────────────────────────────┘
                        │
                        ▼
      ┌───────────────────────────────────┐
      │              Service               │
      └───────────────────────────────────┘
                        │
                        ▼
      ┌───────────────────────────────────┐
      │     Domain (Model & Repository)    │
      └───────────────────────────────────┘
```

## Dependencies

External libraries and tools used in the Auth Service:

- **gin**: Web framework for creating HTTP endpoints.
- **bcrypt**: Library for hashing and verifying passwords.
- **jwt-go**: Library for generating and parsing JWT tokens.
- **mysql**: MySQL driver for connecting to the database.