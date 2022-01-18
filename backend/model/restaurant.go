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
	ID          string  `gorm:"type:varchar(50)"`
	Name        string  `gorm:"type:varchar(50)"`
	Address     string  `gorm:"type:varchar(100)"`
	DeliveryFee float64 `gorm:"type:float"`
	ImgPath     string  `gorm:"type:varchar(100)"`
	TypeofMeal  string  `gorm:"type:varchar(100)"`
	Rating      float64 `gorm:"type:float"`
}
