package main

import (
	"backend/controller"
	_ "backend/model"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) *gin.Engine {
	customerRoutes := r.Group("/customer")
	customerRoutes.POST("/register", controller.Register)
	customerRoutes.POST("/login", controller.Login)
	customerRoutes.GET("/delete", controller.Delete)
	return r
}

func RatingRoute(r *gin.Engine) *gin.Engine {
	customerRoutes := r.Group("/rating")
	customerRoutes.POST("/create", controller.CreateRating)
	customerRoutes.GET("/get", controller.GetRating)
	return r
}

func CuisineRoute(r *gin.Engine) *gin.Engine {
	cuisineRoutes := r.Group("/cuisine")
	cuisineRoutes.POST("/create", controller.CreateCuisine)
	cuisineRoutes.POST("/delete", controller.DeleteCuisine)
	// cuisineRoutes.POST("/update", controller.Update)
	cuisineRoutes.GET("/read", controller.ReadCuisine)
	return r
}

func OrderRoute(r *gin.Engine) *gin.Engine {
	orderRoutes := r.Group("/order")
	orderRoutes.POST("create", controller.CreateOrder)
	orderRoutes.GET("read", controller.ReadOrder)
	return r
}
