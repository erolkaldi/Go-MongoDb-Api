package routes

import (
	controller "mongo-api/controllers"
	middleware "mongo-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("users/getusers", controller.GetUsers())
	incomingRoutes.GET("users/getuser/:user_id", controller.GetUser())
}
