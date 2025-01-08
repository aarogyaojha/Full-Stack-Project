package main

import (
	routes "Full-Stack-Project/routes"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	routes.SetupRoutes()
	port:=os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router :=gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{"sucess":"Access granted for api 1"})
	})

	router.GET("/api-2", func(ctx * gin.Context){
		ctx.JSON(200,gin.H{"success":"Access granted for api 2"})
	})

	router.Run(":"+ port)
}