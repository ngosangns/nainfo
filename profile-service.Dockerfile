# profile-service/Dockerfile
FROM golang:1.22-alpine

WORKDIR /app

COPY profile-service/go.mod profile-service/go.sum .env ./
COPY ./shared /shared
RUN go mod download

COPY ./profile-service .

RUN go build -o profile-service profile-service

ENV MYSQL_DSN="user:password@tcp(mysql:3306)/dbname"
ENV PROFILE_SERVICE_ADDRESS=":8001"

EXPOSE 8001

CMD ["./profile-service"]