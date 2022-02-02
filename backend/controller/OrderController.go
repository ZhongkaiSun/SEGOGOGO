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
	userName := requestOrder.UserName
	restaurantName := requestOrder.RestaurantName
	orderDate := requestOrder.OrderDate
	price := requestOrder.Price
	cuisineName := requestOrder.CuisineName

	if isUserExsit(DB, userName) || userName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "User doesn't exist, please bind a valid username"})
		return
	}

	newOrder := model.Order{
		UserName:       userName,
		RestaurantName: restaurantName,
		OrderDate:      orderDate,
		Price:          price,
		CuisineName:    cuisineName,
	}
	DB.Create(&newOrder)
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
	userName := requestOrder.UserName
	log.Println(userName)
	if isUserExsit(DB, userName) || userName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "User doesn't exist, please bind a valid username"})
		return
	}
	var newOrders []model.Order
	DB.Where("username = ?", userName).Find(&newOrders)
	c.JSON(http.StatusOK, newOrders)
}

func isUserExsit(db *gorm.DB, userName string) bool {
	var order model.Order
	db.Where("username = ?", userName).First((&order))
	return order.UserName == ""
}
