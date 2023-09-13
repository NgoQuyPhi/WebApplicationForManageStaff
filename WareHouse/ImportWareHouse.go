package warehouse

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ImportWareHouse(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data datastruct.ImportProduct

		err := ctx.ShouldBind(&data)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		now := time.Now().UTC()
		data.ImportTime = &now

		err = db.Table("importwarehouse").Create(&data).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		var tempdata datastruct.Product

		err = db.Table("warehouse").Where("ProductId = ?", data.ProductId).First(&tempdata).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		tempdata.Quantity += data.Quantity

		err = db.Table("warehouse").Where("ProductId = ?", data.ProductId).Updates(&tempdata).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Import:": "Success",
		})

	}
}
