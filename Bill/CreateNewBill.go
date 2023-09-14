package bill

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	random "PersonalProject/QlyNV/RandomSeed"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNewBill(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		BillId := random.String(4)
		var data []datastruct.SoldList
		err := ctx.ShouldBind(&data)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}
		now := time.Now().UTC()

		for i := range data {
			data[i].BillId = BillId
			data[i].SoldTime = &now
		}

		err = db.Table("soldhistory").Create(&data).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		for i := range data {
			err = db.Exec("UPDATE warehouse SET Quantity = Quantity - ? WHERE ProductId = ?", data[i].Quantity, data[i].ProductId).Error
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Error:": err,
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Create New Bill:": "Success",
		})

	}
}
