# auth-service/Dockerfile
FROM golang:1.22-alpine

COPY auth-service/go.mod auth-service/go.sum ./
COPY ./shared /shared

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN rm -rf /shared

EXPOSE 8000