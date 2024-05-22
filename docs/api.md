
# Nainfo API Documentation

This document outlines the API endpoints for the Nainfo system. The API provides endpoints for user authentication, profile management, and other functionalities.

## Authentication

### Endpoint: `/auth/login`

**Method:** POST

**Request Body:**

```json
{
  "username": "username",
  "password": "password"
}
```

**Response Body:**

```json
{
  "token": "your_jwt_token"
}
```

**Example:**

```json
{
  "username": "testuser",
  "password": "testpassword"
}
```

### Endpoint: `/auth/register`

**Method:** POST

**Request Body:**

```json
{
  "username": "username",
  "password": "password"
}
```

**Response Body:**

**Status Code:** 204 (No Content)

**Example:**

```json
{
  "username": "newuser",
  "password": "newpassword"
}
```

## Profile Service

### Endpoint: `/profile/me`

**Method:** GET

**Request Headers:**

* `Authorization`: `Bearer your_jwt_token`

**Response Body:**

```json
{
  "username": "username",
  "email": "email@example.com",
  "name": "John Doe",
  "description": "A brief description",
  "address": "123 Main Street",
  "facebook": "facebook.com/john.doe",
  "linkedin": "linkedin.com/in/john-doe",
  "github": "github.com/johndoe"
}
```

**Example:**

```
curl -H "Authorization: Bearer your_jwt_token" http://localhost:8001/profile/me
```

### Endpoint: `/profile/me`

**Method:** PUT

**Request Headers:**

* `Authorization`: `Bearer your_jwt_token`

**Request Body:**

```json
{
  "email": "newemail@example.com",
  "name": "John Doe",
  "description": "A brief description",
  "address": "123 Main Street",
  "facebook": "facebook.com/john.doe",
  "linkedin": "linkedin.com/in/john-doe",
  "github": "github.com/johndoe"
}
```

**Response Body:**

**Status Code:** 204 (No Content)

**Example:**

```
curl -H "Authorization: Bearer your_jwt_token" -X PUT -d '{"email": "newemail@example.com", "name": "John Doe", "description": "A brief description", "address": "123 Main Street", "facebook": "facebook.com/john.doe", "linkedin": "linkedin.com/in/john-doe", "github": "github.com/johndoe"}' http://localhost:8001/profile/me
```


## API Gateway

The API Gateway routes requests to the appropriate services based on the path and authentication. It handles authorization and provides a centralized entry point for clients to interact with the Nainfo system.

### Endpoint: `/auth/*`

**Method:** ANY

**Request Headers:**

* `Authorization`: `Bearer your_jwt_token` (optional for login/register)

**Proxies to:** Auth Service

### Endpoint: `/profile/*`

**Method:** ANY

**Request Headers:**

* `Authorization`: `Bearer your_jwt_token` (required)

**Proxies to:** Profile Service


## Error Handling

The API uses standard HTTP status codes for error handling. 

* **400 Bad Request:** Invalid request parameters.
* **401 Unauthorized:** Missing or invalid authentication token.
* **403 Forbidden:** Insufficient permissions.
* **404 Not Found:** Resource not found.
* **500 Internal Server Error:** An internal error occurred.

## Versioning

The API uses semantic versioning. The current version is **v1.0**.

## Future Development

* Implement more API endpoints for different functionalities.
* Add more detailed documentation for each endpoint.
* Improve error handling and logging.
* Implement rate limiting and security measures.