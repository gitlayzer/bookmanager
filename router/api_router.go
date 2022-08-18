package router

import (
	"bookmanager/controller"
	"bookmanager/middleware"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)

	v1 := r.Group("/api/v1")
	v1.POST("/book", middleware.AuthMiddleware(), controller.CreateBookHandler)
	v1.GET("/book", controller.GetBookListHandler)
	v1.GET("/book/:id", controller.GetBookHandler)
	v1.PUT("/book", controller.UpdateBookHandler)
	v1.DELETE("/book", controller.DeleteBookHandler)
}
