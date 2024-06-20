package router

import (
	"database/sql"
	"profile-service/application/handler"
	"profile-service/domain/repository"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB, profileRepository repository.ProfileRepository) *gin.Engine {
	r := gin.Default()

	profileHandler := handler.NewProfileHandler(profileRepository)

	r.PUT("/profile/me", profileHandler.UpdateProfile)
	r.GET("/profile/me", profileHandler.GetProfile)
	r.POST("/profile/avatar/upload", profileHandler.UploadAvatar)

	return r
}
