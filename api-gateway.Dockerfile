# api-gateway/Dockerfile
FROM golang:1.22-alpine

WORKDIR /app

COPY api-gateway/go.mod api-gateway/go.sum .env ./
COPY ./shared /shared
RUN go mod download

COPY ./api-gateway .

RUN go build -o api-gateway

ENV AUTH_SERVICE_ADDRESS="http://auth-service:8000"
ENV PROFILE_SERVICE_ADDRESS="http://profile-service:8001"

EXPOSE 8080

CMD ["./api-gateway"]