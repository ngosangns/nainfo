package router

import (
	"auth-service/application/handler"
	"auth-service/infrastructure/persistence"
	"database/sql"
	"shared/config"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	db, err := sql.Open("mysql", config.MySQLDSN())
	if err != nil {
		panic(err)
	}
	userRepository := persistence.NewMySQLUserRepository(db)
	authHandler := handler.NewAuthHandler(userRepository)

	r.POST("/auth/login", authHandler.Login)
	r.POST("/auth/register", authHandler.Register)

	return r
}
