package timekeepping

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTimeKeeppingDataByStaffId(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data []datastruct.TimeKeepping

		StaffId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		err = db.Table("timekeepping").
			Where("StaffId = ?", StaffId).
			Find(&data).Error

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"TimeKeepping Calendar:": data,
		})
	}
}
