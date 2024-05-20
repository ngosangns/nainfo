This file will document the API endpoints provided by the auth service, including request and response formats.

```markdown
# Auth Service API

## POST /login

Endpoint for logging in a user.

### Request

```http
POST /login HTTP/1.1
Content-Type: application/json

{
    "username": "exampleuser",
    "password": "examplepassword"
}
```

### Response

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
    "token": "jwt-token-example"
}
```

### Error Responses

```http
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
    "error": "Invalid request payload"
}
```

```http
HTTP/1.1 401 Unauthorized
Content-Type: application/json

{
    "error": "Invalid username or password"
}
```

## POST /register

Endpoint for registering a new user.

### Request

```http
POST /register HTTP/1.1
Content-Type: application/json

{
    "username": "newuser",
    "password": "newpassword",
    "email": "email@example.com"
}
```

### Response

```http
HTTP/1.1 201 Created
Content-Type: application/json

{
    "message": "registration successful"
}
```

### Error Responses

```http
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
    "error": "Invalid request payload"
}
```

```http
HTTP/1.1 500 Internal Server Error
Content-Type: application/json

{
    "error": "Could not register user"
}
```