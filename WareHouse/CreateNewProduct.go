package warehouse

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	random "PersonalProject/QlyNV/RandomSeed"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNewProduct(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data datastruct.Product

		err := ctx.ShouldBind(&data)

		data.ProductId = random.String(6)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		err = db.Table("warehouse").Create(&data).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Add New Product:": "Success",
		})

	}
}
