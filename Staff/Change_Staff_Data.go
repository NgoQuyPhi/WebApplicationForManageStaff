package staff

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ChangeStaffData(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		StaffId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		var data datastruct.StaffData

		err = ctx.ShouldBind(&data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		err = db.Table("managestaff").
			Where("StaffId = ?", StaffId).
			Updates(&data).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}
	}
}
