package v1

import (
	wakatime "WakaTImeGo/model/entity/wakatime/v1"
	"WakaTImeGo/service/heartbeat"
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
	//get today's heartbeat
	start, end := utils.GetDayTimeString(time.Now())
	heartbeats := heartbeat.GetHeartbeatByTime(start, end)
	logger.Info(heartbeats)

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
	c.JSON(http.StatusOK, todayData)
}
