package controller

import (
	"backend/common"
	"backend/model"
	_ "fmt"
	_ "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Create(c *gin.Context) {
	DB := common.GetDB()
	var requestCuisine model.Cuisine
	err := c.ShouldBindJSON(&requestCuisine)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}
	name := requestCuisine.Name
	restaurantId := requestCuisine.RestaurantId
	price := requestCuisine.Price
	calories := requestCuisine.Calories

	if isCuisineExsit(DB, name) || name == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid name"})
		return
	}
	//must be a valid restaurant: length/ not exist
	if isRestaurantExsit(DB, restaurantId) || restaurantId == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please bind a valid restaurantId"})
		return
	}

	if price < 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid price"})
		return
	}

	if calories < 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid calories value "})
		return
	}

	newCuisine := model.Cuisine{
		Name:         name,
		RestaurantId: restaurantId,
		Price:        price,
		Calories:     calories,
	}
	DB.Create(&newCuisine)

}

func isCuisineExsit(db *gorm.DB, name string) bool {
	var cuisine model.Cuisine
	db.Where("name = ?", name).First(&cuisine)
	return cuisine.Name == ""
}

func isRestaurantExsit(db *gorm.DB, restaurantId string) bool {
	var cuisine model.Cuisine
	db.Where("restaurant_id = ?", cuisine).First(&cuisine)
	return cuisine.RestaurantId == ""
}

// func Delete
