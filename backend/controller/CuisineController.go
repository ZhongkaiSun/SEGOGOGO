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

func CreateCuisine(c *gin.Context) {
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

	if isCuisineExsit(DB, name, restaurantId) || name == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid name"})
		return
	}

	if isRestaurantExsit(DB, restaurantId) || restaurantId == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Restaurant doesn't exist, please bind a valid restaurantId"})
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

func DeleteCuisine(c *gin.Context) {
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

	if isCuisineExsit(DB, name, restaurantId) || name == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid name"})
		return
	}

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
	DB.Delete(&newCuisine)
}

//以restaurantId获取
func ReadCuisine(c *gin.Context) {
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

	restaurantId := requestCuisine.RestaurantId

	if isRestaurantExsit(DB, restaurantId) || restaurantId == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please bind a valid restaurantId"})
		return
	}

	var newCuisines []model.Cuisine
	DB.Where("restaurant_id = ?", restaurantId).Find(&newCuisines)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": newCuisines, "msg": "Successfully"})
}

//以什么去查/改？
// func UpdateCuisine(c *gin.Context) {

// }

func isCuisineExsit(db *gorm.DB, name string, restaurantId string) bool {
	var cuisine model.Cuisine
	db.Where("name = ? AND restaurant_id = ?", name, restaurantId).First(&cuisine)
	return cuisine.Name == ""
}

func isRestaurantExsit(db *gorm.DB, restaurantId string) bool {
	var cuisine model.Cuisine
	db.Where("restaurant_id = ?", restaurantId).First(&cuisine)
	return cuisine.RestaurantId == ""
}
