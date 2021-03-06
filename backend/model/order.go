package model

// CREATE TABLE cuisines (
// 	name VARCHAR ( 50 ),
// 	restaurant_id VARCHAR ( 50 ),
// 	price FLOAT,
// 	calories INT,
// 	PRIMARY KEY (name, restaurant_id)
// );

//create query

type Order struct {
	Username       string  `gorm:"type:varchar(50);primary_key" json:"username" form:"username"`
	RestaurantName string  `gorm:"type:varchar(50)" json:"restaurantName"`
	OrderDate      string  `gorm:"type:varchar(50);primary_key" time_format:"02/02/2022" json:"orderDate"`
	Price          float64 `gorm:"type:float" json:"price"`
	CuisineName    string  `gorm:"type:varchar(100);primary_key" json:"cuisineName"`
}
