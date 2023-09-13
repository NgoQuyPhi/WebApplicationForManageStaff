package warehouse

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetWarehouseData(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data []datastruct.Product

		err := db.Table("warehouse").Find(&data).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"": data,
			})
		}
	}
}
