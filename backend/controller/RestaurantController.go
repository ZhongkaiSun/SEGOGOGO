package controller

import (
	"backend/common"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//无商家端口，增删改暂无

//可以输入餐厅名字查找，若无输入则展示全部餐厅
func ReadRestaurant(c *gin.Context) {
	DB := common.GetDB()
	var requestRestaurant model.Restaurant
	err := c.ShouldBindQuery(&requestRestaurant)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(422, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}
	name := requestRestaurant.Name

	//name? resId?? unique
	if name != "" {
		if !isRestaurantNameExsit(DB, name) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Restaurant doesn't exist, please search a valid restaurant"})
			return
		} else {
			var newRestaurant model.Restaurant
			DB.Where("name = ?", name).First(&newRestaurant)
			c.Header("Access-Control-Allow-Origin", "*")
			c.JSON(http.StatusOK, gin.H{"code": 200, "data": newRestaurant, "msg": "Successfully"})
		}
	} else { //get all restaurants
		var newRestaurantList []model.Restaurant
		DB.Find(&newRestaurantList)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": newRestaurantList, "msg": "Successfully"})
	}
}

func isRestaurantNameExsit(db *gorm.DB, name string) bool {
	var restaurant model.Restaurant
	db.Where("name = ?", name).First(&restaurant)
	return restaurant.Name != ""
}
