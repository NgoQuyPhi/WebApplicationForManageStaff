package bill

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteData(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		BillId := ctx.Param("billid")
		ProductId := ctx.Param("productid")
		var data datastruct.SoldList

		err := db.Table("soldlist").Where("BillId = ? AND ProductId = ?", BillId, ProductId).First(&data).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}
		err = db.Table("soldlist").Where("BillId = ? AND ProductId = ?", BillId, ProductId).Delete(nil).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		err = db.Exec("UPDATE warehouse SET Quantity = Quantity + ? WHERE ProductId = ?", data.Quantity, data.ProductId).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Change Bill": "Success",
		})

	}
}
