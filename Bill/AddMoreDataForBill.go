package bill

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddNewProductData(db gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		BillId := ctx.Param("billid")

		var data datastruct.SoldList

		err := ctx.ShouldBind(&data)
		data.BillId = BillId
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		err = db.
			Table("soldhistory").
			Create(&data).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

	}
}
