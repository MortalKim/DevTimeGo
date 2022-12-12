package v1

import (
	wakatime "WakaTImeGo/model/entity/wakatime/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * @Author: Kim
 * @Description: Wakatime Api v1 today fetch
 * @File:  today
 * @Date: 12/12/2022 3:37 PM
 */

func GetToday(c *gin.Context) {
	todayData := wakatime.Today{}
	todayData.Data = wakatime.TodayData{}
	todayData.Data.GrandTotal = wakatime.GrandTotal{
		Digital:      "100 h",
		Hours:        1,
		Minutes:      1,
		Text:         "100 h",
		TotalSeconds: 1,
	}
	todayData.Data.Range = wakatime.Range{
		Date:     time.Now(),
		End:      time.Now(),
		Start:    time.Now(),
		Text:     "100 h",
		Timezone: "100 h",
	}
	c.JSON(http.StatusOK, todayData)
}
