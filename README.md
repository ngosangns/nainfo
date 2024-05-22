# Nainfo

This project implements a microservices-based system for managing user profiles and authentication.

## Overview

Nainfo consists of the following services:

* **Auth Service:** Handles user authentication and registration.
* **Profile Service:** Manages user profile information.
* **API Gateway:** Acts as a centralized entry point for clients to interact with the system.

## Architecture

The system uses Docker Compose to orchestrate the different services and utilizes a MySQL database for data persistence.

## Setup

1. Install Docker and Docker Compose.
2. Clone the repository.
3. Create a `.env` file in the root directory and set the following environment variables:

```
MYSQL_DSN=user:password@host:port/dbname
JWT_SECRET=your_secret_key
AUTH_SERVICE_ADDRESS=:8000
PROFILE_SERVICE_GRPC_ADDRESS=:50051
PROFILE_SERVICE_ADDRESS=:8001
```

4. Run `docker-compose up -d` to start the services.

## Usage

* Access the API Gateway at `http://localhost:8080`.
* See the `api.md` file for API documentation.

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
