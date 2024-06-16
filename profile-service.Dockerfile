# profile-service/Dockerfile
FROM golang:1.22-alpine

COPY profile-service/go.mod profile-service/go.sum ./
COPY ./shared /shared

RUN go mod download
RUN go install github.com/air-verse/air@latest
RUN rm -rf /shared

EXPOSE 8001
