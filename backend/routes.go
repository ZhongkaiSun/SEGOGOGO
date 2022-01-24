package main

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) *gin.Engine {
	customerRoutes := r.Group("/customer")
	customerRoutes.POST("/register", controller.Register)
	customerRoutes.GET("/login", controller.Login)
	customerRoutes.GET("/delete", controller.Delete)
	return r
}
