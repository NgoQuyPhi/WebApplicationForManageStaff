package main

import (
	bill "PersonalProject/QlyNV/Bill"
	dbconnect "PersonalProject/QlyNV/DBconnect"
	timekeepping "PersonalProject/QlyNV/TimeKeepping"
	warehouse "PersonalProject/QlyNV/WareHouse"

	staff "PersonalProject/QlyNV/Staff"

	"github.com/gin-gonic/gin"
)

func main() {
	dbconnect.DBCreate()
	db := dbconnect.ConnectToDB()

	r := gin.Default()

	Staff := r.Group("/Staff")
	{
		Staff.GET("/alldata", staff.Get_Staff_Data(db))
		Staff.POST("/createnewdata", staff.Create_New_Staff_Data(db))
		Staff.GET("/staffdata/:id", staff.Get_Staff_Data_By_Id(db))
		Staff.PUT("/change/:id", staff.ChangeStaffData(db))
		Staff.DELETE("/deletestaff/:id", staff.Delete_Staff_Data(db))
	}

	TimeKeepping := r.Group("/tk")
	{
		TimeKeepping.POST("/checkin/:id", timekeepping.CheckIn(db))
		TimeKeepping.PUT("/checkout/:id", timekeepping.CheckOut(db))
		TimeKeepping.GET("/data/:id", timekeepping.GetTimeKeeppingDataByStaffId(db))
	}

	WareHouse := r.Group("/warehouse")
	{
		WareHouse.POST("/newproduct", warehouse.CreateNewProduct(db))
		WareHouse.POST("/import", warehouse.ImportWareHouse(db))
		WareHouse.GET("/checkwarehouse", warehouse.GetWarehouseData(db))
	}

	Bill := r.Group("/bill")
	{
		Bill.POST("/newbill", bill.CreateNewBill(db))
		Bill.GET("/showbill/:billid", bill.ShowBill(db))

	}

	r.Run()

}
