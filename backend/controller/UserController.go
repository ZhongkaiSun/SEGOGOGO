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

func Delete(c *gin.Context) {
	DB := common.GetDB()
	var requestCustomer model.Customer
	err := c.ShouldBindJSON(&requestCustomer)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}
	username := requestCustomer.Username
	password := requestCustomer.Password

	targetCustomer := model.Customer{}
	DB.Where("username = ?", username).First(&targetCustomer)
	if targetCustomer.Password != password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Does not exist or the password is wrong"})
		return
	} else {
		DB.Delete(&targetCustomer)
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "Successfully deleted"})
		return
	}
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	var requestCustomer model.Customer
	err := c.ShouldBindJSON(&requestCustomer)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}

	username := requestCustomer.Username
	password := requestCustomer.Password

	targetCustomer := model.Customer{}
	DB.Where("username = ?", username).First(&targetCustomer)
	if targetCustomer.Password != password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "The password is wrong"})
		return
	}

	token, err := common.ReleaseToken(targetCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error : %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "Successfully log in"})
}

// CREATE TABLE customers (
// 	username VARCHAR ( 50 ) PRIMARY KEY,
// 	password VARCHAR ( 50 ) NOT NULL,
// 	address_line1 VARCHAR ( 50 ) NOT NULL,
// 	address_line2 VARCHAR ( 50 ),
// 	phone VARCHAR ( 50 ) );
func Register(c *gin.Context) {
	DB := common.GetDB()
	var requestCustomer model.Customer
	err := c.ShouldBindJSON(&requestCustomer)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}

	username := requestCustomer.Username
	password := requestCustomer.Password
	phone := requestCustomer.Phone

	if isCustomerExist(DB, username) || username == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Create a new username"})
		return
	}

	if len(phone) != 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "len of phone must be 10"})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "len of pwd must be longer than 6"})
		return
	}

	newCustomer := model.Customer{
		Username:     username,
		Password:     password,
		AddressLine1: requestCustomer.AddressLine1,
		AddressLine2: requestCustomer.AddressLine2,
		Phone:        phone,
		Email:        requestCustomer.Email,
	}
	DB.Create(&newCustomer)

	token, err := common.ReleaseToken(newCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "system error"})
		log.Printf("token generate error : %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "Successfully"})
}

func isCustomerExist(db *gorm.DB, username string) bool {
	var customer model.Customer
	db.Where("username = ?", username).First(&customer)
	return customer.Username != ""
}
