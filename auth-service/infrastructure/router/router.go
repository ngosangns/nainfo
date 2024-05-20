package router

import (
	"auth-service/application/handler"
	"auth-service/infrastructure/persistence"
	"database/sql"
	"shared/config"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	db, _ := sql.Open("mysql", config.MySQLDSN())
	userRepository := persistence.NewMySQLUserRepository(db)
	authHandler := handler.NewAuthHandler(userRepository)

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
