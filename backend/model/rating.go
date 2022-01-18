package model

// CREATE TABLE ratings (
// 	id int PRIMARY KEY,
// 	username VARCHAR (50),
// 	restaurant_id VARCHAR ( 50 ),
// 	star INT,
// 	rating_date date
// );

type Rating struct {
	ID           string `gorm:"type:varchar(50)"`
	Username     string `gorm:"type:varchar(50)"`
	RestaurantId string `gorm:"type:varchar(50)"`
	Star         int    `gorm:"type:int"`
	Comment      string `gorm:"type:varchar(255)"`
	rating_date  string `gorm:"type:int"`
}
