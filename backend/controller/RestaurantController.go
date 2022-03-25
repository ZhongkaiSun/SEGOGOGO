package controller

import (
	"backend/common"
	"backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//无商家端口，增删改暂无

//可以输入餐厅名字查找，若无输入则展示全部餐厅
func ReadRestaurant(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Origin", "*")
	DB := common.GetDB()
	var requestRestaurant model.Restaurant
	err := c.ShouldBindQuery(&requestRestaurant)
	if err != nil {
		c.JSON(422, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}

	name := requestRestaurant.Name
	typeOfMeal := requestRestaurant.TypeofMeal
	//name? resId?? unique
	if name != "" {
		if !isRestaurantNameExsit(DB, name) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Restaurant doesn't exist, please search a valid restaurant"})
			return
		} else {
			var newRestaurant model.Restaurant
			DB.Where("name = ?", name).First(&newRestaurant)
			c.JSON(http.StatusOK, gin.H{"code": 200, "data": newRestaurant, "msg": "Successfully"})
		}
	} else if typeOfMeal != "" { //get all restaurants
		var newRestaurantList []model.Restaurant
		DB.Where("typeof_meal = ?", typeOfMeal).Find(&newRestaurantList)
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": newRestaurantList, "msg": "Successfully"})
	} else {
		fmt.Println("all empty")
	}
}

func isRestaurantNameExsit(db *gorm.DB, name string) bool {
	var restaurant model.Restaurant
	db.Where("name = ?", name).First(&restaurant)
	return restaurant.Name != ""
}
