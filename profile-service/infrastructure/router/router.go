package router

import (
	"database/sql"
	"profile-service/application/handler"
	"profile-service/infrastructure/persistence"
	"shared/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	db, _ := sql.Open("mysql", config.MySQLDSN())
	profileRepository := persistence.NewMySQLProfileRepository(db)
	profileHandler := handler.NewProfileHandler(profileRepository)

	r.PUT("/profile", profileHandler.UpdateProfile)
	r.GET("/profile/:username", profileHandler.GetProfile)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
