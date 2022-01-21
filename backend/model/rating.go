package model

// CREATE TABLE ratings (
// 	id int PRIMARY KEY,
// 	username VARCHAR (50),
// 	restaurant_id VARCHAR ( 50 ),
// 	star INT,
// 	rating_date date
// );

type Rating struct {
	ID           string `gorm:"type:varchar(50)" json:"id"`
	Username     string `gorm:"type:varchar(50)" json:"username"`
	RestaurantId string `gorm:"type:varchar(50)" json:"restaurantId"`
	Star         int    `gorm:"type:int" json:"star"`
	Comment      string `gorm:"type:varchar(255)" json:"comment"`
	RatingDate   string `gorm:"type:varchar(50)" json:"ratingDate"`
}
