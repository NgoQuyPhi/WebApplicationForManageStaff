package main

import (
	dbconnect "PersonalProject/QlyNV/DBconnect"
)

func main() {
	db := dbconnect.DbConnection()
	defer db.Close()

}
