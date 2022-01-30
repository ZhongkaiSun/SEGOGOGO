package main

import (
	"backend/common"
	_ "backend/controller"
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
	r = RegisterRoute(r)
	// route model
	r = ModelRoute(r)
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
