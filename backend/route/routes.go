package route

import (
	"backend/controller"
	_ "backend/model"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// route registration
	r = CustomerRoute(r)
	// route cuisine
	r = CuisineRoute(r)
	// route rating
	r = RatingRoute(r)
	// route order
	r = OrderRoute(r)
	//route restaurant
	r = RestaurantRoute(r)
	return r
}

func CustomerRoute(r *gin.Engine) *gin.Engine {
	customerRoutes := r.Group("/customer")
	customerRoutes.POST("/register", controller.Register)
	customerRoutes.GET("/login", controller.Login)
	customerRoutes.GET("/delete", controller.DeleteUser)
	customerRoutes.GET("/read/:token", controller.ReadUser)
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

func RestaurantRoute(r *gin.Engine) *gin.Engine {
	restaurantRoutes := r.Group("restaurant")
	// restaurantRoutes.POST("create",)
	// restaurantRoutes.POST("delete",)
	// restaurantRoutes.POST("update",)
	restaurantRoutes.GET("read", controller.ReadRestaurant)
	return r
}
