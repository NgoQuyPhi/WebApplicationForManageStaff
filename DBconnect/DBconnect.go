package dbconnect

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func DbConnection() *sql.DB {
	dns := "ngoquyphi:ngoquyphi@tcp(127.0.0.1:3306)/"

	db, err := sql.Open("mysql", dns)
	PanicErr(err)

	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS ManageWatchShop")
	PanicErr(err)

	dns = "ngoquyphi:ngoquyphi@tcp(127.0.0.1:3306)/ManageWatchShop"

	db, err = sql.Open("mysql", dns)
	PanicErr(err)
	defer db.Close()

	_, err = db.Exec("CREATE Table IF NOT EXISTS ManageStaff(StaffId INT PRIMARY KEY,DepartmentId CHAR(4),StaffName VARCHAR(50), PhoneNumber CHAR(10), Address VARCHAR(50),Salary DOUBLE)")
	PanicErr(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS TimeKeepping(StaffId INT PRIMARY KEY,CheckIn TIMESTAMP, CheckOut TIMESTAMP,FOREIGN KEY (StaffId) REFERENCES ManageStaff(StaffId))")
	PanicErr(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS WareHouse(ProductId CHAR(6) PRIMARY KEY, ProductName VARCHAR(30),Price DOUBLE, Quantity INT)")
	PanicErr(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS SoldHistory(ProductId CHAR(6),StaffId INT,SellTime TIMESTAMP,Price DOUBLE)")
	PanicErr(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS ImportWarehouse(ProductId CHAR(6), Quantity INT,ImportTime TIMESTAMP)")
	PanicErr(err)
	return db
}
