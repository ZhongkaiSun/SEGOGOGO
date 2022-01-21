package model

// CREATE TABLE restaurant (
// 	ID VARCHAR ( 50 ) PRIMARY KEY,
// 	NAME VARCHAR ( 50 ),
// 	Address VARCHAR ( 100 ),
// 	deliveryFee FLOAT,
// 	imgPath VARCHAR ( 100 ),
// 	type VARCHAR ( 50 ),
// 	rating FLOAT
// );

type Restaurant struct {
	ID          string  `gorm:"type:varchar(50);primary_key" json:"id"`
	Name        string  `gorm:"type:varchar(50)" json:"name"`
	Address     string  `gorm:"type:varchar(100)" json:"address"`
	DeliveryFee float64 `gorm:"type:float" json:"deliveryFee"`
	ImgPath     string  `gorm:"type:varchar(100)" json:"imgPath"`
	TypeofMeal  string  `gorm:"type:varchar(100)" json:"typeofMeal"`
	Rating      float64 `gorm:"type:float" json:"rating"`
}
