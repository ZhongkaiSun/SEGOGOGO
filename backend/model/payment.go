package model

// CREATE TABLE payment (
// 	username VARCHAR ( 50 ) PRIMARY KEY,
// 	card_number VARCHAR ( 50 ),
// 	exp_date VARCHAR ( 50 ),
// 	security_code VARCHAR ( 50 ),
// 	address_line1 VARCHAR ( 50 ),
// 	address_line2 VARCHAR ( 50 ),
// 	city VARCHAR ( 50 ),
// 	state VARCHAR ( 50 ),
// 	zipcode VARCHAR ( 50 ),
// )

type Payment struct {
	Username     string `gorm:"type:varchar(50);primary_key" json:"username" binding:"required"`
	CardNumber   string `gorm:"type:varchar(50)" json:"cardNumber"`
	ExpDate      string `gorm:"type:varchar(50)" json:"expDate"`
	SecurityCode string `gorm:"type:varchar(50)" json:"securityCode"`
	AddressLine1 string `gorm:"type:varchar(50)" json:"addressLine1"`
	AddressLine2 string `gorm:"type:varchar(50)" json:"addressLine2"`
	City         string `gorm:"type:varchar(50)" json:"city"`
	State        string `gorm:"type:varchar(50)" json:"state"`
	Zipcode      string `gorm:"type:varchar(50)" json:"zipcode"`
}
