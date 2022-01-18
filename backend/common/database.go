package common

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB {
	driverName := viper.GetString("Mysql.driverName")
	host := viper.GetString("Mysql.host")
	port := viper.GetString("Mysql.port")
	database := viper.GetString("Mysql.database")
	username := viper.GetString("Mysql.username")
	password := viper.GetString("Mysql.password")
	charset := viper.GetString("Mysql.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	} else {
		fmt.Println("Successfully connect to database")
	}

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
