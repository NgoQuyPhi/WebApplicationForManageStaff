package dbconnect

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() (db *gorm.DB) {
	dsn := "ngoquyphi:ngoquyphi@tcp(127.0.0.1:3306)/ManageWatchShop"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
