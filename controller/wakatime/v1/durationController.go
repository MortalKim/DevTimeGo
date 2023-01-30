package v1

import (
	"WakaTImeGo/constant"
	"WakaTImeGo/model/entity/request/duration"
	durationService "WakaTImeGo/service/duration"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Kim
 * @Description:
 * @File:  durationController
 * @Date: 12/19/2022 10:35 AM
 */

func GetDurationsByParams(c *gin.Context) {
	//Get Json Params
	var searchParams duration.SearchParams

	err := c.ShouldBind(&searchParams)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": err.Error()})
		return
	}
	//If UserID is empty, get UserID from header
	if searchParams.UserID == "" {
		searchParams.UserID = c.Request.Header.Get(constant.DECRYPTED_USER_ID)
	}
	durations := durationService.GetDurationByParams(searchParams)
	c.JSON(200, durations)
}
