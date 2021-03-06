package controller

import (
	"backend/common"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// type Rating struct {
// 	ID           string `gorm:"type:varchar(50);primary_key" json:"id"`
// 	Username     string `gorm:"type:varchar(50)" json:"username"`
// 	RestaurantName string `gorm:"type:varchar(50)" json:"restaurantName"`
// 	Star         int    `gorm:"type:int" json:"star"`
// 	Comment      string `gorm:"type:varchar(255)" json:"comment"`
// 	RatingDate   string `gorm:"type:varchar(50)" json:"ratingDate"`
// }
func CreateRating(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Origin", "*")
	DB := common.GetDB()
	var requestRating model.Rating
	err := c.ShouldBindJSON(&requestRating)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(422, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}
	username := requestRating.Username
	restaurantName := requestRating.RestaurantName
	star := requestRating.Star
	comment := requestRating.Comment
	ratingDate := requestRating.RatingDate
	if !isCustomerExist(DB, username) || username == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Customer does not exist"})
		return
	}

	if !isRestaurantExist(DB, restaurantName) || restaurantName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Restaurant does not exist"})
		return
	}

	newRating := model.Rating{
		Username:       username,
		RestaurantName: restaurantName,
		Star:           star,
		Comment:        comment,
		RatingDate:     ratingDate,
	}
	DB.Create(&newRating)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "Successfully"})
}

func GetRating(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Origin", "*")
	DB := common.GetDB()
	var requestRating model.Rating
	err := c.ShouldBindQuery(&requestRating)
	if err != nil {
		c.JSON(422, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}

	restaurantName := requestRating.RestaurantName
	if !isRestaurantExist(DB, restaurantName) || restaurantName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Restaurant does not exist!"})
		return
	}

	targetRatings := []model.Rating{}
	DB.Where("restaurant_name = ?", restaurantName).Find(&targetRatings)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": targetRatings, "msg": "Successfully"})
}

func isRestaurantExist(db *gorm.DB, restaurantName string) bool {
	var restaurant model.Restaurant
	db.Where("name = ?", restaurantName).First(&restaurant)
	return restaurant.Name != ""
}
