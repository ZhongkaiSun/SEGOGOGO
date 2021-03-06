package main

import (
	"backend/common"
	"backend/config"
	_ "backend/controller"
	_ "backend/model"
	"backend/route"
	_ "fmt"

	"github.com/spf13/viper"
)

func main() {
	config.InitConfig("main")
	common.InitDatabase()
	DB := common.GetDB()
	defer DB.Close()

	r := route.SetupRouter()

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
