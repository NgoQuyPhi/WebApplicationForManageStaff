package staff

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get_Staff_Data_By_Id(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data datastruct.StaffData

		StaffID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		if err := db.Table("managestaff").
			Where("StaffId = ?", StaffID).
			First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
