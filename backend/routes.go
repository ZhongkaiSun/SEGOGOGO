package main

import (
	"backend/controller"
	"backend/model"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) *gin.Engine {
	customerRoutes := r.Group("/customer")
	customerRoutes.POST("/register", controller.Register)
	customerRoutes.GET("/login", controller.Login)
	customerRoutes.GET("/delete", controller.Delete)
	return r
}

func ModelRoute(r *gin.Engine) *gin.Engine {
	modelRoutes := r.Group("/model")
	modelRoutes.POST("/createCuisine", model.CreateCuisine)
	//modelRoutes.POST("/deleteCuisine", model.DeleteCuisine)
	//modelRoutes.POST("/updateCuisine", model.UpdateCuisine)
	//modelRoutes.GET("/readCuisine", model.ReadCuisine)

	return r
}
