package model

// CREATE TABLE customers (
// 	username VARCHAR ( 50 ) PRIMARY KEY,
// 	password VARCHAR ( 50 ) NOT NULL,
// 	address_line1 VARCHAR ( 50 ) NOT NULL,
// 	address_line2 VARCHAR ( 50 ),
// 	phone VARCHAR ( 50 ) );
type Customer struct {
	Username     string `gorm:"type:varchar(50);primary_key" json:"username" form:"username" binding:"required"`
	Password     string `gorm:"type:varchar(50)" json:"password" form:"password" binding:"required"`
	AddressLine1 string `gorm:"type:varchar(50)" json:"addressLine1"`
	AddressLine2 string `gorm:"type:varchar(50)" json:"addressLine2"`
	Phone        string `gorm:"type:varchar(50)" json:"phone"`
	Email        string `gorm:"type:varchar(50)" json:"email"`
	City         string `gorm:"type:varchar(50)" json:"city"`
	State        string `gorm:"type:varchar(50)" json:"state"`
	Zipcode      string `gorm:"type:varchar(50)" json:"zipcode"`
}

//payment api (create, read)
