package controller

import (
	"backend/common"
	"backend/model"
	_ "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateOrder(c *gin.Context) {
	DB := common.GetDB()
	var requestOrder model.Order
	err := c.ShouldBindJSON(&requestOrder)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}
	username := requestOrder.Username
	restaurantName := requestOrder.RestaurantName
	orderDate := requestOrder.OrderDate
	price := requestOrder.Price
	cuisineName := requestOrder.CuisineName

	if isUserExsit(DB, username) || username == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "User doesn't exist, please bind a valid username"})
		return
	}

	newOrder := model.Order{
		Username:       username,
		RestaurantName: restaurantName,
		OrderDate:      orderDate,
		Price:          price,
		CuisineName:    cuisineName,
	}
	DB.Create(&newOrder)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "Successfully"})
}

func ReadOrder(c *gin.Context) {
	DB := common.GetDB()
	var requestOrder model.Order
	err := c.ShouldBindJSON(&requestOrder)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}
	username := requestOrder.Username
	log.Println(username)
	if isUserExsit(DB, username) || username == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "User doesn't exist, please bind a valid username"})
		return
	}
	newOrders := []model.Order{}
	DB.Where("username = ?", username).Find(&newOrders)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": newOrders, "msg": "Successfully"})
}

func isUserExsit(db *gorm.DB, username string) bool {
	var order model.Order
	db.Where("username = ?", username).First((&order))
	return order.Username == ""
}
