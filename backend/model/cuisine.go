package model

// CREATE TABLE cuisines (
// 	name VARCHAR ( 50 ),
// 	restaurant_id VARCHAR ( 50 ),
// 	price FLOAT,
// 	calories INT,
// 	PRIMARY KEY (name, restaurant_id)
// );

type Cuisine struct {
	Name         string  `gorm:"type:varchar(50)"`
	RestaurantId string  `gorm:"type:varchar(50)"`
	Price        float64 `gorm:"type:float"`
	Calories     int     `gorm:"type:int"`
}
