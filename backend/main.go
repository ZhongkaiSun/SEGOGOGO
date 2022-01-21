package main

import (
	"backend/common"
	"backend/controller"
	_ "backend/model"
	_ "fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()

	r := gin.Default()
	// route registration
	customerRoutes := r.Group("/customer")
	customerRoutes.POST("/register", controller.Register)
	customerRoutes.GET("/login", controller.Login)
	// customerRoutes.GET("/delete", controller.Delete)
	// route registration
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

	// var res model.Restaurant
	// DB.Table("restaurants").Where("ID=?", "F0001").Scan(&res)
	// fmt.Println(res)

	// var rat model.Rating
	// DB.Table("ratings").Where("ID=?", "R0001").Scan(&rat)
	// fmt.Println(rat)

	// var cui model.Cuisine
	// DB.Table("cuisines").Where("Name=?", "8PC Nuggets A La Carte").Scan(&cui)
	// fmt.Println(cui)

	// var cus model.Customer
	// DB.Table("customers").Where("Username=?", "ZhongkaiSun").Scan(&cus)
	// fmt.Println(cus)
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
