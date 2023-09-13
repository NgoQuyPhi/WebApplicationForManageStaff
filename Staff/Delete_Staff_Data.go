package staff

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete_Staff_Data(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		StaffId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		err = db.Table("managestaff").Where("StaffId = ?", StaffId).Delete(nil).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Delete Staff Data:": "Success",
		})
	}
}
