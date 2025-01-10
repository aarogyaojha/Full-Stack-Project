package routes

import (
    controllers "github.com/aarogyaojha/Full-Stack-Project/controllers"
    "github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.POST("users/signup", controllers.SignUp)
    incomingRoutes.POST("users/login", controllers.LogIn)
}