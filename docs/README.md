### Tools

- protoc-gen-go-grpc:

```sh
go install google.golang.org/grpc/cmd/protoc-gen-go_grpc@latest
# If protoc-gen-go_grpc not found in PATH
# Linux
ln -s ~/.go/bin/protoc-gen-go-grpc ~/.go/bin/protoc-gen-go_grpc
# Windows
mklink %userprofile%\go\bin\protoc-gen-go_grpc.exe %userprofile%\go\bin\protoc-gen-go-grpc.exe
```

### Running the Services with Docker

With the `Dockerfile` and `docker-compose.yml` files set up, you can now run all your services with a single command:

```bash
docker-compose up --build
```

This command will build the Docker images for each service and start the containers as defined in the `docker-compose.yml` file.
