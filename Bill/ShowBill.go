package bill

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowBill(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data []datastruct.BillCalculation
		BillId := ctx.Param("billid")
		err := db.Table("soldhistory").
			Select("soldhistory.ProductId,warehouse.ProductName, warehouse.Price,soldhistory.Quantity").
			Joins("LEFT JOIN warehouse ON soldhistory.ProductId = warehouse.ProductId").
			Where("soldhistory.BillId = ?", BillId).
			Find(&data).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		var totalbill float32 = 0

		for i := range data {
			totalbill += data[i].Price * float32(data[i].Quantity)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"": PrintBill(BillId, data, totalbill),
		})

	}
}

func PrintBill(billid string, data []datastruct.BillCalculation, totalbill float32) string {
	var Bill string = fmt.Sprint("BillId: ", billid)

	for i := range data {
		s := fmt.Sprint(" ProductName: ", data[i].ProductName, " Quantity: ", data[i].Quantity, "\n")
		Bill += s
	}

	Bill += fmt.Sprint("TotalBill: ", totalbill)

	return Bill
}
