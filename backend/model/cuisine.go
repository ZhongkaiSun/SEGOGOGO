package model

import (
	"backend/common"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// CREATE TABLE cuisines (
// 	name VARCHAR ( 50 ),
// 	restaurant_id VARCHAR ( 50 ),
// 	price FLOAT,
// 	calories INT,
// 	PRIMARY KEY (name, restaurant_id)
// );

type Cuisine struct {
	Name         string  `gorm:"type:varchar(50);primary_key" json:"name"`
	RestaurantId string  `gorm:"type:varchar(50);primary_key" json:"restaurantId"`
	Price        float64 `gorm:"type:float" json:"price"`
	Calories     int     `gorm:"type:int" json:"calories"`
}

func CreateCuisine(c *gin.Context) {
	DB := common.GetDB()
	var requestCuisine Cuisine
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

	newCuisine := Cuisine{
		Name:         name,
		RestaurantId: restaurantId,
		Price:        price,
		Calories:     calories,
	}
	DB.Create(&newCuisine)

}

func isCuisineExsit(db *gorm.DB, name string) bool {
	var cuisine Cuisine
	db.Where("name = ?", name).First(&cuisine)
	return cuisine.Name == ""
}

func isRestaurantExsit(db *gorm.DB, restaurantId string) bool {
	var cuisine Cuisine
	db.Where("restaurant_id = ?", cuisine).First(&cuisine)
	return cuisine.RestaurantId == ""
}
