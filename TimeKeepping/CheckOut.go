package timekeepping

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckOut(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		now := time.Now().UTC()
		StaffId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}
		data := datastruct.TimeKeepping{
			StaffId:  StaffId,
			CheckOut: &now,
		}

		err = db.Table("timekeepping").
			Where("StaffId = ? AND DAY(CheckIn) = ?", StaffId, now.Day()).
			Updates(&data).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"CheckOut:": "Success",
		})
	}
}
