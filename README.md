### Tools

- `go install github.com/swaggo/swag/cmd/swag@latest`

### Running the Services

To run the services, you can use the following commands. Make sure your MySQL service is running and accessible.

1. Start auth-service:
   ```bash
   cd auth-service
   go run main.go
   ```

2. Start profile-service:
   ```bash
   cd profile-service
   go run main.go
   ```

3. Start the API Gateway:
   ```bash
   cd api-gateway
   go run main.go
   ```

---

### Running the Services with Docker

With the `Dockerfile` and `docker-compose.yml` files set up, you can now run all your services with a single command:

```bash
docker-compose up --build
```

This command will build the Docker images for each service and start the containers as defined in the `docker-compose.yml` file.

You can access:
- The **Auth Service** on `http://localhost:8000`
- The **Profile Service** on `http://localhost:8001`
- The **API Gateway** on `http://localhost:8080`

Ensure MySQL is up and running before other services try to connect to it, which is managed by `depends_on` in the `docker-compose.yml` file. Also, make sure to change the environment variables (`MYSQL_DSN`, `JWT_SECRET`, etc.) as per your actual setup for security and configurations.