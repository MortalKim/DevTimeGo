package v1

import (
	"WakaTImeGo/constant"
	"WakaTImeGo/model/entity"
	"WakaTImeGo/model/entity/wakatime/v1/response"
	"WakaTImeGo/task"
	"WakaTImeGo/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"net/http"
)

/**
 * @Author: Kim
 * @Description:
 * @File:  heartbeat
 * @Date: 12/12/2022 5:05 PM
 */

var log = logrus.New()

// SaveHeartbeat save heartbeat
func SaveHeartbeat(c *gin.Context) {
	//get heartbeats data
	var heartbeats entity.Heartbeat
	data, err := c.GetRawData()

	//use custom jsoniter to parse data
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(data, &heartbeats)

	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": err.Error()})
		return
	}
	log.Info(heartbeats)

	//get machine info
	machineName := c.Request.Header.Get("X-Machine-Name")
	//get user info by token
	userId := c.Request.Header.Get(constant.DECRYPTED_USER_ID)
	//get user agent
	userAgent := c.Request.Header.Get("User-Agent")
	//get os info and editor info
	opSys, editor, _ := utils.ParseUserAgent(userAgent)

	res := response.HeartbeatBulkResponse{}
	heartbeats.UserID = userId
	heartbeats.OperatingSystem = opSys
	heartbeats.Editor = editor
	heartbeats.Machine = machineName
	heartbeats.HashSelf()
	err2 := heartbeats.Add()
	if err2 != nil {
		log.Error(err)
	} else {
		//make response
		r := make([]interface{}, 2)
		r[0] = nil
		r[1] = http.StatusCreated
		res.Responses = append(res.Responses, r)
	}

	c.JSON(http.StatusCreated, res)
}

// SaveHeartbeats save heartbeats bulk
func SaveHeartbeats(c *gin.Context) {
	//get heartbeats data
	var heartbeats []entity.Heartbeat
	data, err := c.GetRawData()

	//use custom jsoniter to parse data
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(data, &heartbeats)

	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": err.Error()})
		return
	}
	log.Info(heartbeats)

	//get machine info
	machineName := c.Request.Header.Get("X-Machine-Name")
	//get user info by token
	userId := c.Request.Header.Get(constant.DECRYPTED_USER_ID)
	//get user agent
	userAgent := c.Request.Header.Get("User-Agent")
	//get os info and editor info
	opSys, editor, _ := utils.ParseUserAgent(userAgent)

	res := response.HeartbeatBulkResponse{}

	for _, h := range heartbeats {
		h.UserID = userId
		h.OperatingSystem = opSys
		h.Editor = editor
		h.Machine = machineName
		h.HashSelf()
		err := h.Add()
		success := true
		if err != nil {
			if _, ok := err.(*mysql.MySQLError); ok {
				//if duplicate key error, skip, or else return error
				if err.(*mysql.MySQLError).Number != 1062 {
					success = false
				}
			} else {
				success = false
			}

			log.Error(err)
		}

		if success {
			//make response
			r := make([]interface{}, 2)
			r[0] = nil
			r[1] = http.StatusCreated
			res.Responses = append(res.Responses, r)
		}
	}
	task.NewDurationDeliveryTask(userId)
	c.JSON(http.StatusCreated, res)
}
