package main

import (
	"github.com/Salavei/golang_gin/controllers"
	"github.com/Salavei/golang_gin/initializers"
	"github.com/Salavei/golang_gin/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()

	r.POST("/signup/", controllers.SignUp)
	r.POST("/login/", controllers.Login)
	r.GET("/validate/", middleware.RequireAuth, controllers.Validate)

	r.Run(":8000")
}
