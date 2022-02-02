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
// 	RestaurantId string `gorm:"type:varchar(50)" json:"restaurantId"`
// 	Star         int    `gorm:"type:int" json:"star"`
// 	Comment      string `gorm:"type:varchar(255)" json:"comment"`
// 	RatingDate   string `gorm:"type:varchar(50)" json:"ratingDate"`
// }
func CreateRating(c *gin.Context) {
	DB := common.GetDB()
	var requestRating model.Rating
	err := c.ShouldBindJSON(&requestRating)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}
	id := requestRating.ID
	username := requestRating.Username
	restaurantId := requestRating.RestaurantId
	star := requestRating.Star
	comment := requestRating.Comment
	ratingDate := requestRating.RatingDate
	if !isCustomerExist(DB, username) || username == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Customer does not exist"})
		return
	}

	if !isRestaurantExist(DB, restaurantId) || restaurantId == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Restaurant does not exist"})
		return
	}

	newRating := model.Rating{
		ID:           id,
		Username:     username,
		RestaurantId: restaurantId,
		Star:         star,
		Comment:      comment,
		RatingDate:   ratingDate,
	}
	DB.Create(&newRating)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "Successfully"})
}

func GetRating(c *gin.Context) {
	DB := common.GetDB()
	var requestRating model.Rating
	err := c.ShouldBindQuery(&requestRating)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":   "Binding error",
			"error": err,
			"data":  gin.H{},
		})
		return
	}

	restaurantId := requestRating.RestaurantId
	if !isRestaurantExist(DB, restaurantId) || restaurantId == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "msg": "Create a new username"})
		return
	}

	targetRatings := []model.Rating{}
	DB.Where("restaurant_id = ?", restaurantId).Find(&targetRatings)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": targetRatings, "msg": "Successfully"})
}

func isRestaurantExist(db *gorm.DB, restaurantId string) bool {
	var restaurant model.Restaurant
	db.Where("ID = ?", restaurantId).First(&restaurant)
	return restaurant.ID != ""
}
