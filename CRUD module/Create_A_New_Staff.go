package crudmodule

import (
	datastruct "PersonalProject/QlyNV/DataStruct"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create_New_Staff_Data(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data datastruct.NewStaffData
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"Error": err,
			})
			return
		}

		if err := db.Table("ManageStaff").Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error:": err,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Create New Data:": "Success",
		})

	}
}
