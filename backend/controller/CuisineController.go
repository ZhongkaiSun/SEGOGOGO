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
	restaurantName := requestCuisine.RestaurantName
	price := requestCuisine.Price
	calories := requestCuisine.Calories
	if !isRestaurantExsit(DB, restaurantName) || restaurantName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Restaurant doesn't exist, please bind a valid restaurantName"})
		return
	}

	if isCuisineExsit(DB, name, restaurantName) || name == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Cuisine already exists"})
		return
	}

	// if price < 0 {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid price"})
	// 	return
	// }

	// if calories < 0 {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid calories value "})
	// 	return
	// }

	newCuisine := model.Cuisine{
		Name:           name,
		RestaurantName: restaurantName,
		Price:          price,
		Calories:       calories,
	}
	DB.Create(&newCuisine)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "Successfully"})
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
	restaurantName := requestCuisine.RestaurantName
	price := requestCuisine.Price
	calories := requestCuisine.Calories

	if !isRestaurantExsit(DB, restaurantName) || restaurantName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please bind a valid restaurantName"})
		return
	}

	if !isCuisineExsit(DB, name, restaurantName) || name == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid name"})
		return
	}

	// if price < 0 {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid price"})
	// 	return
	// }

	// if calories < 0 {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please create a valid calories value "})
	// 	return
	// }

	deleteCuisine := model.Cuisine{
		Name:           name,
		RestaurantName: restaurantName,
		Price:          price,
		Calories:       calories,
	}
	DB.Delete(&deleteCuisine)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "Successfully"})
}

//以restaurantName获取
func ReadCuisine(c *gin.Context) {
	DB := common.GetDB()
	var requestCuisine model.Cuisine
	err := c.ShouldBindQuery(&requestCuisine)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}

	restaurantName := requestCuisine.RestaurantName

	if !isRestaurantExsit(DB, restaurantName) || restaurantName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Please bind a valid restaurantName"})
		return
	}

	var newCuisines []model.Cuisine
	DB.Where("restaurant_name = ?", restaurantName).Find(&newCuisines)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": newCuisines, "msg": "Successfully"})
}

//以什么去查/改？
// func UpdateCuisine(c *gin.Context) {

// }

func isCuisineExsit(db *gorm.DB, name string, restaurantName string) bool {
	var cuisine model.Cuisine
	db.Where("name = ? AND restaurant_name = ?", name, restaurantName).First(&cuisine)
	return cuisine.Name != ""
}

func isRestaurantExsit(db *gorm.DB, restaurantName string) bool {
	var cuisine model.Cuisine
	db.Where("restaurant_name = ?", restaurantName).First(&cuisine)
	return cuisine.RestaurantName != ""
}
