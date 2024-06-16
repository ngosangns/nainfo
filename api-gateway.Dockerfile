# api-gateway/Dockerfile
FROM golang:1.22-alpine

COPY api-gateway/go.mod api-gateway/go.sum ./
COPY ./shared /shared

RUN go mod download
RUN go install github.com/air-verse/air@latest
RUN rm -rf /shared

EXPOSE 8080
