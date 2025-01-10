package routes

import (
	controllers "github.com/aarogyaojha/Full-Stack-Project/controllers"
	"github.com/aarogyaojha/Full-Stack-Project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
