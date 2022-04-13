package controller

import (
	"backend/common"
	"backend/model"
	"fmt"
	_ "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateUpdatePayment(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	DB := common.GetDB()
	var requestPayment model.Payment
	err := c.ShouldBindJSON(&requestPayment)
	if err != nil {
		c.JSON(422, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}
	username := requestPayment.Username
	cardHolder := requestPayment.CardHolder
	cardNumber := requestPayment.CardNumber
	expDate := requestPayment.ExpDate
	securityCode := requestPayment.SecurityCode
	addressLine1 := requestPayment.AddressLine1
	addressLine2 := requestPayment.AddressLine2
	city := requestPayment.City
	state := requestPayment.State
	zipcode := requestPayment.Zipcode
	fmt.Println("username is: " + username)
	if !isCustomerExist(DB, username) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "User doesn't exist, please bind a valid username"})
		return
	}
	if isPaymentExsit(DB, username) {
		updatePayment := model.Payment{}
		DB.Model(&updatePayment).Where("username = ?", username).Updates(map[string]interface{}{"card_holder": cardHolder, "card_number": cardNumber, "exp_date": expDate, "security_code": securityCode, "address_line1": addressLine1, "address_line2": addressLine2, "city": city, "state": state, "zipcode": zipcode})
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 200, "data": nil, "msg": "The payment updates successfully"})
		return
	}

	newPayment := model.Payment{
		Username:     username,
		CardHolder:   cardHolder,
		CardNumber:   cardNumber,
		ExpDate:      expDate,
		SecurityCode: securityCode,
		AddressLine1: addressLine1,
		AddressLine2: addressLine2,
		City:         city,
		State:        state,
		Zipcode:      zipcode,
	}
	DB.Create(&newPayment)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "Successfully"})
}

func ReadPayment(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Origin", "*")
	DB := common.GetDB()
	var requestPayment model.Payment
	err := c.ShouldBindQuery(&requestPayment)
	log.Println("username input" + requestPayment.Username)
	if err != nil {
		c.JSON(422, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}
	username := requestPayment.Username
	log.Println(username)
	if !isCustomerExist(DB, username) || username == "" {
		c.JSON(422, gin.H{"code": 422, "data": nil, "msg": "User doesn't exist, please bind a valid username"})
		return
	}
	if !isPaymentExsit(DB, username) {
		c.JSON(422, gin.H{"code": 422, "data": nil, "msg": "The payment doesn't exist!"})
		return
	}

	var newPayment model.Payment
	DB.Where("username = ?", username).First(&newPayment) //有多种支付方式时？
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": newPayment, "msg": "Successfully"})
}

func isPaymentExsit(db *gorm.DB, username string) bool {
	var newPayment model.Payment
	db.Where("username = ?", username).First(&newPayment)
	return newPayment.Username != ""
}
