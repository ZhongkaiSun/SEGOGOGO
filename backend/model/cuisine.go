package model

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
