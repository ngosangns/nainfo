package config

import (
	"os"
)

func MySQLDSN() string {
	return os.Getenv("MYSQL_DSN")
}

func JWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func ProfileServiceAddress() string {
	return os.Getenv("PROFILE_SERVICE_ADDRESS")
}

func ProfileServiceGRPCAddress() string {
	return os.Getenv("PROFILE_SERVICE_GRPC_ADDRESS")
}

func AuthServiceAddress() string {
	return os.Getenv("AUTH_SERVICE_ADDRESS")
}

func AuthServiceGRPCAddress() string {
	return os.Getenv("AUTH_SERVICE_GRPC_ADDRESS")
}

func APIGatewayAddress() string {
	return os.Getenv("API_GATEWAY_ADDRESS")
}

func APIGatewayGRPCAddress() string {
	return os.Getenv("API_GATEWAY_GRPC_ADDRESS")
}

func AdminerAddress() string {
	return os.Getenv("ADMINER_ADDRESS")
}

func AdminerPort() string {
	return os.Getenv("ADMINER_PORT")
}

func AdminerUsername() string {
	return os.Getenv("ADMINER_USERNAME")
}

func AdminerPassword() string {
	return os.Getenv("ADMINER_PASSWORD")
}

func AdminerDatabase() string {
	return os.Getenv("ADMINER_DATABASE")
}

func AdminerHost() string {
	return os.Getenv("ADMINER_HOST")
}

func MinIOHost() string {
	return os.Getenv("MINIO_HOST")
}

func MinIOAccessKey() string {
	return os.Getenv("MINIO_ACCESS_KEY")
}

func MinIOSecretKey() string {
	return os.Getenv("MINIO_SECRET_KEY")
}

func MinIOBucketName() string {
	return os.Getenv("MINIO_BUCKET_NAME")
}
