package model

// CREATE TABLE customers (
// 	username VARCHAR ( 50 ) PRIMARY KEY,
// 	password VARCHAR ( 50 ) NOT NULL,
// 	address_line1 VARCHAR ( 50 ) NOT NULL,
// 	address_line2 VARCHAR ( 50 ),
// 	phone VARCHAR ( 50 ) );
type Customer struct {
	Username     string `gorm:"type:varchar(50)"`
	Password     string `gorm:"type:varchar(50)"`
	AddressLine1 string `gorm:"type:varchar(50)"`
	AddressLine2 string `gorm:"type:varchar(50)"`
	Phone        string `gorm:"type:varchar(50)"`
	Email        string `gorm:"type:varchar(50)"`
}
