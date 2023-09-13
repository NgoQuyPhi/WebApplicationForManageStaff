package timekeepping

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckIn(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		StaffId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		now := time.Now().UTC()

		data := datastruct.TimeKeepping{
			StaffId: StaffId,
			CheckIn: &now,
		}

		day := now.Day()

		if err := db.Table("timekeepping").
			Where("StaffId = ? and DAY(CheckIn) = ?", StaffId, day).
			Delete(nil).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
		}

		err = db.Table("timekeepping").
			Where("StaffId = ?", StaffId).
			Create(&data).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"CheckIn:": "Success",
		})
	}
}
