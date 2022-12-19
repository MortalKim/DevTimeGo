package v1

import (
	"WakaTImeGo/constant"
	wakatime "WakaTImeGo/model/entity/wakatime/v1"
	"WakaTImeGo/service/summaryService"
	"WakaTImeGo/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

/**
 * @Author: Kim
 * @Description: Wakatime Api v1 today fetch
 * @File:  today
 * @Date: 12/12/2022 3:37 PM
 */
var logger = logrus.New()

func GetToday(c *gin.Context) {
	userId := c.Request.Header.Get(constant.DECRYPTED_USER_ID)
	//get today's time
	start, end := utils.GetDayTime(time.Now())
	summary := summaryService.GetSummaryByTimeRange(userId, start, end)

	todayData := wakatime.Summary{}
	todayData.Data = wakatime.SummaryData{}
	todayData.Data.GrandTotal = wakatime.GrandTotal{
		Digital:      "100 h",
		Hours:        1,
		Minutes:      1,
		Text:         "100 h",
		TotalSeconds: 1,
	}
	todayData.Data.Range = wakatime.Range{
		Date:     time.Now().Format("yyyy-mm-dd"),
		End:      time.Now(),
		Start:    time.Now(),
		Text:     "100 h",
		Timezone: "100 h",
	}
	c.JSON(http.StatusOK, summary)
}
